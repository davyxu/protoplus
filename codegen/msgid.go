package codegen

import (
	"encoding/json"
	"flag"
	"github.com/davyxu/protoplus/model"
	"io/ioutil"
)

var flagAutoMsgIDCacheFile = flag.String("AutoMsgIDCacheFile", "AutoMsgIDCacheFile.json", "Specifies auto msgid cache file")

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

func (self *AutoMsgIDCacheFile) AddNewMsg(name string, msgid int) {

	self.Caches = append(self.Caches, MsgInfo{name, msgid})
}

// MsgId:100 Descriptor  FileA
// Descriptor		<- auto gen 101
// MsgID:200 Descriptor FileB
// Descriptor		<- auto gen 201

func genMsgID(d *model.Descriptor) int {

	var msgid = 0
	for _, obj := range d.DescriptorSet.Objects {

		if userMsgID := obj.TagValueInt("MsgID"); userMsgID != 0 {
			msgid = userMsgID
		} else {
			msgid++
		}

		if obj == d {
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
		newMsgID = genMsgID(d)
		cachedMsgIDs.AddNewMsg(d.Name, newMsgID)

		cachedMsgIDs.Save(cacheFileName)
	case userMsgID != 0 && existsMsgID == 0: // 缓冲无记录，用户指定ID，用用户指定的ID
		return userMsgID
	case userMsgID == 0 && existsMsgID != 0: // 缓冲有记录ID，用户没有指定ID，用缓冲ID
		return existsMsgID
	case userMsgID != 0 && existsMsgID != 0: // 缓冲有记录ID，用户指定ID, 优先使用用户ID
		return userMsgID
	}

	return
}

func init() {

	UsefulFunc["StructMsgID"] = func(raw interface{}) (msgid int) {
		d := raw.(*model.Descriptor)

		if d.Kind == model.Kind_Struct {
			msgid = d.TagValueInt("MsgID")
		}

		if *flagAutoMsgIDCacheFile != "" {
			msgid = autogenMsgIDByCacheFile(*flagAutoMsgIDCacheFile, d)
		}

		return
	}
}
