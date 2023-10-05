package lib2

import "fmt"

func init() {
	fmt.Println("init lib2")
}

func GetSMS() int {
	fmt.Println("lib2 sms")
	return 0
}
