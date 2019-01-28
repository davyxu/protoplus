using System;

namespace ProtoPlus
{
    public static class WireFormat
    {
        public enum WireType : uint
        {
            None = 0,
            Varint,
            Bytes,
            Zigzag32,
            Zigzag64,
            Fixed32,
            Fixed64,
        }

        const int TagTypeBits = 3;
        const uint TagTypeMask = (1 << TagTypeBits) - 1;

        public static uint MakeTag(int fieldNumber, WireType type )
        {
            return (uint) (fieldNumber << TagTypeBits) | (uint)type;
        }

        public static void ParseWireTag(uint tag, ref int fieldNumber, ref WireType wt)
        {
            fieldNumber = (int) tag >> TagTypeBits;
            wt = (WireType) (tag & TagTypeMask);
        }

        const int CopyThreshold = 12;
        internal static void CopyBytes(byte[] src, byte[] dst, int srcOffset, int destOffset,  int count)
        {
            if (count < CopyThreshold)
            {
                for (int i = 0; i < count; i++)
                {
                    dst[destOffset+i] = src[srcOffset+i];
                }
            }
            else
            {
                Buffer.BlockCopy(src, srcOffset, dst, destOffset, count);
            }
        }
    }
}
