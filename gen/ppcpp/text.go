package ppcpp

const RegTemplateText = `// Generated by github.com/davyxu/protoplus
#pragma once
#include "MessageRegistry.h"
{{range SourceList $}}
#include "{{.}}"{{end}}

{{range .Structs}}
class {{.Name}}Meta : public MessageMeta
{
public:
	virtual const char* GetMessageName( ) const override{ return "{{$.PackageName}}.{{.Name}}"; }
	virtual int32 GetMessageId( ) const override{ return {{StructMsgID .}}; }
	virtual void* NewMessage() override { return {{$.PackageName}}::{{.Name}}::default_instance().New(); }
}; {{end}}

void StaticRegisterMeta( )
{ {{range .Structs}}
	MessageRegistry::Register(new {{.Name}}Meta); {{end}}
}

`
