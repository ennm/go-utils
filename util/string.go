package util

import "strings"

func Trim(str string) string {

    str = strings.TrimLeft(str, "(")

    str = strings.TrimRight(str, ")")

    return str
}

func SubStr(str string, start, end int64) string {

    return string([]rune(str)[start:end])
}

func UcFirst(str string) string {

    first := SubStr(str, 0, 1)

    return strings.Replace(str, first, strings.ToUpper(first), 1)
}
