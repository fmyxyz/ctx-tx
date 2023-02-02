package main

import (
	"database/sql"
	"reflect"
	"strings"
)

func main() {
	//structMethod(reflect.TypeOf(&xorm.Engine{}), reflect.TypeOf(&xorm.Session{}))
	structMethod(reflect.TypeOf(&sql.DB{}), reflect.TypeOf(&sql.Tx{}))
}

// 打印struct相同签名的method
func structMethod(dbT, txT reflect.Type) {

	dbMethods := getMethodSigns(dbT)
	txMethods := getMethodSigns(txT)

	//
	//for _, dbMethod := range dbMethods {
	//	println(dbMethod)
	//}
	//for _, txMethod := range txMethods {
	//	find := false
	//	for _, dbMethod := range dbMethods {
	//		find = dbMethod == txMethod
	//		if find {
	//			break
	//		}
	//	}
	//	if !find {
	//		println(txMethod)
	//	}
	//}

	for _, txMethod := range txMethods {
		for _, dbMethod := range dbMethods {
			if dbMethod == txMethod {
				println(dbMethod)
			}
		}
	}

	println()

}

func getMethodSigns(dbT reflect.Type) []string {
	dbMethods := []string{}
	for i := 0; i < dbT.NumMethod(); i++ {
		method := dbT.Method(i)
		t := method.Func.Type()
		name := method.Name + "("
		for j := 1; j < t.NumIn(); j++ {
			in := t.In(j)
			s := in.String()
			if j+1 == t.NumIn() && t.IsVariadic() {
				s = "..." + strings.TrimPrefix(s, "[]")
			}
			name += s
			if j+1 != t.NumIn() {
				name += ","
			}
		}
		numOut := t.NumOut()
		if numOut > 0 {
			name += ")("
		} else {
			name += ")"
		}
		for j := 0; j < numOut; j++ {
			name += t.Out(j).String()
			if j+1 != numOut {
				name += ","
			}
		}
		if numOut > 0 {
			name += ")"
		}

		name = strings.ReplaceAll(name, "interface {}", "any")
		dbMethods = append(dbMethods, name)
	}
	return dbMethods
}
