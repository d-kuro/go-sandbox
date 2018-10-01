package main

import (
	"errors"
	"fmt"
)

// 使用しない定数
//const Url = "http://example.com"

type Person struct {
	name string
	// 使用されない
	//age int
}

func main() {
	//unconvert := "unconvert"
	// 意味のないキャスト
	//unconvert = string(unconvert)
	
	//person := &Person{name:"hoge"}
	//fmt.Println(person.name)
	
	// エラーチェックしない
	if err := returnError(); err != nil {
		fmt.Println("error")
	}
}

func returnError() error {
	fmt.Println("return error")
	return errors.New("new error")
}

// 使われない関数
//func unused() {
//	fmt.Println("unused")
//}
