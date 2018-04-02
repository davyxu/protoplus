package msgidutil

import "flag"

var (
	flagAutoMsgIDCacheFile            = flag.String("AutoMsgIDCacheFile", "", "Specifies auto msgid cache file")
	flagShowOverWriteCacheFileWarning = flag.Bool("ShowOverWriteCacheFileWarning", false, "Show warning when over write auto msgid cahce file, default is false")
	flagSuggestMsgIDStart             = flag.Int("SuggestMsgIDStart", 0, "Suggest msgid start, default is 0")
	flagCheckDuplicateMsgID           = flag.Bool("CheckDuplicateMsgID", false, "Check duplicate msgid, default is false")

	// 消息号发生重复时，添加一个自定义的字符串可以错开
	flagMsgIDSalt = flag.String("MsgIDSalt", "", "extra string add to input msgname in order to prevent hash collide")
)
