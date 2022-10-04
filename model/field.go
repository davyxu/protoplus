package model

import "strings"

type FieldDescriptor struct {
	Comment
	TagSet

	Name string
	Type string

	Kind Kind // 原始类型/结构体/枚举

	Tag     int  `json:",omitempty"`
	Repeatd bool `json:",omitempty"`

	MapKey   string
	MapValue string

	Descriptor *Descriptor `json:"-"` // 字段归属的父级描述符
}

func (self *FieldDescriptor) IsMap() bool {
	return self.MapKey != "" && self.MapValue != ""
}

func (self *FieldDescriptor) ParseType(str string) {

	if strings.Contains(str, "map") {
		println(str)
	}

	self.Type = SchemeType2Type[str]
	if self.Type != "" {
		self.Kind = Kind_Primitive
	} else {
		// 复杂类型
		self.Type = str
	}
}
