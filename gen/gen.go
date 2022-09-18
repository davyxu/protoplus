package gen

import (
	"github.com/davyxu/protoplus/model"
)

type Context struct {
	*model.DescriptorSet
	OutputFileName string
	ClassBase      string
}
