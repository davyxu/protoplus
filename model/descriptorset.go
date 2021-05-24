package model

type DescriptorSet struct {
	Objects     []*Descriptor `json:",omitempty"`
	PackageName string
	Codec       string

	DependentSource []string                  // 按文件管理的描述符取出时, 这个字段有效. 本文件依赖的其他source的symbiol
	SourceName      string                    // 按文件管理的描述符取出时, 这个字段有效. 表示本DescriptorSet的文件名
	dsBySource      map[string]*DescriptorSet // 按文件名管理的描述符集合
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

func (self *DescriptorSet) addDependentSource(name string) {

	if self.SourceName == name {
		return
	}

	for _, n := range self.DependentSource {
		if n == name {
			return
		}
	}

	self.DependentSource = append(self.DependentSource, name)
}

func (self *DescriptorSet) DescriptorSetBySource() map[string]*DescriptorSet {
	if self.dsBySource != nil {
		return self.dsBySource
	}

	self.dsBySource = map[string]*DescriptorSet{}

	for _, obj := range self.Objects {
		ds := self.dsBySource[obj.SrcName]
		if ds == nil {
			ds = &DescriptorSet{
				PackageName: self.PackageName,
				Codec:       self.Codec,
				SourceName:  obj.SrcName,
			}
			self.dsBySource[obj.SrcName] = ds
		}

		ds.AddObject(obj)
	}

	for _, file := range self.dsBySource {
		for _, st := range file.Structs() {

			for _, fd := range st.Fields {

				switch fd.Kind {
				case Kind_Struct, Kind_Enum:
					refTarget := self.ObjectByName(fd.Type)
					if refTarget != nil {
						file.addDependentSource(refTarget.SrcName)
					}
				}
			}
		}
	}

	return self.dsBySource
}
