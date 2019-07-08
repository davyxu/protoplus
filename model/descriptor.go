package model

import (
	"errors"
	"fmt"
)

// 结构体或枚举
type Descriptor struct {
	Comment
	TagSet

	Name string

	// 枚举或结构体
	Kind Kind

	// 归属的文件名
	SrcName string

	// 结构体和枚举
	Fields []*FieldDescriptor `json:",omitempty"`

	// 服务调用
	SvcCall []*ServiceCall `json:",omitempty"`

	DescriptorSet *DescriptorSet `json:"-"`
}

func (self *Descriptor) FieldByName(name string) *FieldDescriptor {

	for _, o := range self.Fields {
		if o.Name == name {
			return o
		}
	}

	return nil
}

func (self *Descriptor) FieldNameExists(name string) bool {

	for _, o := range self.Fields {
		if o.Name == name {
			return true
		}
	}

	return false
}

func (self *Descriptor) CallNameExists(name string) bool {

	for _, o := range self.SvcCall {
		if o.Name == name {
			return true
		}
	}

	return false
}

func (self *Descriptor) FieldTagExists(tag int) bool {

	// 没填不会重复
	if tag == 0 {
		return false
	}

	for _, o := range self.Fields {
		if o.Tag == tag {
			return true
		}
	}

	return false
}

func (self *Descriptor) AddField(fd *FieldDescriptor) {
	self.Fields = append(self.Fields, fd)
}

func (self *Descriptor) AddSvcCall(sc *ServiceCall) {
	self.SvcCall = append(self.SvcCall, sc)
}

func (self *Descriptor) Size() (size int32) {

	for _, f := range self.Fields {

		if f.Repeatd {
			panic(errors.New("Nonsupport repeated"))
		}

		switch f.Kind {
		case Kind_Primitive:
			cs := TypeSize(f.Type)
			if cs == 0 {
				panic(errors.New(fmt.Sprintf("Nonsupport %s", f.Type)))
			}
			size += cs
		case Kind_Enum:
			size += 4
		case Kind_Struct:
			if cd := self.DescriptorSet.ObjectByName(f.Type); cd != nil {
				size += cd.Size()
			}
		}
	}

	return
}
