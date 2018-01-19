package msgidutil

import "flag"

var (
	flagAutoMsgIDCacheFile            = flag.String("AutoMsgIDCacheFile", "AutoMsgIDCacheFile.json", "Specifies auto msgid cache file")
	flagShowOverWriteCacheFileWarning = flag.Bool("ShowOverWriteCacheFileWarning", false, "Show warning when over write auto msgid cahce file, default is false")
	flagSuggestMsgIDStart             = flag.Int("SuggestMsgIDStart", 0, "Suggest msgid start, default is 0")
	flagCheckDuplicateMsgID           = flag.Bool("CheckDuplicateMsgID", false, "Check duplicate msgid, default is false")
)
