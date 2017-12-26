package parser

import (
	"encoding/json"
	"github.com/davyxu/golexer"
	"github.com/davyxu/protoplus/model"
	"strings"
	"testing"
)

func checkString(t *testing.T, script string) *model.DescriptorSet {

	if ds, err := ParseString(script); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		if data, err := json.Marshal(ds); err != nil {
			t.Error(err)
			t.FailNow()
		} else {
			t.Log(string(data))
		}

		return ds
	}

	return nil
}

func mustError(t *testing.T, script string, expectErrStr string) {

	if _, err := ParseString(script); err == nil || !strings.Contains(err.Error(), expectErrStr) {
		t.Log("[error not match]", err)
		t.FailNow()
	}
}

func TestDuplicateName(t *testing.T) {

	mustError(t, `
enum Vocation {

}
enum Vocation {

}
`, "Duplicate name: Vocation")

	mustError(t, `
struct Vocation {

}
struct Vocation {

}
`, "Duplicate name: Vocation")

}

func TestFieldTypeNotFound(t *testing.T) {

	mustError(t, `
struct Vocation {
	n Node
}

`, "type not found: Vocation")

}

func TestComment(t *testing.T) {

	ds := checkString(t, `

	// 枚举头注释
	enum Enum{

		// 枚举字段注释
		A // 枚举字段尾注释
	}

`)

	if ds.ObjectByName("Enum").Leading != "枚举头注释" {
		t.FailNow()
	}

	fd := ds.ObjectByName("Enum").FieldByName("A")
	if fd.Leading != "枚举字段注释" {
		t.FailNow()
	}

	if fd.Trailing != "枚举字段尾注释" {
		t.FailNow()
	}
}

func TestTags(t *testing.T) {

	ds := checkString(t, `

	[MsgID:201 CSRequestOnce:true Dir:"client->game"]
	struct s1 {

	}

	struct s2{
		[x: 1 y :2]
		f int32
	}
`)

	if ds.ObjectByName("s1").TagValueInt("MsgID") != 201 {
		t.FailNow()
	}

	if ds.ObjectByName("s1").TagValueBool("CSRequestOnce") != true {
		t.FailNow()
	}

	if ds.ObjectByName("s1").TagValueString("Dir") != "client->game" {
		t.FailNow()
	}

	fd := ds.ObjectByName("s2").FieldByName("f")

	if fd.TagValueInt("x") != 1 {
		t.FailNow()
	}

	if fd.TagValueInt("y") != 2 {
		t.FailNow()
	}
}

func TestParseString(t *testing.T) {

	checkString(t, `

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

struct Node{
	node []Node
}

	`)

}

func TestTagsNoColon(t *testing.T) {

	ds := checkString(t, `

	[MsgID:201 AutoMsgID  CSRequestOnce Dir:"client->game"]
	struct s1 {

	}

`)

	if ds.ObjectByName("s1").TagValueInt("MsgID") != 201 {
		t.FailNow()
	}

	if _, exists := ds.ObjectByName("s1").TagValueByKey("AutoMsgID"); exists != true {
		t.FailNow()
	}

	if _, exists := ds.ObjectByName("s1").TagValueByKey("CSRequestOnce"); exists != true {
		t.FailNow()
	}

	if ds.ObjectByName("s1").TagValueString("Dir") != "client->game" {
		t.FailNow()
	}

}

func sizeOfMustError(t *testing.T, script, name string, expectErrStr string) {

	ds := checkString(t, script)

	defer golexer.ErrorCatcher(func(err error) {

		if err == nil || !strings.Contains(err.Error(), expectErrStr) {
			t.Log("[error not match]", err)
			t.FailNow()
		}

	})

	ds.ObjectByName(name).Size()

}

func TestSizeOfNonsupport(t *testing.T) {

	sizeOfMustError(t, `

		struct PhoneNumber {

			number string

			type int32
		}
	`, "PhoneNumber", "Nonsupport string")

	sizeOfMustError(t, `

		struct PhoneNumber {

			number []int64

			type int32
		}
	`, "PhoneNumber", "Nonsupport repeated")

}

func TestSizeOf(t *testing.T) {

	ds := checkString(t, `

enum Sex
{
	Man
	Woman
}

struct PhoneNumber {

	number int64

	type int32
}


struct Person {

	id  int32

	phone PhoneNumber

	phone2 PhoneNumber

	sex Sex

}

`)

	if size := ds.ObjectByName("Person").Size(); size != 32 {
		t.FailNow()
	}

}
