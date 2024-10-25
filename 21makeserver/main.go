package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("making req in golang")
	PerformGetRequest()

}

func PerformGetRequest() {
	const url = "http://localhost:3000/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	defer resp.Body.Close()

	fmt.Println("status code:", resp.StatusCode)
	fmt.Println("content length:", resp.ContentLength)

	var responsestring strings.Builder
	content, _ := ioutil.ReadAll(resp.Body)

	// another way
	bytecode, _ := responsestring.Write(content)
	fmt.Println(bytecode)
	fmt.Println(responsestring.String()) // home page

	fmt.Println(string(content)) // home page
}

// &{200 OK 200 HTTP/1.1 1 1 map[Connection:[keep-alive] Content-Length:[9] Content-Type:[text/html; charset=utf-8] Date:[Fri, 25 Oct 2024 10:09:06 GMT] Etag:[W/"9-MGNMHSlmVAdeK1+AtfLnHw6r5jA"] Keep-Alive:[timeout=5] X-Powered-By:[Express]] 0xc0001a0040 9 [] false false map[] 0xc00007a280 <nil>}
// status code: 200
// content length: 9
// home page
