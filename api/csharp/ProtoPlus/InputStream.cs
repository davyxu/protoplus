using System;
using System.Collections.Generic;
using System.Text;

namespace ProtoPlus
{
    public class InputStream
    {
        int _pos;
        byte[] _buffer;
        int _len;

        public int Position => _pos;

        public int SpaceLeft => _len - _pos;
        
        public static Func<Type, IProtoStruct> CreateStructFunc = new Func<Type, IProtoStruct>(CreateStruct);

        public static T CreateStruct<T>() where T: IProtoStruct
        {
            return (T)CreateStruct(typeof(T));
        }

        public static IProtoStruct CreateStruct(Type t)
        {
            return MessageMeta.NewStruct(t );            
        }        

        

        public InputStream()
        {

        }

        public InputStream(byte[] _buffer)
        {
            Init(_buffer, 0, _buffer.Length );
        }

        public void Init(byte[] buffer, int offset, int length )
        {
            _pos = offset;
            _buffer = buffer;
            _len = length;
        }

        InputStream NewLimitStream(int size)
        {
            var stream = new InputStream();
            stream.Init(this._buffer, Position, Position + size);            
            return stream;
        }

        void CheckBuffer(int requireSize)
        {
            if (_len - _pos < requireSize)
            {
                throw new Exception("No enough buffer");
            }
        }

        #region Proto API

        public void ReadBytes(WireFormat.WireType wt, ref byte[] value)
        {
            if (wt == WireFormat.WireType.Bytes)
            {
                var count = ReadVarint32();

                if (value == null)
                {
                    value = new byte[count];
                }

                WireFormat.CopyBytes(_buffer, value, _pos,0, (int)count);
                _pos += (int)count;
            }
            else
            {
                throw new Exception("Invalid Wireformat");
            }
        }

        public void ReadBool(WireFormat.WireType wt, ref bool value )
        {
            if (wt == WireFormat.WireType.Varint)
            {
                value = ReadVarint32() != 0;
            }
            else
            {
                throw new Exception("Invalid Wireformat");
            }
        }

        public void ReadBool(WireFormat.WireType wt, ref List<bool> value)
        {
            if (wt == WireFormat.WireType.Bytes)
            {
                var count = ReadVarint32();

                if (value == null)
                {
                    value = new List<bool>((int)count);
                }

                for (int i = 0; i < count; i++)
                {
                    value.Add(ReadRawByte() != 0);                    
                }
            }
            else
            {
                throw new Exception("Invalid Wireformat");
            }
        }

        public void ReadEnum<T>(WireFormat.WireType wt, ref T value) where T : struct, IConvertible
        {
            int evalue = 0;
            ReadInt32(wt, ref evalue);
            value = (T)Enum.ToObject(typeof(T), evalue);
        }

        public void ReadEnum<T>(WireFormat.WireType wt, ref List<T> value) where T : struct, IConvertible
        {
            if (wt == WireFormat.WireType.Bytes)
            {
                int size = (int)ReadVarint32();

                var stream = NewLimitStream(size);                

                if (value == null)
                {
                    value = new List<T>();
                }

                while (stream.SpaceLeft > 0)
                {
                    int element = 0;
                    stream.ReadInt32(WireFormat.WireType.Varint, ref element);                   
                    value.Add((T)Enum.ToObject(typeof(T), element));
                }

                _pos += size;
            }
            else
            {
                throw new Exception("Invalid Wireformat");
            }
        }

        public void ReadInt32(WireFormat.WireType wt, ref int value)
        {
            switch (wt)
            {
                case WireFormat.WireType.Varint:
                    value = (int) ReadVarint32();
                    break;
                case WireFormat.WireType.Fixed32:
                    value = (int) ReadFixed32();
                    break;
                default:
                    throw new Exception("Invalid Wireformat");
            }
        }
       

