package route

import (
	"encoding/json"
	"fmt"
	"github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/gen"
	"github.com/davyxu/protoplus/model"
	"io/ioutil"
)

// 输出到文件
func GenJson(ctx *gen.Context) error {

	data, err := genJsonData(ctx)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(ctx.OutputFileName, data, 0666)
}

// 将json输出到标准输出
func OutputJson(ctx *gen.Context) error {

	data, err := genJsonData(ctx)

	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}

func genJsonData(ctx *gen.Context) ([]byte, error) {

	var rt model.RouteTable

	for _, d := range ctx.Structs() {

		msgDir := parseMessage(d)
		msgID := codegen.StructMsgID(d)

		if msgDir.Valid() {
			rt.Rule = append(rt.Rule, &model.RouteRule{
				MsgName: ctx.PackageName + "." + d.Name,
				SvcName: msgDir.To,
				Router:  msgDir.Mid,
				MsgID:   msgID,
			})
		}
	}

	data, err := json.MarshalIndent(&rt, "", "\t")

	if err != nil {
		return nil, err
	}

	return data, nil
}
