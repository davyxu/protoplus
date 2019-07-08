package model

type DescriptorSet struct {
	Objects     []*Descriptor `json:",omitempty"`
	PackageName string
}

func (self *DescriptorSet) Services() (ret []*Descriptor) {

	for _, o := range self.Objects {
		if o.Kind == Kind_Service {
			ret = append(ret, o)
		}
	}

	return
}

func (self *DescriptorSet) Structs() (ret []*Descriptor) {

	for _, o := range self.Objects {
		if o.Kind == Kind_Struct {
			ret = append(ret, o)
		}
	}

	return
}

func (self *DescriptorSet) Enums() (ret []*Descriptor) {

	for _, o := range self.Objects {
		if o.Kind == Kind_Enum {
			ret = append(ret, o)
		}
	}

	return
}

func (self *DescriptorSet) ObjectNameExists(name string) bool {

	for _, o := range self.Objects {
		if o.Name == name {
			return true
		}
	}

	return false
}

func (self *DescriptorSet) ObjectByName(name string) *Descriptor {

	for _, o := range self.Objects {
		if o.Name == name {
			return o
		}
	}

	return nil
}

func (self *DescriptorSet) AddObject(d *Descriptor) {
	self.Objects = append(self.Objects, d)
}
