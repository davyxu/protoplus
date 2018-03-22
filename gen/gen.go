package gen

import (
	"flag"
	"github.com/davyxu/protoplus/model"
)

type GenFunc func(*model.DescriptorSet, string) error

type Generator struct {
	FlagName    string
	FlagComment string

	GenFunc GenFunc

	flagStr *string

	flagBool *bool

	//builtin bool

	UseBoolFlag bool
}

func (self *Generator) flagString() string {
	if self.flagStr == nil {
		return ""
	}

	return *self.flagStr
}

func (self *Generator) gen(dset *model.DescriptorSet) error {

	if self.UseBoolFlag && *self.flagBool {
		return self.GenFunc(dset, "")
	} else if self.flagString() != "" {

		//if !self.builtin {
		//	fmt.Printf("%s:\n", self.FlagName)
		//}

		return self.GenFunc(dset, self.flagString())
	}

	return nil
}

func (self *Generator) init() {

	if self.UseBoolFlag {
		self.flagBool = flag.Bool(self.FlagName, false, self.FlagComment)
	} else {
		self.flagStr = flag.String(self.FlagName, "", self.FlagComment)
	}

}

var generators []*Generator

func RegisterGenerator(genList ...*Generator) {

	for _, gen := range genList {
		gen.init()
	}

	generators = append(generators, genList...)
}

func errorCatcher(errFunc func(error)) {

	err := recover()

	switch err.(type) {
	// 运行时错误
	case interface {
		RuntimeError()
	}:

		// 继续外抛， 方便调试
		panic(err)

	case error:
		errFunc(err.(error))
	case nil:
	default:
		panic(err)
	}
}

func RunGenerator(dset *model.DescriptorSet) (retErr error) {

	defer errorCatcher(func(genErr error) {

		retErr = genErr

	})

	for _, gen := range generators {

		if err := gen.gen(dset); err != nil {
			return err
		}

	}

	return nil
}
