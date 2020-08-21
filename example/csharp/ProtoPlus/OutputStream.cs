using System;
using System.Collections.Generic;
using System.Text;

namespace ProtoPlus
{
    public partial class OutputStream
    {
        int _pos;
        byte[] _buffer;
        int _len;        
        bool _extend;

        public int Position => _pos;

        public int SpaceLeft => _len - _pos;

        public OutputStream(byte[] buffer )
        {
            Init(buffer, 0, buffer.Length, false);
        }

        public void Init(byte[] buffer, int offset, int length, bool extend)
        {
            _pos = offset;
            _buffer = buffer;            
            _len = length;
            _extend = extend;
        }

        

        void CheckBuffer(int requireSize)
        {
            if (_len - _pos < requireSize)
            {
                if (_extend)
                {
                    var newdata = new byte[2 * _buffer.Length + requireSize];
                    WireFormat.CopyBytes(_buffer, newdata, 0, 0,  _buffer.Length);
                    _len = newdata.Length;
                    _buffer = newdata;                 
                }
                else
                {
                    throw new Exception("No enough buffer");
                }
            }            
        }

        void WriteTag(int fieldNumber, WireFormat.WireType type)
        {
            WriteVarint32(WireFormat.MakeTag(fieldNumber, type));
        }

        #region Proto API

        public void WriteBytes(int fieldIndex, byte[] value)
        {
            if (value == null || value.Length == 0)
            {
                return;
            }

            WriteTag(fieldIndex, WireFormat.WireType.Bytes);

            WriteVarint32((uint)value.Length);
            WriteRawBytes(value, value.Length);
        }


        public void WriteBool(int fieldIndex, bool value)
        {
            if (!value)
                return;

            WriteTag(fieldIndex, WireFormat.WireType.Varint);
            WriteByte(1);
        }

        public void WriteBool(int fieldIndex, List<bool> value)
        {
            if (value == null || value.Count == 0)
            {
                return;
            }

            WriteTag(fieldIndex, WireFormat.WireType.Bytes);

            WriteVarint32((uint)value.Count);

            for (int i = 0; i < value.Count; i++)
            {
                WriteByte(value[i] ? (byte)1: (byte)0);
            }
        }

        public void WriteEnum<T>(int fieldIndex, T value) where  T: struct, IConvertible
        {
            var v = Convert.ToInt32(value);
            WriteInt32(fieldIndex, v);
        }

        public void WriteEnum<T>(int fieldIndex, List<T> value) where T : struct, IConvertible
        {
            if (value == null || value.Count == 0)
            {
                return;
            }

            WriteTag(fieldIndex, WireFormat.WireType.Bytes);

            int size = 0;
            for (int i = 0; i < value.Count; i++)
            {
                size += GetVarint32Size((uint)Convert.ToInt32(value[i]));
            }

            WriteVarint32((uint)size);
            for (int i = 0; i < value.Count; i++)
            {
                WriteVarint32((uint)Convert.ToInt32(value[i]));
            }
        }

        public void WriteInt32(int fieldIndex, int value)
        {
            if (value == 0)
            {
                return;
            }

            if (value > 0)
            {
                WriteTag(fieldIndex, WireFormat.WireType.Varint);
                WriteVarint32((uint)value);
            }
            else
            {
                WriteTag(fieldIndex, WireFormat.WireType.Fixed32);
                WriteFixed32((uint)value);
            }
        }

        public void WriteInt32(int fieldIndex, List<int> value)
        {
            if (value == null || value.Count == 0)
            {
                return;
            }

            WriteTag(fieldIndex, WireFormat.WireType.Bytes);

            int size = 0;
            for (int i = 0; i < value.Count; i++)
            {
                size += GetVarint32Size((uint) value[i]);
            }

            WriteVarint32((uint) size);
            for (int i = 0; i < value.Count; i++)
            {
                WriteVarint32((uint)value[i]);
            }
        }

        public void WriteUInt32(int fieldIndex, uint value)
        {
            if (value == 0)
            {
                return;
            }

            WriteTag(fieldIndex, WireFormat.WireType.Varint);
            WriteVarint32(value);
        }

