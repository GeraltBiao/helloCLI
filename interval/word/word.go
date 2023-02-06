package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

// 下划线转大写驼峰
func UnderscoreToUpperCamelClass(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// 下划线转小写驼峰
func UnderscoreToLowerCamelClass(s string) string {
	s = UnderscoreToUpperCamelClass(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 驼峰转下划线
func CamelCaseToUnderscore(s string) string {
	var ret []rune
	for i, v := range s {
		if i == 0 {
			ret = append(ret, unicode.ToLower(v))
			continue
		}
		if unicode.IsUpper(v) {
			ret = append(ret, '_')
		}
		ret = append(ret, unicode.ToLower(v))
	}
	return string(ret)
}