        public void ReadInt32(WireFormat.WireType wt, ref List<int> value)
        {
            if (wt == WireFormat.WireType.Bytes)
            {
                int size = (int)ReadVarint32();

                var stream = NewLimitStream(size);

                if (value == null)
                {
                    value = new List<int>();
                }
                
                while(stream.SpaceLeft > 0 )
                {
                    int element = 0;
                    stream.ReadInt32(WireFormat.WireType.Varint, ref element);
                    value.Add(element);                    
                }

                _pos += size;
            }
            else
            {
                throw new Exception("Invalid Wireformat");
            }
        }

        public void ReadUInt32(WireFormat.WireType wt, ref uint value)
        {
            switch (wt)
            {
                case WireFormat.WireType.Varint:
                    value = ReadVarint32();
                    break;                
                default:
                    throw new Exception("Invalid Wireformat");
            }
        }

        public void ReadUInt32(WireFormat.WireType wt, ref List<uint> value)
        {
            if (wt == WireFormat.WireType.Bytes)
            {
                int size = (int)ReadVarint32();

                var stream = NewLimitStream(size);

                if (value == null)
                {
                    value = new List<uint>();
                }

                while (stream.SpaceLeft > 0)
                {
                    uint element = 0;
                    stream.ReadUInt32(WireFormat.WireType.Varint, ref element);
                    value.Add(element);
                }

                _pos += size;
            }
            else
            {
                throw new Exception("Invalid Wireformat");
            }
        }


        public void ReadInt64(WireFormat.WireType wt, ref long value)
        {
            switch (wt)
            {
                case WireFormat.WireType.Varint:
                    value = (int)ReadVarint64();
                    break;
                case WireFormat.WireType.Fixed64:
                    value = (int)ReadFixed64();
                    break;
                default:
                    throw new Exception("Invalid Wireformat");
            }
        }

        public void ReadInt64(WireFormat.WireType wt, ref List<long> value)
        {
            if (wt == WireFormat.WireType.Bytes)
            {
                int size = (int)ReadVarint32();

                var stream = NewLimitStream(size);

                if (value == null)
                {
                    value = new List<long>();
                }

                while (stream.SpaceLeft > 0)
                {
                    long element = 0;
                    stream.ReadInt64(WireFormat.WireType.Varint, ref element);
                    value.Add(element);
                }

                _pos += size;
            }
            else
            {
                throw new Exception("Invalid Wireformat");
            }
        }

        public void ReadUInt64(WireFormat.WireType wt, ref ulong value)
        {
            switch (wt)
            {
                case WireFormat.WireType.Varint:
                    value = ReadVarint64();
                    break;
                default:
                    throw new Exception("Invalid Wireformat");
            }
        }

        public void ReadUInt64(WireFormat.WireType wt, ref List<ulong> value)
        {
            if (wt == WireFormat.WireType.Bytes)
            {
                int size = (int)ReadVarint32();

                var stream = NewLimitStream(size);

                if (value == null)
                {
                    value = new List<ulong>();
                }

                while (stream.SpaceLeft > 0)
                {
                    ulong element = 0;
                    stream.ReadUInt64(WireFormat.WireType.Varint, ref element);
                    value.Add(element);
                }

                _pos += size;
            }
            else
            {
                throw new Exception("Invalid Wireformat");
            }
        }

        public void ReadFloat(WireFormat.WireType wt, ref float value)
        {
            switch (wt)
            {
                case WireFormat.WireType.Fixed32:
                    value = BitConverter.ToSingle(_buffer, _pos);
                    _pos += 4;
                    break;
                default:
                    throw new Exception("Invalid Wireformat");
            }
        }

        public void ReadFloat(WireFormat.WireType wt, ref List<float> value)
        {
            if (wt == WireFormat.WireType.Bytes)
            {
                int size = (int)ReadVarint32();

                var stream = NewLimitStream(size);

                if (value == null)
                {
                    value = new List<float>(size/4);
                }
                
                while (stream.SpaceLeft > 0)
                {
                    float element = 0;
                    stream.ReadFloat(WireFormat.WireType.Fixed32, ref element);
                    value.Add(element);
                }

                _pos += size;
            }
            else
            {
                throw new Exception("Invalid Wireformat");
            }
        }

