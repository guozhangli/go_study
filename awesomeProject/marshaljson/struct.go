package main

import (
	"bytes"
	"reflect"
)

func writeStruct(buff *bytes.Buffer,value reflect.Value)error{
	valueType:=value.Type()
	buff.WriteString("{")
	defer buff.WriteString("}")
	for i:=0;i<value.NumField();i++{
		fieldValue:=value.Field(i)
		fieldType:=valueType.Field(i)
		buff.WriteString("\"")
		buff.WriteString(fieldType.Name)
		buff.WriteString("\":")
		writeAny(buff,fieldValue)
		if i<value.NumField()-1{
			buff.WriteString(",")
		}
	}
	return nil
}