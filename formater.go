package log

import (
	"bytes"
	"fmt"
	"reflect"
)

type Format int8

var (
	JSON   Format = 0
	STRING Format = 1
)

func (f Format) Format(buf *bytes.Buffer, level Level, keyvals ...interface{}) {
	switch f {
	case JSON:
		f.Json(buf, level, keyvals...)
	case STRING:
		f.String(buf, level, keyvals...)
	}
}

func (f Format) Json(buf *bytes.Buffer, level Level, keyvals ...interface{}) {
	fmt.Fprintf(buf, "{ \"level\"=\"%s\"", level.String())
	for i := 0; i < len(keyvals); i += 2 {
		if keyvals[i+1] == nil {
			continue
		}
		_, _ = fmt.Fprintf(buf, ", \"%s\":", keyvals[i])
		Json(buf, reflect.ValueOf(keyvals[i+1]))
	}
	buf.WriteString(" }")
}

func Json(buf *bytes.Buffer, value reflect.Value) {
	if value.Kind() == reflect.Interface {
		if value.IsNil() {
			buf.WriteString("null")
			return
		}
		value = value.Elem()
	}
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			buf.WriteString("null")
			return
		}
		value = value.Elem()
	}
	if value.Kind() == reflect.Struct {
		l := value.NumField()
		if l == 0 {
			buf.WriteString("{}")
			return
		}
		i := 0
		for ; i < l; i++ {
			field := value.Type().Field(i)
			val := value.FieldByName(field.Name)
			if !val.IsValid() || !val.CanInterface() {
				continue
			}
			name, ok := NameTag(field)
			if !ok {
				continue
			}
			fmt.Fprintf(buf, "{ \"%s\": ", name)
			Json(buf, value.FieldByName(field.Name))
			break
		}
		for i++; i < l; i++ {
			field := value.Type().Field(i)
			val := value.FieldByName(field.Name)
			if !val.IsValid() || !val.CanInterface() {
				continue
			}
			name, ok := NameTag(field)
			if !ok {
				continue
			}
			fmt.Fprintf(buf, ", \"%s\": ", name)
			Json(buf, val)
		}
		buf.WriteString(" }")
		return
	}
	fmt.Fprintf(buf, "\"%v\"", value.Interface())
}

func (f Format) String(buf *bytes.Buffer, level Level, keyvals ...interface{}) {
	fmt.Fprintf(buf, "level=%s", level.String())
	for i := 0; i < len(keyvals); i += 2 {
		_, _ = fmt.Fprintf(buf, " %s=%v", keyvals[i], keyvals[i+1])
	}
}

type Formater interface {
	HasValuer(kvs []interface{}) bool
	Bind(st interface{}) []interface{}
}

type formater struct {
	prefix    []interface{}
	hasValuer bool
}

func NewFormater(st interface{}) Formater {
	fmter := &formater{}
	fmter.prefix = fmter.Bind(st)
	return fmter
}

func (f *formater) HasValuer(kvs []interface{}) bool {
	if f.hasValuer {
		return true
	}
	if kvs == nil {
		return false
	}
	for _, v := range kvs {
		if _, ok := v.(Valuer); ok {
			return true
		}
	}
	return false
}

func (f *formater) PrefixCopy() []interface{} {
	if f.prefix == nil {
		return []interface{}{}
	}
	prefix := make([]interface{}, len(f.prefix))
	copy(prefix, f.prefix)
	return prefix
}

func (f *formater) Bind(st interface{}) []interface{} {
	switch st.(type) {
	case []interface{}:
		return f.BindKeyvals(st.([]interface{})...)
	case map[string]interface{}:
		return f.BindMap(st.(map[string]interface{}))
	default:
		return f.BindStruct(st)
	}
}

func (f *formater) BindStruct(st interface{}) []interface{} {
	prefix := f.PrefixCopy()
	if st == nil {
		return prefix
	}
	var value reflect.Value
	value = reflect.ValueOf(st)
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return prefix
		}
		value = value.Elem()
	}
	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		name, ok := NameTag(field)
		if !ok {
			continue
		}
		prefix = append(prefix, name, value.FieldByName(field.Name).Interface())
	}
	return prefix
}

func (f *formater) BindKeyvals(kvs ...interface{}) []interface{} {
	prefix := f.PrefixCopy()
	prefix = append(prefix, kvs...)
	return prefix
}

func (f *formater) BindMap(kvs map[string]interface{}) []interface{} {
	prefix := f.PrefixCopy()
	l := len(kvs)
	if l == 0 {
		return prefix
	}
	for k, v := range kvs {
		prefix = append(prefix, k, v)
	}
	return prefix
}
