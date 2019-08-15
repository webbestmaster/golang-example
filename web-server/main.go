package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening on port :3000")

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		bytesLength, fprintfError := fmt.Fprintf(writer, "Hello World")

		if fprintfError != nil {
			fmt.Printf(fprintfError.Error())
		}

		fmt.Println("Sent bytes:", bytesLength)
	})

	listenAndServeError := http.ListenAndServe(":3000", nil)

	if listenAndServeError != nil {
		fmt.Printf(listenAndServeError.Error())
	}
}
