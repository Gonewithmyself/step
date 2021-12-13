package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	sd := newSdesc(&orderVo{})
	sd.genProto()
}

type sdesc struct {
	name string
	fds  []*reflect.StructField
}

func (sd *sdesc) genProto() {
	var buf strings.Builder
	buf.WriteString(fmt.Sprintf("message %v { \n", sd.name))
	for i, fd := range sd.fds {
		buf.WriteString(fmt.Sprintf(" %v %v = %v;\n", fd.Type, firstLower(fd.Name), i+1))
	}
	buf.WriteString("}\n")
	fmt.Println(buf.String())
}

func firstLower(s string) string {
	return strings.ToLower(s[0:1]) + s[1:]
}

func newSdesc(v interface{}) *sdesc {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
	}

	if vv.Kind() != reflect.Struct {
		panic(vv.Kind())
	}

	sd := &sdesc{}
	ty := vv.Type()
	sd.name = ty.Name()
	n := ty.NumField()
	sd.fds = make([]*reflect.StructField, n)
	for i := 0; i < n; i++ {
		fd := ty.Field(i)
		sd.fds[i] = &fd
	}
	return sd
}
