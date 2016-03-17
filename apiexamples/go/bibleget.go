package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func main() {

	resp, err := http.Get("http://bibletag.xyz:8080/tag/Gospel")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	
	fmt.Println(string(body))

}