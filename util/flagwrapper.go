package util

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/davyxu/protoplus/model"
	"github.com/davyxu/protoplus/parser"
	"io/ioutil"
)

type GenFunc func(*model.DescriptorSet, string) error

type Generator struct {
	FlagName    string
	FlagComment string

	GenFunc GenFunc

	flagStr *string

	flagBool *bool

	builtin bool

	useBoolFlag bool
}

func (self *Generator) flagString() string {
	if self.flagStr == nil {
		return ""
	}

	return *self.flagStr
}

func (self *Generator) gen(dset *model.DescriptorSet) error {

	if self.useBoolFlag && *self.flagBool {
		return self.GenFunc(dset, "")
	} else if self.flagString() != "" {

		if !self.builtin {
			fmt.Printf("%s:\n", self.FlagName)
		}

		return self.GenFunc(dset, self.flagString())
	}

	return nil
}

func (self *Generator) init() {

	if self.useBoolFlag {
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

func init() {

	RegisterGenerator(&Generator{
		FlagName:    "json_out",
		FlagComment: "json output",
		builtin:     true,
		GenFunc: func(dset *model.DescriptorSet, fileName string) error {

			data, err := json.Marshal(dset)

			if err != nil {
				return err
			}

			return ioutil.WriteFile(fileName, data, 0666)
		},
	},
		&Generator{
			FlagName:    "json",
			FlagComment: "json text output to std out",
			builtin:     true,
			useBoolFlag: true,
			GenFunc: func(dset *model.DescriptorSet, fileName string) error {

				data, err := json.Marshal(dset)

				if err != nil {
					return err
				}

				fmt.Println(string(data))
				return nil
			},
		})
}

func RunGenerator() error {

	flag.Parse()

	var dset model.DescriptorSet

	err := parser.ParseFileList(&dset, flag.Args()...)
	if err != nil {
		return err
	}

	for _, gen := range generators {

		if err := gen.gen(&dset); err != nil {
			return err
		}

	}

	return nil
}
