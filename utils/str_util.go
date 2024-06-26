package utils

import (
	"strconv"
	"strings"
)

// ToInt string -> int
func ToInt(val string) (int, error) {
	return strconv.Atoi(val)
}

// ToInt64 string -> int64
func ToInt64(val string) (int64, error) {
	return strconv.ParseInt(val, 10, 64)
}

// ToFloat string -> float64
func ToFloat(val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}

// Int64ToString int64 -> string
func Int64ToString(val int64) string {
	return strconv.FormatInt(val, 10)
}

// Float64ToString float64 -> string
func Float64ToString(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}

// Split string -> []string
func Split(val string, ch rune) []string {
	values := strings.FieldsFunc(val, func(r rune) bool {
		return r == ch
	})

	return values
}

// SplitToFloat string -> []float64
func SplitToFloat(val string, ch rune) []float64 {
	var ret []float64
	values := Split(val, ch)
	for _, val := range values {
		if v, err := strconv.ParseFloat(val, 64); err == nil {
			ret = append(ret, v)
		}
	}
	return ret
}

// SplitToInt string -> []int64
func SplitToInt(val string, ch rune) []int64 {
	var ret []int64
	values := Split(val, ch)
	for _, val := range values {
		if v, err := strconv.ParseInt(val, 10, 64); err == nil {
			ret = append(ret, v)
		}
	}
	return ret
}

// Trim string remove backspace
func Trim(val string) string {
	return strings.Trim(val, " ")
}

// ToBool string -> bool
func ToBool(val string) (bool, error) {
	return strconv.ParseBool(val)
}

// BoolToString bool -> string
func BoolToString(val bool) string {
	return strconv.FormatBool(val)
}

// IsEmpty  func
func IsEmpty(str string) (isEmpty bool) {
	return len(Trim(str)) == 0
}

// ArrayInt64ToString []int64 to string
func ArrayInt64ToString(elems []int64, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return Int64ToString(elems[0])
	}

	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(Int64ToString(elems[i]))
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(Int64ToString(elems[0]))
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(Int64ToString(s))
	}
	return b.String()
}

// ReplaceStrEmptyAndLine  Replace all line breaks and spaces
func ReplaceStrEmptyAndLine(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}
