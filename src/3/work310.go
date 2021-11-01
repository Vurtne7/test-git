package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma3("123.1231.123"))
	fmt.Println(comma3("123"))
	fmt.Println(comma3("123.456"))
	fmt.Println(comma3("1234.456"))
	fmt.Println(comma3("12345.456"))
	fmt.Println(comma3("123456.456"))
	fmt.Println(comma3("1234567.456"))
	fmt.Println(comma3("+123456789.456"))
	fmt.Println(comma3("-123456789.456"))
	fmt.Println(comma3("-12345678.456"))
	fmt.Println(comma3("-1234567.456"))
	fmt.Println(comma3("-123456.456"))
	fmt.Println(comma3("-12345.456"))
}

func comma3(s string) string {
	var buff bytes.Buffer
	var point string

	pointCount := strings.Count(s, ".")
	if pointCount > 1 {
		return s
	}
	if pointCount == 1 {
		strSlices := strings.Split(s, ".")
		s = strSlices[0]
		point = "." + strSlices[1]
	}

	// 写入正负数符号
	if strings.HasPrefix(s, "+") {
		s = s[1:]
		buff.WriteString("+")
	}
	if strings.HasPrefix(s, "-") {
		s = s[1:]
		buff.WriteString("-")
	}

	sLen := len(s)
	if sLen <= 3 {
		buff.WriteString(s)
		buff.WriteString(point)
		return buff.String()
	}

	var strSlices2 []string
	for i := sLen; i >= 0; i -= 3 {
		if i >= 3 {
			strSlices2 = append(strSlices2, s[i-3:i])
		} else if i > 0 {
			strSlices2 = append(strSlices2, s[0:i])
		}
	}
	for i := len(strSlices2) - 1; i >= 0; i-- {
		buff.WriteString(strSlices2[i])
		if i > 0 {
			buff.WriteString(",")
		}
	}

	// 写入小数点
	buff.WriteString(point)
	return buff.String()
}
