package main

import (
	"fmt"
	"reflect"
	"time"
)

type Stu struct {
	Str  string
	Time time.Time
}

func TestTime() {
	stu := Stu{Str: "test", Time: time.Now()}
	print(stu)
}

func print(t interface{}) {
	getType := reflect.TypeOf(t)
	getValue := reflect.ValueOf(t)
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		switch field.Type.String() {
		case "time.Time":
			fmt.Printf("%s: %v = %v\n", field.Name, field.Type, getValue.Field(1).Interface().(time.Time))
			break
		default:
			value := getValue.Field(i)
			fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value.String())
			break
		}
	}
}

func main() {
	TestTime()
}