        public void ReadDouble(WireFormat.WireType wt, ref double value)
        {
            switch (wt)
            {
                case WireFormat.WireType.Fixed64:
                    value = BitConverter.ToDouble(_buffer, _pos);
                    _pos += 8;
                    break;
                default:
                    throw new Exception("Invalid Wireformat");
            }
        }

        public void ReadDouble(WireFormat.WireType wt, ref List<double> value)
        {
            if (wt == WireFormat.WireType.Bytes)
            {
                int size = (int)ReadVarint32();

                var stream = NewLimitStream(size);

                if (value == null)
                {
                    value = new List<double>(size / 8);
                }

                while (stream.SpaceLeft > 0)
                {
                    double element = 0;
                    stream.ReadDouble(WireFormat.WireType.Fixed64, ref element);
                    value.Add(element);
                }

                _pos += size;
            }
            else
            {
                throw new Exception("Invalid Wireformat");
            }
        }

        public void ReadString(WireFormat.WireType wt, ref string value)
        {
            switch (wt)
            {
                case WireFormat.WireType.Bytes:
                    int len = (int)ReadVarint32();
                    if (len > 0)
                    {
                        value = Encoding.UTF8.GetString(_buffer, _pos, len);
                        _pos += len;
                    }
                    break;
                default:
                    throw new Exception("Invalid Wireformat");
            }
        }

        public void ReadString(WireFormat.WireType wt, ref List<string> value)
        {
            if (wt == WireFormat.WireType.Bytes)
            {
                if (value == null)
                {
                    value = new List<string>();
                }

                string str = string.Empty;
                ReadString(wt, ref str);
                value.Add(str);
            }
            else
            {
                throw new Exception("Invalid Wireformat");
            }
        }

        public void ReadStruct<T>(WireFormat.WireType wt, ref List<T> value) where T : IProtoStruct, new()
        {
            if (wt == WireFormat.WireType.Bytes)
            {
                if (value == null)
                {
                    value = new List<T>();
                }

                T s = default(T);
                ReadStruct(wt, ref s);
                value.Add(s);
            }
            else
            {
                throw new Exception("Invalid Wireformat");
            }
        }

        public void ReadStruct<T>(WireFormat.WireType wt, ref T value) where T:IProtoStruct
        {
            switch (wt)
            {
                case WireFormat.WireType.Bytes:
                    int size = (int)ReadVarint32();
                    if (size > 0)
                    {
                        var stream = NewLimitStream(size);

                        if (value == null)
                        {
                            value = (T)CreateStruct(typeof(T));
                        }
                        stream.Unmarshal(value);
                        _pos += size;
                    }
                    break;
                default:
                    throw new Exception("Invalid Wireformat");
            }
        }

        #endregion

        public void Unmarshal(IProtoStruct s)
        {
            while(SpaceLeft > 0 )
            {
                var tag = ReadVarint32();

                int fieldNumber = -1;
                WireFormat.WireType type = WireFormat.WireType.None;
                WireFormat.ParseWireTag(tag, ref fieldNumber, ref type);

                if (s.Unmarshal(this, fieldNumber, type))
                {
                    SkipField(type);
                }
            }
        }

        void SkipField(WireFormat.WireType type )
        {
            switch (type)
            {
                case WireFormat.WireType.Varint:
                    ReadVarint32();
                    break;
                case WireFormat.WireType.Bytes:
                    var len = ReadVarint32();                    
                    _pos += (int)len;
                    break;
                case WireFormat.WireType.Fixed32:
                    ReadFixed32();
                    break;
                case WireFormat.WireType.Fixed64:
                    ReadFixed64();
                    break;
                default:
                    throw new Exception("Unknown wire type to skip");
            }
        }

        internal byte ReadRawByte()
        {
            if (_pos == _len)
            {
                throw new Exception("TruncatedMessage");
            }

            return _buffer[_pos++];
        }

