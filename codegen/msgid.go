package codegen

import (
	"github.com/davyxu/protoplus/model"
	"strings"
)

// 字符串转为16位整形值
func stringHash(s string) (hash uint16) {

	for _, c := range s {
		ch := uint16(c)
		hash = hash + ((hash) << 5) + ch + (ch << 7)
	}

	return
}

func StructMsgID(d *model.Descriptor) (msgid int) {
	if !IsMessage(d) {
		return 0
	}

	if d.Kind == model.Kind_Struct {
		msgid = d.TagValueInt("MsgID")
	}

	if msgid == 0 {
		// 这里不能引入packageName, 不同语言的package设置为不同, 造成消息id不一致
		msgid = int(stringHash(strings.ToLower(d.Name)))
	}

	return
}

func init() {
	UsefulFunc["StructMsgID"] = StructMsgID
}
