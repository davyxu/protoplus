package text

import (
	"github.com/davyxu/ulexer"
	"reflect"
)

func detectEnd(lex *ulexer.Lexer, literal string) bool {
	if literal != "" {
		lex.Read(ulexer.WhiteSpace())
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

		fieldTk := ulexer.NextToken(lex, ulexer.Identifier())

		tField, fieldExists := tObj.FieldByName(fieldTk.String())
		vField := vObj.FieldByName(fieldTk.String())

		ulexer.NextToken(lex, ulexer.Contain(":"))

		lex.Read(ulexer.WhiteSpace())

		parseValue(lex, func(tk *ulexer.Token) {
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

func parseValue(lex *ulexer.Lexer, action ulexer.MatchAction) {
	lex.SelectAction(
		[]ulexer.Matcher{ulexer.String(),
			ulexer.Numeral(),
			ulexer.Bool(),
			ulexer.Contain("["),
			ulexer.Contain("{"),
			ulexer.Identifier(),
		},
		[]ulexer.MatchAction{
			action,
			action,
			action,
			action,
			action,
			action,
		},
	)
}

func parseElement(lex *ulexer.Lexer, endLiteral string, onValue ulexer.MatchAction, onEnd func()) {
	for {

		if detectEnd(lex, endLiteral) {
			onEnd()
			break
		}

		parseValue(lex, onValue)
	}

}

func parseArray(lex *ulexer.Lexer, tField reflect.Type, vObj reflect.Value, endLiteral string) {

	switch tField.Elem().Kind() {
	case reflect.Int32:

		// 枚举
		if vObj.Kind() != reflect.Int32 {

			list := reflect.MakeSlice(vObj.Type(), 0, 0)

			parseElement(lex, endLiteral, func(tk *ulexer.Token) {

				vElm := reflect.ValueOf(tk.Int32())

				list = reflect.Append(list, vElm.Convert(vObj.Type().Elem()))
			}, func() {

				vObj.Set(list)
			})

		} else {
			var value []int32
			// TODO 按整数parse
			parseElement(lex, endLiteral, func(tk *ulexer.Token) {
				value = append(value, tk.Int32())
			}, func() {

				vObj.Set(reflect.ValueOf(value))
			})
		}

	case reflect.Int64:
		var value []int64
		// TODO 按整数parse
		parseElement(lex, endLiteral, func(tk *ulexer.Token) {
			value = append(value, tk.Int64())
		}, func() {
			vObj.Set(reflect.ValueOf(value))
		})
	case reflect.Uint64:
		var value []uint64
		// TODO 按整数parse
		parseElement(lex, endLiteral, func(tk *ulexer.Token) {
			value = append(value, tk.UInt64())
		}, func() {
			vObj.Set(reflect.ValueOf(value))
		})
	case reflect.Float32:
		var value []float32
		parseElement(lex, endLiteral, func(tk *ulexer.Token) {
			value = append(value, tk.Float32())
		}, func() {
			vObj.Set(reflect.ValueOf(value))
		})
	case reflect.Float64:
		var value []float64
		parseElement(lex, endLiteral, func(tk *ulexer.Token) {
			value = append(value, tk.Float64())
		}, func() {
			vObj.Set(reflect.ValueOf(value))
		})
	case reflect.Uint32:
		var value []uint32
		// TODO 按整数parse
		parseElement(lex, endLiteral, func(tk *ulexer.Token) {
			value = append(value, tk.UInt32())
		}, func() {
			vObj.Set(reflect.ValueOf(value))
		})
	case reflect.Bool:
		var value []bool
		parseElement(lex, endLiteral, func(tk *ulexer.Token) {
			value = append(value, tk.Bool())
		}, func() {
			vObj.Set(reflect.ValueOf(value))
		})
	case reflect.String:
		var value []string
		parseElement(lex, endLiteral, func(tk *ulexer.Token) {
			value = append(value, tk.String())
		}, func() {
			vObj.Set(reflect.ValueOf(value))
		})
	case reflect.Uint8:
		var value []byte
		// TODO 按整数parse
		parseElement(lex, endLiteral, func(tk *ulexer.Token) {
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

			lex.Read(ulexer.WhiteSpace())
			lex.Expect(ulexer.Contain("{"))

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

func UnmarshalText(s string, obj interface{}) error {

	vObj := safeValueOf(obj)
	tObj := safeTypeOf(obj)

	lex := ulexer.NewLexer(s)

	err := lex.Run(func(lex *ulexer.Lexer) {

		parseStruct(lex, tObj, vObj, "")

	})

	return err
}
