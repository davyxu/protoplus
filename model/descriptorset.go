package model

type DescriptorSet struct {
	Objects []*Descriptor
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
