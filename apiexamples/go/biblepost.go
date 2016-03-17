package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
)

func main() {

    url := "http://bibletag.xyz:8080/tag"

    var jsonStr = []byte(`{
		"tag": "worry",
		"book": "1 peter",
		"chapter": 5,
		"startVerse": 7,
		"endVerse": 7
	}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))

}