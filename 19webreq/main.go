package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://github.com/TanmayChaurasia24"

func main() {
	fmt.Println("handling web request")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("response is:", response)
	fmt.Printf("type of response is: %T\n", response) // type of response is: *http.Response
	defer response.Body.Close()
	fmt.Println(response.Body)
	res, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("response after reading the req:", string(res))
}
