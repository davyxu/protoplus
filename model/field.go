package model

type FieldDescriptor struct {
	Comment
	TagSet

	Name    string
	Type    string
	Kind    Kind
	Tag     int  `json:",omitempty"`
	Repeatd bool `json:",omitempty"`
}

func (self *FieldDescriptor) ParseType(str string) {

	self.Type = Str2Type[str]
	if self.Type != "" {
		self.Kind = Kind_Primitive
	} else {
		// 复杂类型
		self.Type = str
	}
}
