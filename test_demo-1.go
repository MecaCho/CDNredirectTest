package main

import (
	"bytes"
	"fmt"
	//	"html/template"
	//	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	//	io.WriteString(w, "Hello World!\n")
	http.Redirect(w, r, "http://114.115.145.32:8081/", 302)
	//	io.WriteString(w, "http://162.3.210.33:5000")

}

func testpostHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://162.3.210.33:5000/123"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"eventName":"cdn-redirect"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "cdn-redirect-test")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response Body:", string(body))
	filep, _ := os.Create("./1.txt")
	filep.WriteString(string(body))
	//	result := []rune(string(body))
	//	result1 := strings.Split(string(body), ":")[5]
	ipRed := strings.Split(string(body), "/")[2]
	//	fmt.Println("#######", result1)
	fmt.Println("#######", ipRed)
	//	return result1
	//	return "Success"

	http.Redirect(w, r, "http://114.115.145.32:8081/", 302)
}

func main() {

	http.HandleFunc("/p2", testpostHandler)
	http.HandleFunc("/p3", testHandler)
	//http.Redirect(w,r,)
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
