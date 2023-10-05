package main

import (
	// 导入就会执行init方法
	"Ax/lib"
	//. "Ax/lib" 不推荐
	// "Ax/lib"
	mx "Ax/lib2"
	"fmt"

	//导入但是不使用（报错）  调用init 方法
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
	// 发现上面的引用依赖报错 执行 go mod tidy
)

func main() {
	mx.GetSMS()
	lib.GetSums()
	fmt.Println("main")
}

func init() {
	fmt.Println("main init")
}
