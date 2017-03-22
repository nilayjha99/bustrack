package tools

import (
	"fmt"
	"strings"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func UpdateBuilder(tablename string, attr map[string]string, col string, val string) string {
	var base string
	base = fmt.Sprintf("update %s set ", tablename)

	for k, v := range attr {
		base += fmt.Sprintf("%s = %s, ", k, v)
		fmt.Println(base)
	}

	base = strings.TrimSuffix(base, ", ")
	base += fmt.Sprintf(" where %s = %s", col, val)
	return base
}
