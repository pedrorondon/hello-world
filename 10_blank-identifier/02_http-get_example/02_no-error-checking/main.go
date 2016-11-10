package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func man() {
	res, _ := http.Get("http://www.mcloed.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