        public void WriteUInt32(int fieldIndex, List<uint> value)
        {
            if (value == null || value.Count == 0)
            {
                return;
            }

            WriteTag(fieldIndex, WireFormat.WireType.Bytes);

            int size = 0;
            for (int i = 0; i < value.Count; i++)
            {
                size += GetVarint32Size(value[i]);
            }

            WriteVarint32((uint)size);
            for (int i = 0; i < value.Count; i++)
            {
                WriteVarint32(value[i]);
            }
        }

        public void WriteInt64(int fieldIndex, long value)
        {
            if (value == 0)
            {
                return;
            }

            if (value > 0)
            {
                WriteTag(fieldIndex, WireFormat.WireType.Varint);
                WriteVarint64((ulong)value);
            }
            else
            {
                WriteTag(fieldIndex, WireFormat.WireType.Fixed64);
                WriteFixed64((ulong)value);
            }
        }

        public void WriteInt64(int fieldIndex,List<long> value)
        {
            if (value == null || value.Count == 0)
            {
                return;
            }

            WriteTag(fieldIndex, WireFormat.WireType.Bytes);

            int size = 0;
            for (int i = 0; i < value.Count; i++)
            {
                size += ComputeRawVarint64Size((ulong)value[i]);
            }

            WriteVarint32((uint)size);
            for (int i = 0; i < value.Count; i++)
            {
                WriteVarint64((ulong)value[i]);
            }
        }

        public void WriteUInt64(int fieldIndex, ulong value)
        {
            if (value == 0)
            {
                return;
            }

            WriteTag(fieldIndex, WireFormat.WireType.Varint);
            WriteVarint64(value);
        }

        public void WriteUInt64(int fieldIndex, List<ulong> value)
        {
            if (value == null || value.Count == 0)
            {
                return;
            }

            WriteTag(fieldIndex, WireFormat.WireType.Bytes);

            int size = 0;
            for (int i = 0; i < value.Count; i++)
            {
                size += ComputeRawVarint64Size(value[i]);
            }

            WriteVarint32((uint)size);
            for (int i = 0; i < value.Count; i++)
            {
                WriteVarint64(value[i]);
            }
        }

        public void WriteFloat(int fieldIndex, float value)
        {            
            if (value == 0F)
            {
                return;
            }

            WriteTag(fieldIndex, WireFormat.WireType.Fixed32);

            CheckBuffer(4);
            FastBitConverter.GetBytes(_buffer, _pos, value);
            _pos += 4; ;
        }

        public void WriteFloat(int fieldIndex, List<float> value)
        {
            if (value == null || value.Count == 0)
            {
                return;
            }

            WriteTag(fieldIndex, WireFormat.WireType.Bytes);
            WriteVarint32((uint)value.Count * 4);

            CheckBuffer(value.Count * 4);
            

            for (int i = 0; i < value.Count; i++)
            {
                FastBitConverter.GetBytes(_buffer, _pos, value[i]);
                _pos += 4;
            }
        }

        public void WriteDouble(int fieldIndex, double value)
        {
            if (value == 0D)
            {
                return;
            }

            WriteTag(fieldIndex, WireFormat.WireType.Fixed64);

            CheckBuffer(8);
            FastBitConverter.GetBytes(_buffer, _pos, value);
            _pos += 8;
        }


        public void WriteDouble(int fieldIndex, List<double> value)
        {
            if (value == null || value.Count == 0)
            {
                return;
            }

            WriteTag(fieldIndex, WireFormat.WireType.Bytes);
            WriteVarint32((uint)value.Count * 8);

            CheckBuffer(value.Count * 8);


            for (int i = 0; i < value.Count; i++)
            {
                FastBitConverter.GetBytes(_buffer, _pos, value[i]);
                _pos += 8;
            }
        }

        public void WriteString(int fieldIndex, string value)
        {
            if (String.IsNullOrEmpty(value))
            {
                return;
            }


            WriteTag(fieldIndex, WireFormat.WireType.Bytes);

            WriteStringBytes(value);
        }

