package msgidutil

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/model"
	"io/ioutil"
	"os"
)

var skipDupCacheMsgIDTips bool // 跳过重复缓存消息ID的警告提示

var msgNameByMsgID = map[int]string{}

type MsgInfo struct {
	Name  string
	MsgID int
}

type AutoMsgIDCacheFile struct {
	Caches []MsgInfo
}

func (self *AutoMsgIDCacheFile) Load(cacheFileName string) {
	existsFileData, _ := ioutil.ReadFile(cacheFileName)

	json.Unmarshal(existsFileData, self)
}

func (self *AutoMsgIDCacheFile) Save(cacheFileName string) {
	data, _ := json.MarshalIndent(self, "", "\t")

	ioutil.WriteFile(cacheFileName, data, 0666)
}

func (self *AutoMsgIDCacheFile) ExistsMsgID(name string) int {
	for _, m := range self.Caches {
		if m.Name == name {
			return m.MsgID
		}
	}

	return 0
}

func (self *AutoMsgIDCacheFile) GetNameByID(msgid int) string {

	for _, m := range self.Caches {
		if m.MsgID == msgid {
			return m.Name
		}
	}

	return ""
}

func (self *AutoMsgIDCacheFile) AddNewMsg(name string, msgid int) {

	self.Caches = append(self.Caches, MsgInfo{name, msgid})
}

func (self *AutoMsgIDCacheFile) ModifyByMsgID(msgid int, name string) {

	for index, m := range self.Caches {
		if m.MsgID == msgid {
			self.Caches[index].Name = name
			return
		}
	}

}

// MsgId:100 Descriptor  FileA
// Descriptor		<- auto gen 101
// MsgID:200 Descriptor FileB
// Descriptor		<- auto gen 201

func genMsgID(d *model.Descriptor, cacheFileName string, cachedMsgIDs *AutoMsgIDCacheFile) int {

	var msgid = 0
	for _, obj := range d.DescriptorSet.Objects {

		userMsgID := obj.TagValueInt("MsgID")

		if userMsgID == 0 && !obj.TagExists("AutoMsgID") {
			continue
		}

		if userMsgID > 0 {
			msgid = userMsgID
		} else {
			msgid++
		}

		if obj == d {

			if existsName := cachedMsgIDs.GetNameByID(msgid); existsName != "" {

				if !skipDupCacheMsgIDTips && *flagShowOverWriteCacheFileWarning {

					fmt.Println("Warning: auto generate msg id has exists in automsgidcache file, the file will be overwrited.")
					bufio.NewReader(os.Stdin).ReadString('\n')

					skipDupCacheMsgIDTips = true
				}

				// msgid已存在,msgid拿给该消息使用
				cachedMsgIDs.ModifyByMsgID(msgid, d.Name)
			} else {
				// msgid不存在，添加
				cachedMsgIDs.AddNewMsg(d.Name, msgid)
			}

			cachedMsgIDs.Save(cacheFileName)

			return msgid
		}
	}

	// 不会运行到这里的
	return 0
}

func autogenMsgIDByCacheFile(cacheFileName string, d *model.Descriptor) (newMsgID int) {

	var cachedMsgIDs AutoMsgIDCacheFile
	cachedMsgIDs.Load(cacheFileName)

	// 协议里用户指定的ID
	userMsgID := d.TagValueInt("MsgID")

	// 文件中，这个消息已经记忆的ID
	existsMsgID := cachedMsgIDs.ExistsMsgID(d.Name)

	switch {
	case userMsgID == 0 && existsMsgID == 0: // 缓冲无记录，用户没有指定ID，生成新ID
		newMsgID = genMsgID(d, cacheFileName, &cachedMsgIDs)
	case userMsgID != 0 && existsMsgID == 0: // 缓冲无记录，用户指定ID，用用户指定的ID
		return userMsgID
	case userMsgID == 0 && existsMsgID != 0: // 缓冲有记录ID，用户没有指定ID，用缓冲ID
		return existsMsgID
	case userMsgID != 0 && existsMsgID != 0: // 缓冲有记录ID，用户指定ID, 优先使用用户ID
		return userMsgID
	}

	return
}

// 字符串转为16位整形值
func StringHash(s string) (hash uint16) {

	for _, c := range s {

		ch := uint16(c)

		hash = hash + ((hash) << 5) + ch + (ch << 7)
	}

	return
}

func StructMsgID(d *model.Descriptor) (msgid int) {
	if !codegen.IsMessage(d) {
		return 0
	}

	if d.Kind == model.Kind_Struct {
		msgid = d.TagValueInt("MsgID")
	}

	if *flagAutoMsgIDCacheFile != "" {
		msgid = autogenMsgIDByCacheFile(*flagAutoMsgIDCacheFile, d)
	} else if msgid == 0 {
		msgid = int(StringHash(d.DescriptorSet.PackageName + d.Name + *flagMsgIDSalt))
	}

	if *flagCheckDuplicateMsgID {

		oldName, exists := msgNameByMsgID[msgid]
		if exists && d.Name != oldName {
			panic(errors.New(fmt.Sprintf("%s's msgid(%d) has used by %s", d.Name, msgid, oldName)))
		}

		msgNameByMsgID[msgid] = d.Name

	}

	return
}

func init() {

	codegen.UsefulFunc["StructMsgID"] = StructMsgID
}
