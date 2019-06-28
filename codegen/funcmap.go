package codegen

import (
	"github.com/davyxu/protoplus/model"
	"reflect"
	"strings"
	"text/template"
)

var UsefulFunc = template.FuncMap{}

func TagNumber(d *model.Descriptor, fd *model.FieldDescriptor) (tag int) {
	tag = -1
	for _, libfd := range d.Fields {

		if libfd.Tag != 0 {
			tag = libfd.Tag
		} else {
			tag++
		}

		if libfd == fd {
			return tag
		}
	}

	return 0
}

func init() {
	UsefulFunc["ObjectLeadingComment"] = func(raw interface{}) (ret string) {

		d := raw.(*model.Descriptor)

		if d.Leading != "" {

			for index, line := range strings.Split(d.Leading, "\n") {
				if index > 0 {
					ret += "\n"
				}
				ret += "// " + line
			}

			return
		}

		return
	}

	UsefulFunc["FieldTrailingComment"] = func(raw interface{}) string {

		fd := raw.(*model.FieldDescriptor)

		if fd.Trailing != "" {
			return "// " + fd.Trailing
		}

		return ""
	}

	UsefulFunc["TagNumber"] = func(rawD, rawFD interface{}) (tag int) {
		d := rawD.(*model.Descriptor)
		fd := rawFD.(*model.FieldDescriptor)

		return TagNumber(d, fd)
	}

	// 类pb的协议都使用这个Tag生成
	UsefulFunc["PbTagNumber"] = func(rawD, rawFD interface{}) (tag int) {
		d := rawD.(*model.Descriptor)
		fd := rawFD.(*model.FieldDescriptor)

		// 枚举从0自动生成
		if d.Kind == model.Kind_Enum {
			return TagNumber(d, fd)
		} else {
			// 自动从1开始, pb不允许字段tag为0
			return TagNumber(d, fd) + 1
		}

	}

	UsefulFunc["IsMessage"] = IsMessage

	// 生成Json尾巴的逗号，rawIdx为当前遍历索引，rawSlice传切片
	UsefulFunc["GenJsonTailComma"] = func(rawIdx, rawSlice interface{}) string {

		index := rawIdx.(int)

		total := reflect.Indirect(reflect.ValueOf(rawSlice)).Len()

		if index < total-1 {
			return ","
		}

		return ""
	}

}

func IsMessage(d *model.Descriptor) bool {
	return d.TagValueInt("MsgID") > 0 || d.TagExists("AutoMsgID")
}
