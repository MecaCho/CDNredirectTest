package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!\n")

	io.WriteString(w, "http://162.3.210.33:5000")

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
	if r.Method == "GET" {
		t, err := template.ParseFiles("hello.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
		return
	}
	if r.Method == "POST" {
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
		ipRed := strings.Split(string(body), "/")[2]
		fmt.Println("#######", ipRed)

		http.Redirect(w, r, ipRed, 302)
	}
	//http.Redirect(w, r, ipRed, 302)
}

func main() {
	//	result := postHandler()
	//postHandler()
	//	fmt.Println("result", result)
	http.HandleFunc("/p2", testpostHandler)
	//http.Redirect(w,r,)
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
