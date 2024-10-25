package main

import (
	"fmt"
	"net/url"
)

const URL string = "http://127.0.0.1:5500/20handleURL/index.html?course=reactnative&golang"

func main() {
	fmt.Println("handling url in golang")
	fmt.Println("URL:", URL)

	//parsing
	parsedURL, err := url.Parse(URL)
	if err != nil {
		panic(err)
	}
	fmt.Println("host is:", parsedURL.Host)         // host is: 127.0.0.1:5500
	fmt.Println("scheme is:", parsedURL.Scheme)     // scheme is: http
	fmt.Println("path is:", parsedURL.Path)         // path is: /20handleURL/index.html
	fmt.Println("rawquery is:", parsedURL.RawQuery) // they same like params in javascript // rawquery is: course=reactnative&golang
	fmt.Println("port is:", parsedURL.Port())       // port is: 5500

	qparams := parsedURL.Query()
	fmt.Println(qparams)                                   // map[course:[reactnative] golang:[]]
	fmt.Printf("types of query params are: %T\n", qparams) // types of query params are: url.Values

	// making url from parts
	partsURL := &url.URL{
		Scheme:   "https",
		Host:     "127.0.0.1:5500",
		Path:     "/20handleURL/index.html",
		RawQuery: "course=reactnative&golang",
	}

	anotherURL := partsURL.String()
	fmt.Println(anotherURL) // https://127.0.0.1:5500/20handleURL/index.html?course=reactnative&golang
}
