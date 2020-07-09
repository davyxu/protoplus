package text

import (
	"github.com/davyxu/ulexer"
	"github.com/davyxu/ureflect"
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

func parseStruct(lex *ulexer.Lexer, tObj *ureflect.Type, vObj ureflect.Pointer, endLiteral string) {
	for {

		if detectEnd(lex, endLiteral) {
			break
		}

		fieldTk := ulexer.NextToken(lex, ulexer.Identifier())

		tField := tObj.FieldByName(fieldTk.String())

		//fmt.Println(tObj.Name(), tField.Name)

		if tField.Name == "BytesSlice" {
			tField.Name = tField.Name
		}

		ulexer.NextToken(lex, ulexer.Contain(":"))

		lex.Read(ulexer.WhiteSpace())

		parseValue(lex, func(tk *ulexer.Token) {
			if tField == nil {
				return
			}

			switch tField.Type().Kind() {
			case ureflect.Bool:
				tField.SetBool(vObj, tk.Bool())
			case ureflect.String:
				tField.SetString(vObj, tk.String())
			case ureflect.Int32:
				tField.SetInt32(vObj, tk.Int32())
			case ureflect.Int64:
				tField.SetInt64(vObj, tk.Int64())
			case ureflect.UInt32:
				tField.SetUInt32(vObj, tk.UInt32())
			case ureflect.UInt64:
				tField.SetUInt64(vObj, tk.UInt64())
			case ureflect.Float32:
				tField.SetFloat32(vObj, tk.Float32())
			case ureflect.Float64:
				tField.SetFloat64(vObj, tk.Float64())
			case ureflect.Slice:
				parseArray(lex, tField, vObj, "]")
			case ureflect.Struct:
				parseStruct(lex, tField.Type(), tField.Value(vObj), "}")
			default:
				panic("unsupport field kind " + tField.Type().Kind().String())
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

func parseArray(lex *ulexer.Lexer, tField *ureflect.Field, vObj ureflect.Pointer, endLiteral string) {

	switch tField.Type().Elem().Kind() {
	case ureflect.Int32:
		var value []int32
		// TODO 按整数parse
		parseElement(lex, endLiteral, func(tk *ulexer.Token) {
			value = append(value, tk.Int32())
		}, func() {
			tField.SetInt32Slice(vObj, value)
		})
	case ureflect.Int64:
		var value []int64
		// TODO 按整数parse
		parseElement(lex, endLiteral, func(tk *ulexer.Token) {
			value = append(value, tk.Int64())
		}, func() {
			tField.SetInt64Slice(vObj, value)
		})
	case ureflect.UInt64:
		var value []uint64
		// TODO 按整数parse
		parseElement(lex, endLiteral, func(tk *ulexer.Token) {
			value = append(value, tk.UInt64())
		}, func() {
			tField.SetUInt64Slice(vObj, value)
		})
	case ureflect.Float32:
		var value []float32
		parseElement(lex, endLiteral, func(tk *ulexer.Token) {
			value = append(value, tk.Float32())
		}, func() {
			tField.SetFloat32Slice(vObj, value)
		})
	case ureflect.Float64:
		var value []float64
		parseElement(lex, endLiteral, func(tk *ulexer.Token) {
			value = append(value, tk.Float64())
		}, func() {
			tField.SetFloat64Slice(vObj, value)
		})
	case ureflect.UInt32:
		var value []uint32
		// TODO 按整数parse
		parseElement(lex, endLiteral, func(tk *ulexer.Token) {
			value = append(value, tk.UInt32())
		}, func() {
			tField.SetUInt32Slice(vObj, value)
		})
	case ureflect.Bool:
		var value []bool
		parseElement(lex, endLiteral, func(tk *ulexer.Token) {
			value = append(value, tk.Bool())
		}, func() {
			tField.SetBoolSlice(vObj, value)
		})
	case ureflect.String:
		var value []string
		parseElement(lex, endLiteral, func(tk *ulexer.Token) {
			value = append(value, tk.String())
		}, func() {
			tField.SetStringSlice(vObj, value)
		})
	case ureflect.UInt8:
		var value []byte
		// TODO 按整数parse
		parseElement(lex, endLiteral, func(tk *ulexer.Token) {
			value = append(value, tk.UInt8())
		}, func() {
			tField.SetBytes(vObj, value)
		})
	case ureflect.Struct:

		list := tField.Type().New()
		for {

			if detectEnd(lex, endLiteral) {
				break
			}

			lex.Read(ulexer.WhiteSpace())
			lex.Expect(ulexer.Contain("{"))

			element := tField.Type().Elem().New()

			parseStruct(lex, tField.Type().Elem(), element, "}")

			tField.SetValue(vObj, ureflect.PointerOf(list))
		}

	default:
		panic("unsupport array element " + tField.Type().Kind().String())
	}

}

func UnmarshalText(s string, obj interface{}) error {

	vObj := ureflect.PointerOf(obj)
	tObj := ureflect.TypeOf(obj)

	lex := ulexer.NewLexer(s)

	err := lex.Run(func(lex *ulexer.Lexer) {

		parseStruct(lex, tObj, vObj, "")

	})

	return err
}
