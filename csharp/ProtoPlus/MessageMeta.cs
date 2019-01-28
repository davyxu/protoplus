using System;
using System.Collections.Generic;

namespace ProtoPlus
{
    public class MetaInfo
    {
        public Func<IProtoStruct> CreateFunc;
        public ushort ID;
        public Type Type;
        public string EndPoint;
    }

    public partial class MessageMeta
    {
        private Dictionary<ushort, MetaInfo> metaByID = new Dictionary<ushort, MetaInfo>();
        private Dictionary<Type, MetaInfo> metaByType = new Dictionary<Type, MetaInfo>();

        public void RegisterMeta(MetaInfo info)
        {
            metaByID.Add(info.ID, info);
            metaByType.Add(info.Type, info);
        }

        public MetaInfo GetMetaByID(ushort msgid)
        {
            MetaInfo value;
            if (metaByID.TryGetValue(msgid, out value))
            {
                return value;
            }

            return null;
        }

        public MetaInfo GetMetaByType(Type t)
        {
            MetaInfo value;
            if (metaByType.TryGetValue(t, out value))
            {
                return value;
            }

            return null;
        }

        // 取消息由哪个端点发
        public bool GetMessageEndpoint(IProtoStruct msg, ref string index)
        {
            var meta = GetMetaByType(msg.GetType());
            if (meta == null)
                return false;

            index = meta.EndPoint;

            return true;
        }

        public IProtoStruct CreateMessageByID(ushort msgid)
        {
            var meta = GetMetaByID(msgid);
            if (meta == null)
                return null;

            return NewStruct(meta.Type, true);            
        }

        public static IProtoStruct NewStruct(Type t, bool initfield)
        {
            var s = Activator.CreateInstance(t) as IProtoStruct;
            if (s == null)
            {
                return null;
            }

            if (initfield)
            {
                s.Init();
            }

            return s;
        }

        public ushort GetMsgIDByType(Type t)
        {
            var meta = GetMetaByType(t);
            if (meta == null)
                return 0;

            return meta.ID;
        }
    }
}
