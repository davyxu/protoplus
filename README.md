# protoplus
适用于游戏开发的协议格式、代码生成器及相关开发API

# Why ProtoPlus？
Protobuf作为行业通用的协议格式已经存在多年，各种历史遗留问题及语言设计上的缺陷对游戏开发阻碍非常大。

为了提高协议编写效率，增加协议使用灵活度，急需一套强于Protobuf的工具集提高开发效率和使用便捷性，ProtoPlus应运而生！

# 特性对比

___ | ProtoPlus | Protobuf
---|---|---
生成字段Tag | 自动/手动 | 手动
生成枚举值 | 自动/手动 | 手动
注释扩展 | API简单 | API复杂
字段扩展 | 专属API支持 | 手动解析注释
结构扩展 | 专属API支持 | 手动解析注释
消息ID生成 |  自动/手动 | 不支持
路由表生成 |  支持 | 不支持
扩展方式 | 描述文件输出JSON | 复杂的插件格式，调试复杂

注：

* ProtoPlus中的枚举值，不再需要为了兼容C++而必须加上前缀，保证全局枚举值唯一

* ProtoPlus二进制编码格式与Protobuf一致，方便调试、分析、优化

# 描述文件格式

```

enum Vocation {
	Monkey
	Monk
	Pig
}

struct PhoneNumber {

	number string

	type int32
}


struct Person {

	name string

	id  int32

	email string

	phone PhoneNumber

	voc Vocation
}

struct AddressBook {

	person []Person
}


```



# 协议类型及输出语言类型对应

描述 |ProtoPlus | Go | C# | Protobuf
---|---|---|---|---
32位整形| int32 | int32 | int | int32
64位整形 | int64|int64 | long | int64
无符号32位整形| uint32|uint32 | uint | uint32
无符号64位整形| uint64|uint64 | ulong|uint64
字符串| string|string|string | string
单精度浮点数| float32|float32 | float | float
双精度浮点数|float64|float64| double | double
二进制数据 | bytes | []byte | byte[] |repeated byte
枚举| enum | int32类型常量封装| enum | enum
布尔| bool | bool| bool | bool
结构体| struct | struct| class | message


# 编译

```
	go get -u -v github.com/davyxu/protoplus/cmd/protoplus
```

# 功能

## 输出ProtoPlus编码消息序列化的Go源码

命令行示例:
```bash
protoplus --ppgo_out=msg_gen.go --package=proto proto1.proto proto2.proto
```

参数说明:
* ppgo_out

    Go源码文件名
    
* package

    指定输出时的Go包名

## 输出消息绑定的Go源码

输出源码被引用时, 自动注册到cellnet的消息Registry中

命令行示例:
```bash
protoplus --ppgoreg_out=msg_gen.go --package=proto proto1.proto proto2.proto
```

参数说明:
* ppgoreg_out

  Go源码文件名

* package

  指定输出时的Go包名

* codec   
    生成消息注册的默认编码，如在消息中指定编码时，优先使用指定的编码
  
    
## 输出ProtoPlus编码的消息序列化C#源码

输出的C#源码, 需要配合[ProtoPlus C# SDK](https://github.com/davyxu/protoplus/tree/master/api/csharp/ProtoPlus) 使用

命令行示例:
```bash
protoplus --ppcs_out=MsgGen.cs --package=Proto proto1.proto proto2.proto
```

参数说明:
* ppcs_out

   C#源码文件名
    
* package

    指定输出时的C#命名空间
    
* classbase

    C#代码生成时，消息类默认基类名称, 默认基类为IProtoStruct


## 输出消息绑定的C#源码

输出的C#源码, 需要配合[ProtoPlus C# SDK](https://github.com/davyxu/protoplus/tree/master/api/csharp/ProtoPlus) 使用

命令行示例:
```bash
protoplus --ppcsreg_out=MsgGen.cs --package=Proto proto1.proto proto2.proto
```

参数说明:
* ppcsreg_out

  C#源码文件名

* package

  指定输出时的C#命名空间


## 输出Protobuf协议描述文件

输出的Protobuf协议描述文件,可使用protoc编译器编译

命令行示例:
```bash
protoplus --pbscheme_out=pb.proto --package=proto proto1.proto proto2.proto
```

参数说明:
* pbscheme_out

  生成protobuf 3.0协议文件
    
* package

    指定输出时的Protobuf包名


## 输出ProtoPlus描述文件

ProtoPlus协议描述文件可输出为JSON格式, 方便插件及工具链获取ProtoPlus的格式信息

命令行示例:
```bash
protoplus --ppscheme_out=pp.json proto1.proto proto2.proto
```

参数说明:
* ppscheme_out

  JSON格式的ProtoPlus协议描述文件, 格式参见 [协议描述定义](https://github.com/davyxu/protoplus/tree/master/model/descriptorset.go)

也可将描述文件的JSON直接输出,示例如下:

```bash
protoplus --ppscheme proto1.proto proto2.proto
```
## 输出路由配置

在结构体上添加MsgDir字段,标记结构体作为消息时的消息流转方向, 格式范例:
```protoplus
[MsgDir: "client -> game"]
struct LoginREQ {
}
```
路由信息将输出为JSON格式:
```json
{
  "Rule": [
     {
        "MsgName": "proto.LoginREQ",
        "SvcName": "game"
     }
  ]
}
```
### MsgDir格式说明

MsgDir的格式包含3个部分
```
    MsgDir: From -> Mid -> To
```

* From

    消息的发起方,路由一般忽略此部分信息
    
* Mid

    消息的中间处理, 一般为网关或者路由等
    
* To

    消息的目标送达点,消息处理最终方, 一般为某种消息服务, 路由表的SvcName字段读取该字段

命令行示例:
```bash
protoplus --route_out=route.json --package=proto proto1.proto proto2.proto
```

参数说明:
* route_out

  JSON格式的ProtoPlus路由配置, 格式参见 [路由定义](https://github.com/davyxu/protoplus/tree/master/model/route.go)

* package

    指定输出时消息的包名
    
也可将路由信息的JSON直接输出,示例如下:

```bash
protoplus --route --package=proto proto1.proto proto2.proto
```

# 备注

感觉不错请star, 谢谢!

提交bug及特性: https://github.com/davyxu/protoplus/issues
