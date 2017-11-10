package meta

import "bytes"

type FileDescriptorSet struct {
	Files []*FileDescriptor

	lazyFields []*LazyField
}

func (self *FileDescriptorSet) ResolveAll() error {

	for _, v := range self.lazyFields {
		if _, err := v.Resolve(2); err != nil {
			return err
		}
	}

	return nil
}

func (self *FileDescriptorSet) AddLazyField(lf *LazyField) {
	self.lazyFields = append(self.lazyFields, lf)
}

func (self *FileDescriptorSet) AddFile(file *FileDescriptor) {
	file.FileSet = self
	self.Files = append(self.Files, file)
}

func (self *FileDescriptorSet) parseType(name string) (ft FieldType, structType *Descriptor) {

	for _, file := range self.Files {

		if ft, structType = file.rawParseType(name); ft != FieldType_None {
			return ft, structType
		}
	}

	return FieldType_None, nil
}

func (self *FileDescriptorSet) String() string {

	var bf bytes.Buffer

	bf.WriteString("[File]\n")
	for _, f := range self.Files {
		bf.WriteString(f.FileName)
		bf.WriteString("{\n")
		bf.WriteString(f.String())
		bf.WriteString("}\n")
	}

	bf.WriteString("\n")

	return bf.String()
}

func NewFileDescriptorSet() *FileDescriptorSet {
	return &FileDescriptorSet{}
}
