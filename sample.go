package main

import (
	"fmt"
	"net/http"
	"os"
	"math/rand"
	"strconv"
)

const Port = 8080

func main() {
	http.HandleFunc("/", HandleTopPage)
	err := http.ListenAndServe(":" + strconv.Itoa(Port), nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func HandleTopPage(w http.ResponseWriter, r *http.Request) {
	results := []string{"大吉", "吉", "小吉", "凶", "大凶"}
	fmt.Fprintf(w, "<html><body><p>%s</p></body></html>\n",
		results[rand.Intn(len(results))])
}
