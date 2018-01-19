package msgidutil

import (
	"fmt"
	"github.com/davyxu/protoplus/model"
)

const sugguestMsgIDInterval = 100

func GenSuggestMsgID(dset *model.DescriptorSet) {

	// 段: MsgID/100

	sectionMap := make(map[int]bool)

	for _, d := range dset.Objects {

		if d.Kind != model.Kind_Struct {
			continue
		}
		userMsgID := d.TagValueInt("MsgID")

		if userMsgID == 0 {
			continue
		}

		sectionMap[userMsgID/sugguestMsgIDInterval] = true
	}

	var section = *flagSuggestMsgIDStart / sugguestMsgIDInterval

	for ; ; section++ {

		if _, ok := sectionMap[section]; ok {
			continue
		}

		fmt.Println("Suggest msgid:", section*sugguestMsgIDInterval+1)

		return
	}

}
