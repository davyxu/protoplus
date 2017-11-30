package model

type Descriptor struct {
	Comment
	TagSet

	Name string

	Kind Kind

	SrcName string

	Fields []*FieldDescriptor
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
