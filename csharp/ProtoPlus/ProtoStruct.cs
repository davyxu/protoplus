
namespace ProtoPlus
{
    public interface IProtoStruct
    {
        void Init();

        bool Unmarshal(InputStream stream, int fieldNumber, WireFormat.WireType wt );

        void Marshal(OutputStream stream);

        int GetSize();
    }
}
