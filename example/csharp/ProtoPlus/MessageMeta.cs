using System;
using System.Collections.Generic;

namespace ProtoPlus
{
    public class MetaInfo
    {        
        public ushort ID;           // 消息ID
        public Type Type;           // 消息类型

        // 消息方向
        // 在proto中添加[MsgDir: "client -> game" ], 左边为源, 右边为目标
        public string SourcePeer;   // 消息发起的源
        public string TargetPeer;   // 消息的目标
    }

    // 消息扩展信息集合
    public partial class MessageMeta
    {
        private Dictionary<ushort, MetaInfo> metaByID = new Dictionary<ushort, MetaInfo>();
        private Dictionary<Type, MetaInfo> metaByType = new Dictionary<Type, MetaInfo>();

        // 注册消息的扩展信息
        public void RegisterMeta(MetaInfo info)
        {
            metaByID.Add(info.ID, info);
            metaByType.Add(info.Type, info);
        }

        // 通过ID取信息
        public MetaInfo GetMetaByID(ushort msgid)
        {
            if (metaByID.TryGetValue(msgid, out var value))
            {
                return value;
            }

            return null;
        }

        // 通过类型取信息
        public MetaInfo GetMetaByType(Type t)
        {
            if (metaByType.TryGetValue(t, out var value))
            {
                return value;
            }

            return null;
        }

        // 用类型取消息ID
        public ushort GetMsgIDByType(Type t)
        {
            var meta = GetMetaByType(t);
            if (meta == null)
                return 0;

            return meta.ID;
        }

        // 通过消息ID创建消息
        public IProtoStruct CreateMessageByID(ushort msgid)
        {
            var meta = GetMetaByID(msgid);
            if (meta == null)
                return null;

            return NewStruct(meta.Type);            
        }

        // 通过类型创建消息
        public static IProtoStruct NewStruct(Type t)
        {
            var s = Activator.CreateInstance(t) as IProtoStruct;
            if (s == null)
            {
                return null;
            }

            s.Init();

            return s;
        }


    }
}
