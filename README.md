# protoplus
通用的描述文件及代码生成器及工具集


# 特性

* 提供比Protobuf描述文件更加友好的格式

* 支持插件架构，可以使用任何语言开发自己的协议runtime和代码生成器

* 插件信息交换使用JSON格式(Protobuf使用pb二进制格式)


# 描述文件格式(*.pp)

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

## 特性

* 自动生成tag序列号(base0),也可以手动指定

* 自动生成枚举序号(base0),也可以手动指定

* 类go结构体字段命名方式

* 比Protobuf更方便的导出注释内容做协议扩充

# 支持类型

* int32: 32位整形
* int64: 64位整形
* uint32: 无符号32位整形
* uint64: 无符号64位整形
* string: 字符串
* float32: 单精度浮点数
* float64: 双精度浮点数
* bytes: 二进制数据
* enum: int32封装
* bool: 布尔
* struct 结构体

所有类型前添加[]表示数组

# 编译

```
	go get -u -v github.com/davyxu/proto++
```

# 下载


# 命令行参数


# 使用方法


# 例子

https://github.com/davyxu/proto++/tree/master/example


# 备注

感觉不错请star, 谢谢!

博客: http://www.cppblog.com/sunicdavy

知乎: http://www.zhihu.com/people/sunicdavy

提交bug及特性: https://github.com/davyxu/proto++/issues
