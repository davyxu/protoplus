using System;
using System.Collections.Generic;
using System.Diagnostics;
using Proto;
using ProtoPlus;

namespace Proto
{
    public partial class MySubType : IProtoStruct
    {
        public override bool Equals(object obj)
        {
            if (!(obj is MySubType o))
                return false;

            return EqualsType(o);
        }

        public bool EqualsType(MySubType other)
        {
            return Bool == other.Bool &&
                   Int32 == other.Int32 &&
                   UInt32 == other.UInt32 &&
                   Int64 == other.Int64 &&
                   UInt64 == other.UInt64 &&
                   Float32.Equals(other.Float32) &&
                   Float64.Equals(other.Float64) &&
                   Enum == other.Enum &&
                   string.Equals(Str, other.Str);
        }
      
    }
    
    public partial class MyType : IProtoStruct
    {    
        public override bool Equals(object obj)
        {
            if (!(obj is MyType o))
                return false;

            return EqualsType(o);
        }

        protected bool EqualsType(MyType other)
        {
            return Bool == other.Bool &&
                   Int32 == other.Int32 &&
                   UInt32 == other.UInt32 &&
                   Int64 == other.Int64 &&
                   UInt64 == other.UInt64 &&
                   Float32.Equals(other.Float32) &&
                   Float64.Equals(other.Float64) &&
                   string.Equals(Str, other.Str) &&
                   Enum == other.Enum &&
                   Equals(Struct, other.Struct);
        }

    }
}

namespace Example
{
    class Program
    {
        static MyType MakeMyType()
        {
            return new MyType
            {
                Bool = true,
                Int32 = 200,
                UInt32 = UInt32.MaxValue - 100,
                Int64 = -789,
                UInt64 = 1234567890123456,
                Str = "hello",
                Float32 = 3.14f,
                Float64 = double.MaxValue,
                Enum = MyEnum.Two,
                BoolSlice = new List<bool>{ true, false, true },
                Int32Slice = new List<int>{ 1, 2, 3, 4},
                UInt32Slice = new List<uint> { 100, 200, 300, 400 },
                Int64Slice = new List<long> { 10, 20, 30, 40 },
                UInt64Slice = new List<ulong> { 100, 200, 300, 400 },
                StrSlice = new List<string> { "genji", "dva", "bastion" },
                Float32Slice = new List<float> { 1.1f, 2.1f, 3.2f, 4.5f },
                Float64Slice= new List<double> { 1.1, 2.1, 3.2, 4.5 },
                BytesSlice = new byte[]{ 61, 234, 7 },
                EnumSlice = new List<MyEnum>{MyEnum.Two, MyEnum.Zero, MyEnum.One},

                Struct = new MySubType()
                {
                    Bool = true,
                    Int32 = 128,
                },
                StructSlice = new List<MySubType>
                {
                    new MySubType
                    {
                        Int32 = 100,
                        Str = "x100",
                    },
                    new MySubType
                    {
                        Int32 = 200,
                        Str = "x200",
                    }

                }


            };
        }

        static void TestFull()
        {
            byte[] data = new byte[256];
            var s = new OutputStream(data);

            var myType = MakeMyType();

            s.Marshal(myType);

            var s2 = new InputStream();
            s2.Init(data, 0, s.Position);

            var myType2 = InputStream.CreateStruct<MyType>();

            s2.Unmarshal(myType2);

            Debug.Assert(myType.Equals(myType2));
        }

        static void TestMessage()
        {
            var mm = new MessageMeta();
            MessageMetaRegister.RegisterGeneratedMeta(mm);
            var msg = mm.CreateMessageByID(33606);

            var meta = mm.GetMetaByType(msg.GetType());

            Debug.Assert(meta.ID == 33606);

            Debug.Assert(meta.SourcePeer == "client");
        }

        static void Main(string[] args)
        {
            TestFull();
            TestMessage();
        }
    }

}

