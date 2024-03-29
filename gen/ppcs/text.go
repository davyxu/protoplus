package ppcs

// 报错行号+7
const TemplateText = `// Generated by github.com/davyxu/protoplus
// DO NOT EDIT!
using System;
using System.Collections.Generic;
using ProtoPlus;
#pragma warning disable 162

namespace {{.PackageName}}
{
	{{range $a, $enumobj := .Enums}}
	public enum {{.Name}} 
	{
		{{range .Fields}}
		{{.Name}} = {{PbTagNumber $enumobj .}}, {{end}}
	} {{end}}
	{{range $a, $obj := .Structs}}
	{{ObjectLeadingComment .}}
	public partial class {{$obj.Name}} : {{$.ClassBase}} 
	{
		{{range .Fields}}public {{CSTypeNameFull .}} {{.Name}};
		{{end}}
		#region Serialize Code
		public void Init( )
		{   {{range .Fields}}{{if NoneStructCanInit . }}
			{{.Name}} = new {{CSTypeNameFull .}}();	{{end}}{{end}}{{range .Fields}}{{if StructCanInit .}}
			{{.Name}} = ({{CSTypeNameFull .}}) MessageMeta.NewStruct(typeof({{CSTypeNameFull .}})); {{end}} {{end}}
		}

		public void Marshal(OutputStream stream)
		{ {{range .Fields}} 
			stream.Write{{CodecName .}}({{PbTagNumber $obj .}}, {{.Name}}); {{end}}
		}

		public int GetSize()
		{
			int size = 0; {{range .Fields}} 
			size += OutputStream.Size{{CodecName .}}({{PbTagNumber $obj .}}, {{.Name}}); {{end}}
			return size;
		}

 		public bool Unmarshal(InputStream stream, int fieldNumber, WireFormat.WireType wt)
		{
		 	switch (fieldNumber)
            { {{range .Fields}}
			case {{PbTagNumber $obj .}}:	
				stream.Read{{CodecName .}}(wt, ref {{.Name}});
                break; {{end}}
			default:
				return true;
            }

            return false;
		}
		#endregion
	}
{{end}}
}
`

const RegTemplateText = `// Generated by github.com/davyxu/protoplus
// DO NOT EDIT!
using System;
using System.Collections.Generic;
#pragma warning disable 162

namespace {{.PackageName}}
{
	public class MetaInfo
	{        
		public ushort ID;           // 消息ID
		public Type Type;           // 消息类型

		// 消息方向
		// 在proto中添加[MsgDir: "client -> game" ], 左边为源, 右边为目标
		public string SourcePeer;   // 消息发起的源
		public string TargetPeer;   // 消息的目标

		public string Name;
	}

	public static class MessageVisitor
    {
		public static void Visit(Action<MetaInfo> callback)
		{	{{range .Structs}}
            callback(new MetaInfo
            {
				Type = typeof({{.Name}}),	
				ID = {{StructMsgID .}}, 	
				SourcePeer = "{{GetSourcePeer .}}",
				TargetPeer = "{{GetTargetPeer .}}",
            });{{end}}
		}
    }
}
`
