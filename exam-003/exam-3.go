package main

import "fmt"

func main() {
	var mp map[string]string
	fmt.Println(mp)

	//var mp2 map[string]string = map[string]string()
	mp2 := map[string]string{}
	mp2["test"] = "ok"

	fmt.Println(mp2)

	fmt.Println("Hello-3")
}
