package meta

import (
	"errors"
	"fmt"
	"github.com/davyxu/golexer"
)

type LazyField struct {
	typeName string

	fd *FieldDescriptor

	d *Descriptor

	tp golexer.TokenPos

	miss bool
}

func NewLazyField(typeName string, fd *FieldDescriptor, d *Descriptor, tp golexer.TokenPos) *LazyField {
	return &LazyField{
		typeName: typeName,
		fd:       fd,
		d:        d,
		tp:       tp,
	}
}

func (self *LazyField) Resolve(pass int) (bool, error) {

	self.fd.Type, self.fd.Complex = self.fd.parseType(self.typeName)

	if self.fd.Type == FieldType_None {
		if pass > 1 {

			fmt.Println(self.tp.String())

			return true, errors.New("type not found: " + self.typeName)
		} else {

			self.miss = true
			return true, nil
		}
	}

	return false, nil
}
