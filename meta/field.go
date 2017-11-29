package meta

import (
	"fmt"
)

type FieldDescriptor struct {
	*CommentGroup
	Name    string
	Type    FieldType
	Tag     int         `json:",omitempty"`
	AutoTag int         `json:",omitempty"`
	Repeatd bool        `json:",omitempty"`
	Complex *Descriptor `json:",omitempty"`

	Struct *Descriptor `json:"-"`
}

func (self *FieldDescriptor) TagNumber() int {

	var tag int

	if self.AutoTag == -1 {
		tag = self.Tag
	} else {
		tag = self.AutoTag
	}

	if tag != 0 {
		return self.Struct.TagBase + tag
	}

	return tag
}

func (self *FieldDescriptor) TypeString() (ret string) {
	if self.Repeatd {
		ret = "[]"
	}

	return ret + self.TypeName()
}

func (self *FieldDescriptor) CompatibleTypeString() string {
	return self.typeStr(true)
}

func (self *FieldDescriptor) typeStr(compatible bool) (ret string) {

	return
}

func (self *FieldDescriptor) String() string {

	return fmt.Sprintf("%s %s = %d", self.Name, self.TypeString(), self.TagNumber())
}

func (self *FieldDescriptor) Kind() string {

	return self.Type.String()
}

func (self *FieldDescriptor) TypeName() string {

	switch self.Type {
	case FieldType_Bool:
		return "bool"

	case FieldType_Struct, FieldType_Enum:
		return self.Complex.Name
	default:
		return self.Type.String()
	}

}

func (self *FieldDescriptor) parseType(name string) (ft FieldType, structType *Descriptor) {

	ft = ParseFieldType(name)

	if ft != FieldType_None {
		return ft, nil
	}

	if ft, structType = self.Struct.File.FileSet.parseType(name); ft != FieldType_None {
		return ft, structType
	}

	return FieldType_None, nil

}

func NewFieldDescriptor(d *Descriptor) *FieldDescriptor {
	return &FieldDescriptor{
		Struct:  d,
		AutoTag: -1,
	}
}
