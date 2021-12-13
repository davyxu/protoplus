package ppgo

import (
	"github.com/davyxu/ulexer"
	"reflect"
	"strings"
)

type MatchAction func(tk *ulexer.Token)

func selectAction(self *ulexer.Lexer, mlist []ulexer.Matcher, alist []MatchAction) {

	if len(mlist) != len(alist) {
		panic("Matcher list should equal to Action list length")
	}

	var hit bool
	for index, m := range mlist {
		tk := self.Read(m)

		if tk != nil {

			action := alist[index]
			if action != nil {
				action(tk)
			}
			hit = true
			break
		}
	}

	if !hit {

		var sb strings.Builder

		for index, m := range mlist {
			if index > 0 {
				sb.WriteString(" ")
			}
			sb.WriteString(m.TokenType())
		}

		self.Error("Expect %s", sb.String())
	}

}

func detectEnd(lex *ulexer.Lexer, literal string) bool {
	if literal != "" {
		if tk := lex.Read(ulexer.Contain(literal)); tk.String() == literal {
			return true
		}
	}

	return false
}

func parseStruct(lex *ulexer.Lexer, tObj reflect.Type, vObj reflect.Value, endLiteral string) {
	for {

		if detectEnd(lex, endLiteral) {
			break
		}

		fieldTk := ulexer.Expect(lex, ulexer.Identifier())

		tField, fieldExists := tObj.FieldByName(fieldTk.String())
		vField := vObj.FieldByName(fieldTk.String())

		ulexer.Expect(lex, ulexer.Contain(":"))

		parseMultiValue(lex, []ulexer.Matcher{ulexer.String(),
			ulexer.Numeral(),
			ulexer.Bool(),
			ulexer.Contain("["),
			ulexer.Contain("{"),
			ulexer.Identifier(),
		}, func(tk *ulexer.Token) {
			if !fieldExists {
				return
			}

			switch tField.Type.Kind() {
			case reflect.Bool:
				vField.SetBool(tk.Bool())
			case reflect.String:
				vField.SetString(tk.String())
			case reflect.Int32:
				vField.SetInt(int64(tk.Int32()))
			case reflect.Int64:
				vField.SetInt(int64(tk.Int64()))
			case reflect.Uint32:
				vField.SetUint(uint64(tk.UInt32()))
			case reflect.Uint64:
				vField.SetUint(uint64(tk.UInt64()))
			case reflect.Float32:
				vField.SetFloat(float64(tk.Float32()))
			case reflect.Float64:
				vField.SetFloat(tk.Float64())
			case reflect.Slice:
				parseArray(lex, tField.Type, vField, "]")
			case reflect.Struct:
				parseStruct(lex, tField.Type, vField, "}")
			default:
				panic("unsupport field kind " + tField.Type.Kind().String())
			}
		})
	}
}

func parseMultiValue(lex *ulexer.Lexer, mlist []ulexer.Matcher, action MatchAction) {

	alist := make([]MatchAction, 0, len(mlist))
	for i := 0; i < len(mlist); i++ {
		alist = append(alist, action)
	}

	selectAction(lex,
		mlist,
		alist,
	)
}

func parseElement(lex *ulexer.Lexer, endLiteral string, m ulexer.Matcher, onValue MatchAction, onEnd func()) {
	for {

		if detectEnd(lex, endLiteral) {
			onEnd()
			break
		}

		tk := ulexer.Expect(lex, m)
		onValue(tk)
	}

}

func parseArray(lex *ulexer.Lexer, tField reflect.Type, vObj reflect.Value, endLiteral string) {

	switch tField.Elem().Kind() {
	case reflect.Int32:

		// 枚举
		if vObj.Kind() != reflect.Int32 {

			list := reflect.MakeSlice(vObj.Type(), 0, 0)

			parseElement(lex, endLiteral, ulexer.Integer(), func(tk *ulexer.Token) {

				vElm := reflect.ValueOf(tk.Int32())

				list = reflect.Append(list, vElm.Convert(vObj.Type().Elem()))
			}, func() {

				vObj.Set(list)
			})

		} else {
			var value []int32
			parseElement(lex, endLiteral, ulexer.Integer(), func(tk *ulexer.Token) {
				value = append(value, tk.Int32())
			}, func() {

				vObj.Set(reflect.ValueOf(value))
			})
		}

	case reflect.Int64:
		var value []int64
		parseElement(lex, endLiteral, ulexer.Integer(), func(tk *ulexer.Token) {
			value = append(value, tk.Int64())
		}, func() {
			vObj.Set(reflect.ValueOf(value))
		})
	case reflect.Uint64:
		var value []uint64
		parseElement(lex, endLiteral, ulexer.UInteger(), func(tk *ulexer.Token) {
			value = append(value, tk.UInt64())
		}, func() {
			vObj.Set(reflect.ValueOf(value))
		})
	case reflect.Float32:
		var value []float32
		parseElement(lex, endLiteral, ulexer.Numeral(), func(tk *ulexer.Token) {
			value = append(value, tk.Float32())
		}, func() {
			vObj.Set(reflect.ValueOf(value))
		})
	case reflect.Float64:
		var value []float64
		parseElement(lex, endLiteral, ulexer.Numeral(), func(tk *ulexer.Token) {
			value = append(value, tk.Float64())
		}, func() {
			vObj.Set(reflect.ValueOf(value))
		})
	case reflect.Uint32:
		var value []uint32
		parseElement(lex, endLiteral, ulexer.UInteger(), func(tk *ulexer.Token) {
			value = append(value, tk.UInt32())
		}, func() {
			vObj.Set(reflect.ValueOf(value))
		})
	case reflect.Bool:
		var value []bool
		parseElement(lex, endLiteral, ulexer.Bool(), func(tk *ulexer.Token) {
			value = append(value, tk.Bool())
		}, func() {
			vObj.Set(reflect.ValueOf(value))
		})
	case reflect.String:
		var value []string
		parseElement(lex, endLiteral, ulexer.String(), func(tk *ulexer.Token) {
			value = append(value, tk.String())
		}, func() {
			vObj.Set(reflect.ValueOf(value))
		})
	case reflect.Uint8:
		var value []byte
		parseElement(lex, endLiteral, ulexer.Numeral(), func(tk *ulexer.Token) {
			value = append(value, tk.UInt8())
		}, func() {
			vObj.SetBytes(value)
		})
	case reflect.Struct:

		list := reflect.MakeSlice(tField, 0, 0)
		for {

			if detectEnd(lex, endLiteral) {
				break
			}

			ulexer.Expect(lex, ulexer.Contain("{"))

			element := reflect.New(tField.Elem()).Elem()

			parseStruct(lex, tField.Elem(), element, "}")

			list = reflect.Append(list, element)

			vObj.Set(list)
		}

	default:
		panic("unsupport array element " + tField.Kind().String())
	}

}

func safeValueOf(obj interface{}) reflect.Value {
	vObj := reflect.ValueOf(obj)
	if vObj.Kind() == reflect.Ptr {
		return vObj.Elem()
	}

	return vObj
}

func safeTypeOf(obj interface{}) reflect.Type {
	tObj := reflect.TypeOf(obj)
	if tObj.Kind() == reflect.Ptr {
		return tObj.Elem()
	}

	return tObj
}
