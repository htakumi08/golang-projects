//1,mainサーバとなる場所の設定
//2,各関数をコード

package main

import (
	"fmt"
	"log"
	"net/http"
)

// r*http.Request : 構造体のポインタ→メモリのコピーをしない、効率的
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	// Fprintf()インターフェースに実装する任意のオブジェクトに書き込む
	// wは構造体
	fmt.Fprintf(w, "POST request succesful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s", name)
	fmt.Fprintf(w, "address = %s", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not suported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("/usr/src/app/goprograming/go-server/static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at port 8888\n")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
