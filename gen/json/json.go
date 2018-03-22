package json

import (
	"encoding/json"
	"fmt"
	"github.com/davyxu/protoplus/gen"
	"github.com/davyxu/protoplus/model"
	"io/ioutil"
)

func init() {

	// 输出到文件
	gen.RegisterGenerator(&gen.Generator{
		FlagName:    "json_out",
		FlagComment: "json output to file",
		//builtin:     true,
		GenFunc: func(dset *model.DescriptorSet, fileName string) error {

			data, err := json.Marshal(dset)

			if err != nil {
				return err
			}

			return ioutil.WriteFile(fileName, data, 0666)
		},
	})

	// 输出到标准输出
	gen.RegisterGenerator(&gen.Generator{
		FlagName:    "json",
		FlagComment: "json text output to std out",
		//builtin:     true,
		UseBoolFlag: true,
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
