package model

type FieldDescriptor struct {
	Comment
	TagSet

	Name string
	Type string

	Kind Kind // 原始类型/结构体/枚举

	Tag     int  `json:",omitempty"`
	Repeatd bool `json:",omitempty"`

	Descriptor *Descriptor `json:"-"` // 字段归属的父级描述符
}

func (self *FieldDescriptor) ParseType(str string) {

	self.Type = SchemeType2Type[str]
	if self.Type != "" {
		self.Kind = Kind_Primitive
	} else {
		// 复杂类型
		self.Type = str
	}
}
