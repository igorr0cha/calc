package main

import (
	"net/http"
	"time"
	"strings"
	"strconv"
	"fmt"
	"encoding/json"
)
var ValidPaths = []string{"/result"}

func isPathValid(path string) bool{
	for i := 0 ; i < len(ValidPaths) ; i++{
		if(path==ValidPaths[i]){
			return true
		}
	}
	return false
}

type Serve struct {}

func (Serve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if(r.Method != "GET"){
		w.Write([]byte("{\"code\": 405, \"message\": MethodNotAllowed}"))
		return
	}

	if (!isPathValid(r.URL.Path)){
		w.WriteHeader(404)
		w.Write([]byte("{\"code\": 404, \"message\": invalid path}"))
		return
	}
	
	// trying:

	opQuery:= r.URL.Query().Get("op")
	piece:= strings.Fields(opQuery)

	if len(piece) != 3 {
		w.WriteHeader(406)
		w.Write([]byte("{\"code\": 406, \"message\": Expression NotAcceptable}"))
	}

	n1, err1 := strconv.ParseFloat(piece[0],64)
	n2, err2 := strconv.ParseFloat(piece[2],64)

	if err1!= nil || err2 != nil {
		w.WriteHeader(400)
		http.Error(w,err1.Error(),http.StatusBadRequest)
		return
	}

	op := Operation{
		Num1: n1,
		Num2: n2,
		TypeOp: piece[1],
	}

	Calc:= op.Calculate()
	result, err := Calc(n1,n2)

	if err!= nil {
		http.Error(w,err.Error(), http.StatusBadRequest)
	}

	res := ResJson{
		Result: result,
		Operation: piece,
	}

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w,"Result: %2f", result)

	encode:= json.NewEncoder(w)
	encode.Encode(res)

}

func NewServer(s Serve) *http.Server{
	return &http.Server{
			Addr: ":8080",
			Handler: s,
			ReadTimeout: 1 * time.Second,
			WriteTimeout: 1 * time.Second,
		}
}


