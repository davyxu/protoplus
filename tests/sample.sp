
// 添加好友
[MsgID:501 ]
struct FriendAddREQ
{
	FriendID 	int64
}


[AutoMsgID]
struct FriendAddACK
{
	Result 	int32				
}

struct FriendInfo 
{	
	FriendID 	int64
	FriendName 	string
}


[AutoMsgID]
struct ItemAddREQ
{
	
}

[AutoMsgID]
struct PartyAddREQ
{
	PartyID 	int64
}

struct PartyAddInfo 
{	
	ParytID 	int64
	PartyName 	string
}

// 加入帮会结果
[MsgID:511 ]
struct PartyAddACK
{
	Result 	int32				
	PartyID int64	// 帮会ID
	Info	PartyAddInfo
	
}
