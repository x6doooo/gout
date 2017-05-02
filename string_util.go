package gout

import (
    "regexp"
    "strings"
)

var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

func CamelToSnake(s string) string {
    var a []string
    for _, sub := range camel.FindAllStringSubmatch(s, -1) {
        if sub[1] != "" {
            a = append(a, sub[1])
        }
        if sub[2] != "" {
            a = append(a, sub[2])
        }
    }
    return strings.ToLower(strings.Join(a, "_"))
}

func SnakeToCamel(str string) string {
    arr := strings.Split(str, "_")
    for i, v := range arr {
        arr[i] = strings.Title(v)
    }
    return strings.Join(arr, "")
}
