package util

import (
	"fmt"
	"strconv"
	"strings"
)

func ToString(value interface{}, defaultValue string) string {
	str := strings.TrimSpace(fmt.Sprintf("%v", value))
	if str == "" {
		return defaultValue
	} else {
		return str
	}
}

func ToStringTrim(value interface{}, defaultValue string) string {
	s := fmt.Sprintf("%v", value)
	s = s[1 : len(s)-1]
	str := strings.TrimSpace(s)
	if str == "" {
		return defaultValue
	} else {
		return str
	}
}

func FromStringToUint32(value string) (uint32, error) {
	number, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint32(number), nil
}
