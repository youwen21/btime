package main

import (
	"btime"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "github.com/imroc/biu"
	// "reflect"
	// "testing"
)

type resultS struct {
	Ret [5]uint64 `json:"ret"`
	Err error     `json:"err"`
}

func main() {
	server1()
	// server2()
}

func server1() {
	log.Println("start")
	http.HandleFunc("/", btimeService)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
	log.Println("Listening...")
}

func btimeService(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var startTime string = ""
	var endTime string = ""
	if len(r.Form["start"]) > 0 {
		startTime = r.Form["start"][0]
	} else {
		w.Write([]byte("lost start param"))
	}
	if len(r.Form["end"]) > 0 {
		endTime = r.Form["end"][0]
		// w.Write([]byte("The time is: " + endTime))
	}

	data, err := btime.GetBinary(startTime, endTime)
	if err != nil {
		w.Write([]byte("有错误"))
	}
	result := resultS{data, err}
	// fmt.Println("result is : ")
	// fmt.Println(result)
	json, err := json.Marshal(result)
	// w.Write([]byte("<Br/>"))
	if err != nil {
		w.Write([]byte("转json发生错误 "))
	}
	// fmt.Println("json is : ")
	// fmt.Println(json)
	w.Write(json)
}
