package json

import (
	"encoding/json"
	"fmt"
	"github.com/davyxu/protoplus/gen"
	"github.com/davyxu/protoplus/msgidutil"
	"io/ioutil"
	"strconv"
)

func genJsonData(ctx *gen.Context) (error, []byte) {

	for _, obj := range ctx.DescriptorSet.Objects {
		if obj.TagExists("AutoMsgID") {
			obj.SetTagValue("AutoMsgID", strconv.Itoa(msgidutil.StructMsgID(obj)))
		}

	}

	data, err := json.MarshalIndent(ctx.DescriptorSet, "", "\t")

	if err != nil {
		return err, nil
	}

	return nil, data
}

// 输出到文件
func GenJson(ctx *gen.Context) error {

	err, data := genJsonData(ctx)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(ctx.OutputFileName, data, 0666)
}

// 将json输出到标准输出
func OutputJson(ctx *gen.Context) error {

	err, data := genJsonData(ctx)

	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}