        void WriteStringBytes(string value )
        {           
            int strLen = Encoding.UTF8.GetByteCount(value);

            WriteVarint32((uint)strLen);

            CheckBuffer(strLen);

            // 全ASCII
            if (strLen == value.Length)
            {
                for (int i = 0; i < strLen; i++)
                {
                    _buffer[_pos + i] = (byte)value[i];
                }
            }
            else
            {
                Encoding.UTF8.GetBytes(value, 0, value.Length, _buffer, _pos);

            }

            _pos += strLen;
        }

        public void WriteString(int fieldIndex, List<string> value)
        {
            if (value == null || value.Count == 0)
            {
                return;
            }

            for (int i = 0; i < value.Count; i++)
            {
                WriteTag(fieldIndex, WireFormat.WireType.Bytes);
                WriteStringBytes(value[i]);
            }
        }

        public void WriteStruct(int fieldIndex, IProtoStruct s)
        {            
            if (s == null)
            {
                return;
            }

            var size = s.GetSize();
            if (size == 0)
            {
                return;
            }

            WriteTag(fieldIndex, WireFormat.WireType.Bytes);

            WriteVarint32((uint)size);

            s.Marshal(this);
        }

        public void WriteStruct<T>(int fieldIndex, List<T> value) where T:IProtoStruct
        {
            if (value == null)
                return;

            for (int i = 0; i < value.Count; i++)
            {
                WriteStruct(fieldIndex, value[i]);
            }
        }

        #endregion

        public void Marshal(IProtoStruct s)
        {
            s.Marshal(this);
        }

        void WriteVarint32(uint value)
        {
            // Optimize for the common case of a single byte value
            if (value < 128 && _pos < _len)
            {
                _buffer[_pos++] = (byte)value;
                return;
            }

            while (value > 127 && _pos < _len)
            {
                _buffer[_pos++] = (byte)((value & 0x7F) | 0x80);
                value >>= 7;
            }

            while (value > 127)
            {
                WriteByte((byte)((value & 0x7F) | 0x80));
                value >>= 7;
            }

            if (_pos < _len)
            {
                _buffer[_pos++] = (byte)value;
            }
            else
            {
                WriteByte((byte) value);
            }            
        }

        internal void WriteByte(byte value )
        {
            CheckBuffer(1);
            _buffer[_pos++] = value;
        }

        internal void WriteVarint64(ulong value)
        {
            while (value > 127 && _pos < _len)
            {
                _buffer[_pos++] = (byte)((value & 0x7F) | 0x80);
                value >>= 7;
            }
            while (value > 127)
            {
                WriteByte((byte) ((value & 0x7F) | 0x80));
                value >>= 7;                
            }
            if (_pos < _len)
            {
                _buffer[_pos++] = (byte)value;
            }
            else
            {
                WriteByte((byte) value);
            }            
        }


        internal void WriteFixed32(uint value)
        {
            CheckBuffer(4);
            _buffer[_pos++] = ((byte)value);
            _buffer[_pos++] = ((byte)(value >> 8));
            _buffer[_pos++] = ((byte)(value >> 16));
            _buffer[_pos++] = ((byte)(value >> 24));
        }

        internal void WriteFixed64(ulong value)
        {
            CheckBuffer(8);
            _buffer[_pos++] = ((byte)value);
            _buffer[_pos++] = ((byte)(value >> 8));
            _buffer[_pos++] = ((byte)(value >> 16));
            _buffer[_pos++] = ((byte)(value >> 24));
            _buffer[_pos++] = ((byte)(value >> 32));
            _buffer[_pos++] = ((byte)(value >> 40));
            _buffer[_pos++] = ((byte)(value >> 48));
            _buffer[_pos++] = ((byte)(value >> 56));
        }

        internal void WriteRawBytes(byte[] value, int length )
        {
            CheckBuffer(length);
            WireFormat.CopyBytes(value, _buffer, 0, _pos, length);
            _pos += length;
        }


    }
}
