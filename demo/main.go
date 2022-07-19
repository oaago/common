package main

import (
	"github.com/oaago/common/validate"
)

func main() {
	//a, _ := jwt.GenerateToken("1", map[string]interface{}{
	//	"qqq": "1",
	//})
	//b, r := jwt.ParseToken(a)
	//fmt.Println(b, r)
	//go queue.Test()
	//for i := 0; i < 100; i++ {
	//	queue.Create()
	//}
	inter := []validate.Interfaces{
		{
			Url:       "/aaa/bbb",
			Method:    "post1",
			Ttl:       10,
			Desc:      "1",
			ChildNode: []validate.Interfaces{},
		},
	}
	validate.Input(inter)
	select {}
}
