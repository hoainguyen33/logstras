package log

import "reflect"

var (
	TagName = "log"
)

func NameTag(field reflect.StructField) (string, bool) {
	tag := field.Tag.Get("log")
	if tag == "" {
		return field.Name, true
	}
	if tag == "-" {
		return "", false
	}
	return tag, true
}
