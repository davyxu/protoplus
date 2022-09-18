using System;
using System.Collections.Generic;
using System.Text;

namespace ProtoPlus
{
    public partial class OutputStream
    {
        #region Proto API


        public static int SizeBytes(int fieldIndex, byte[] value)
        {
            if (value == null || value.Length == 0)
            {
                return 0;
            }

            int size = value.Length;

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Bytes)) +
                   GetVarint32Size((uint)size) +
                   size;

        }

        public static int SizeBool(int fieldIndex, bool value)
        {
            if (value == false)
            {
                return 0;
            }

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Varint)) + 1;
        }

        public static int SizeBool(int fieldIndex, List<bool> value)
        {
            if (value == null || value.Count == 0)
            {
                return 0;
            }

            int size = value.Count;

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Bytes)) + 
                   GetVarint32Size((uint)size) +
                   size;
        }

        public static int SizeEnum<T>(int fieldIndex, T value) where T : struct, IConvertible
        {
            var v = Convert.ToInt32(value);
            return SizeInt32(fieldIndex, v);
        }

        public static int SizeEnum<T>(int fieldIndex, List<T> value) where T : struct, IConvertible
        {
            if (value == null || value.Count == 0)
            {
                return 0;
            }

            int size = 0;
            for (int i = 0; i < value.Count; i++)
            {
                size += GetVarint32Size((uint)Convert.ToInt32(value[i]));
            }

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Bytes)) +
                   GetVarint32Size((uint)size) +
                   size;
        }


        public static int SizeFloat(int fieldIndex, float value)
        {
            if (value == 0F)
            {
                return 0;
            }

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Fixed32)) + 4;
        }

        public static int SizeFloat(int fieldIndex, List<float> value)
        {
            if (value == null || value.Count == 0)
            {
                return 0;
            }

            int size = value.Count * 4;

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Bytes)) + 
                   GetVarint32Size((uint)size) +
                   size;
        }


        public static int SizeDouble(int fieldIndex, double value)
        {
            if (value == 0D)
            {
                return 0;
            }

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Fixed64)) + 8;
        }

        public static int SizeDouble(int fieldIndex, List<double> value)
        {
            if (value == null || value.Count == 0)
            {
                return 0;
            }

            int size = value.Count * 8;

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Bytes)) +
                   GetVarint32Size((uint)size) +
                   size;
        }

        public static int SizeInt32(int fieldIndex, int value)
        {
            if (value == 0)
                return 0;

            if (value > 0)
                return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Varint)) +
                       GetVarint32Size((uint)value);

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Fixed32)) + 4;
        }

        public static int SizeInt32(int fieldIndex, List<int> value)
        {
            if (value == null || value.Count == 0)
            {
                return 0;
            }

            int size = 0;
            for (int i = 0; i < value.Count; i++)
            {
                size += GetVarint32Size((uint)value[i]);
            }

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Bytes)) +
                   GetVarint32Size((uint)size) +
                   size;
        }

        public static int SizeUInt32(int fieldIndex, uint value)
        {
            if (value == 0)
                return 0;

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Varint)) +
                   GetVarint32Size(value);
        }


        public static int SizeUInt32(int fieldIndex, List<uint> value)
        {
            if (value == null || value.Count == 0)
            {
                return 0;
            }

            int size = 0;
            for (int i = 0; i < value.Count; i++)
            {
                size += GetVarint32Size(value[i]);
            }

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Bytes)) +
                   GetVarint32Size((uint)size) +
                   size;
        }


        public static int SizeInt64(int fieldIndex, long value)
        {
            if (value == 0)
                return 0;

            if (value > 0 )
                return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Varint)) + 
                       ComputeRawVarint64Size((ulong)value);

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Fixed64)) + 8;
        }

        public static int SizeInt64(int fieldIndex, List<long> value)
        {
            if (value == null || value.Count == 0)
            {
                return 0;
            }

            int size = 0;
            for (int i = 0; i < value.Count; i++)
            {
                size += ComputeRawVarint64Size((ulong)value[i]);
            }

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Bytes)) +
                   GetVarint32Size((uint)size) +
                   size;
        }

        public static int SizeUInt64(int fieldIndex, ulong value)
        {
            if (value == 0)
                return 0;

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Varint)) + 
                   ComputeRawVarint64Size(value);
        }

        public static int SizeUInt64(int fieldIndex, List<ulong> value)
        {
            if (value == null || value.Count == 0)
            {
                return 0;
            }

            int size = 0;
            for (int i = 0; i < value.Count; i++)
            {
                size += ComputeRawVarint64Size(value[i]);
            }

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Bytes)) +
                   GetVarint32Size((uint)size) +
                   size;
        }


        public static int SizeString(int fieldIndex, string value)
        {
            if (string.IsNullOrEmpty(value))
                return 0;

            int strLen = Encoding.UTF8.GetByteCount(value);

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Varint)) + 
                   GetVarint32Size((uint) strLen) + 
                   strLen;
        }

        public static int SizeString(int fieldIndex, List<string> value)
        {
            if (value == null || value.Count == 0)
            {
                return 0;
            }

            int size = 0;
            for (int i = 0; i < value.Count; i++)
            {
                size += SizeString(fieldIndex,value[i]);
            }

            return size;
        }

        public static int SizeStruct(int fieldIndex, IProtoStruct value)
        {
            if (value == null)
                return 0;

            var size = value.GetSize();
            if (size == 0)
            {
                return 0;
            }

            return GetVarint32Size(WireFormat.MakeTag(fieldIndex, WireFormat.WireType.Varint)) + 
                   GetVarint32Size((uint)size) + 
                   size;
        }

        public static int SizeStruct<T>(int fieldIndex, List<T> value) where T:IProtoStruct
        {
            if (value == null || value.Count == 0)
                return 0;

            int size = 0;
            for (int i = 0; i < value.Count; i++)
            {
                size += SizeStruct(fieldIndex, value[i]);
            }

            return size;
        }

        #endregion

        static int GetVarint32Size(uint value)
        {
            if ((value & (0xffffffff << 7)) == 0)
            {
                return 1;
            }
            if ((value & (0xffffffff << 14)) == 0)
            {
                return 2;
            }
            if ((value & (0xffffffff << 21)) == 0)
            {
                return 3;
            }
            if ((value & (0xffffffff << 28)) == 0)
            {
                return 4;
            }
            return 5;
        }

        static int ComputeRawVarint64Size(ulong value)
        {
            if ((value & (0xffffffffffffffffL << 7)) == 0)
            {
                return 1;
            }
            if ((value & (0xffffffffffffffffL << 14)) == 0)
            {
                return 2;
            }
            if ((value & (0xffffffffffffffffL << 21)) == 0)
            {
                return 3;
            }
            if ((value & (0xffffffffffffffffL << 28)) == 0)
            {
                return 4;
            }
            if ((value & (0xffffffffffffffffL << 35)) == 0)
            {
                return 5;
            }
            if ((value & (0xffffffffffffffffL << 42)) == 0)
            {
                return 6;
            }
            if ((value & (0xffffffffffffffffL << 49)) == 0)
            {
                return 7;
            }
            if ((value & (0xffffffffffffffffL << 56)) == 0)
            {
                return 8;
            }
            if ((value & (0xffffffffffffffffL << 63)) == 0)
            {
                return 9;
            }
            return 10;
        }


       


    }
}