        uint SlowReadRawVarint32()
        {
            int tmp = ReadRawByte();
            if (tmp < 128)
            {
                return (uint)tmp;
            }
            int result = tmp & 0x7f;
            if ((tmp = ReadRawByte()) < 128)
            {
                result |= tmp << 7;
            }
            else
            {
                result |= (tmp & 0x7f) << 7;
                if ((tmp = ReadRawByte()) < 128)
                {
                    result |= tmp << 14;
                }
                else
                {
                    result |= (tmp & 0x7f) << 14;
                    if ((tmp = ReadRawByte()) < 128)
                    {
                        result |= tmp << 21;
                    }
                    else
                    {
                        result |= (tmp & 0x7f) << 21;
                        result |= (tmp = ReadRawByte()) << 28;
                        if (tmp >= 128)
                        {
                            // Discard upper 32 bits.
                            for (int i = 0; i < 5; i++)
                            {
                                if (ReadRawByte() < 128)
                                {
                                    return (uint)result;
                                }
                            }
                            throw new Exception("MalformedVarint");
                        }
                    }
                }
            }
            return (uint)result;
        }

        uint ReadFixed32()
        {
            CheckBuffer(4);

            uint b1 = _buffer[_pos++];
            uint b2 = _buffer[_pos++];
            uint b3 = _buffer[_pos++];
            uint b4 = _buffer[_pos++];

            return b1 | (b2 << 8) | (b3 << 16) | (b4 << 24);
        }

        ulong ReadFixed64()
        {
            CheckBuffer(8);
            
            ulong b1 = _buffer[_pos++];
            ulong b2 = _buffer[_pos++];
            ulong b3 = _buffer[_pos++];
            ulong b4 = _buffer[_pos++];
            ulong b5 = _buffer[_pos++];
            ulong b6 = _buffer[_pos++];
            ulong b7 = _buffer[_pos++];
            ulong b8 = _buffer[_pos++];

            return b1 | (b2 << 8) | (b3 << 16) | (b4 << 24)
                   | (b5 << 32) | (b6 << 40) | (b7 << 48) | (b8 << 56);
        }

        ulong ReadVarint64()
        {
            int shift = 0;
            ulong result = 0;
            while (shift < 64)
            {
                byte b = ReadRawByte();
                result |= (ulong)(b & 0x7F) << shift;
                if ((b & 0x80) == 0)
                {
                    return result;
                }
                shift += 7;
            }

            throw new Exception("MalformedVarint");
        }

        uint ReadVarint32()
        {
            if (_pos + 5 > _len)
            {
                return SlowReadRawVarint32();
            }

            int tmp = _buffer[_pos++];
            if (tmp < 128)
            {
                return (uint)tmp;
            }
            int result = tmp & 0x7f;
            if ((tmp = _buffer[_pos++]) < 128)
            {
                result |= tmp << 7;
            }
            else
            {
                result |= (tmp & 0x7f) << 7;
                if ((tmp = _buffer[_pos++]) < 128)
                {
                    result |= tmp << 14;
                }
                else
                {
                    result |= (tmp & 0x7f) << 14;
                    if ((tmp = _buffer[_pos++]) < 128)
                    {
                        result |= tmp << 21;
                    }
                    else
                    {
                        result |= (tmp & 0x7f) << 21;
                        result |= (tmp = _buffer[_pos++]) << 28;
                        if (tmp >= 128)
                        {
                            // Discard upper 32 bits.
                            // Note that this has to use ReadRawByte() as we only ensure we've
                            // got at least 5 bytes at the start of the method. This lets us
                            // use the fast path in more cases, and we rarely hit this section of code.
                            for (int i = 0; i < 5; i++)
                            {
                                if (ReadRawByte() < 128)
                                {
                                    return (uint)result;
                                }
                            }
                            throw new Exception("MalformedVarint");
                        }
                    }
                }
            }
            return (uint)result;
        }
    }
}
