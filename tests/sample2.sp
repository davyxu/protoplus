
// 添加好友
[MsgID:601 ]
struct FriendAddREQ2
{
	FriendID 	int64
}


[AutoMsgID]
struct FriendAddACK2
{
	Result 	int32				
}

struct FriendInfo2 
{	
	FriendID 	int64
	FriendName 	string
}


[AutoMsgID]
struct ItemAddREQ2
{
	
}

[AutoMsgID]
struct PartyAddREQ2
{
	PartyID 	int64
}

struct PartyAddInfo2 
{	
	ParytID 	int64
	PartyName 	string
}

// 加入帮会结果
[MsgID:611 ]
struct PartyAddACK2
{
	Result 	int32				
	PartyID int64	// 帮会ID
	Info	PartyAddInfo
	
}
