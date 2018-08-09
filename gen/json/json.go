package json

import (
	"encoding/json"
	"fmt"
	"github.com/davyxu/protoplus/gen"
	"io/ioutil"
)

// 输出到文件
func GenJson(ctx *gen.Context) error {

	data, err := json.MarshalIndent(ctx.DescriptorSet, "", "\t")

	if err != nil {
		return err
	}

	return ioutil.WriteFile(ctx.OutputFileName, data, 0666)
}

// 将json输出到标准输出
func OutputJson(ctx *gen.Context) error {

	data, err := json.MarshalIndent(ctx.DescriptorSet, "", "\t")

	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}
