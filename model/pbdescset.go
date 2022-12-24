package model

import "sort"

type PBDescriptorSet struct {
	DescriptorSet

	// pb生成文件依赖时, 使用以下字段
	DependentSource []string                    // 按文件管理的描述符取出时, 这个字段有效. 本文件依赖的其他source的symbiol
	SourceName      string                      // 按文件管理的描述符取出时, 这个字段有效. 表示本DescriptorSet的文件名
	dsBySource      map[string]*PBDescriptorSet // 按文件名管理的描述符集合
}

func (self *PBDescriptorSet) addDependentSource(name string) {

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

func (self *PBDescriptorSet) SourceList() (ret []string) {
	for sourceName := range self.DescriptorSetBySource() {
		ret = append(ret, sourceName)
	}

	sort.Strings(ret)
	return
}

func (self *PBDescriptorSet) DescriptorSetBySource() map[string]*PBDescriptorSet {
	if self.dsBySource != nil {
		return self.dsBySource
	}

	self.dsBySource = map[string]*PBDescriptorSet{}

	for _, obj := range self.Objects {
		ds := self.dsBySource[obj.SrcName]
		if ds == nil {
			ds = &PBDescriptorSet{}

			ds.PackageName = self.PackageName
			ds.Codec = self.Codec
			ds.SourceName = obj.SrcName

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
