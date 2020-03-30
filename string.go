package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
)

var (
	_              fmt.Formatter          = (*String)(nil)
	_              json.Marshaler         = (*String)(nil)
	_              encoding.TextMarshaler = (*String)(nil)
	FormatStringFn                        = func(s String, f fmt.State, c rune) {}
)

type String string

func (s String) Format(f fmt.State, c rune) {
	FormatStringFn(s, f, c)
}

func (s String) MarshalJSON() ([]byte, error) {
	var ss State
	s.Format(&ss, 's')
	return json.Marshal(string(ss.b))
}

func (s String) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 's')
	return ss.b, nil
}
