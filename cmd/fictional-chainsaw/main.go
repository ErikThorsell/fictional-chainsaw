package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang/gddo/httputil/header"
)

type OpList struct {
	Operation string
	Numbers   []int
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("This is a calculator!\n"))
}

func help(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Send a POST request to URL/<operation> with a list of integers to try it out.\n"))
	w.Write([]byte("Example data format: {\"operation\":\"sum\",\"numbers\":[1,2,3]}\n"))
	w.Write([]byte("Suported operations are \"sum\" and \"mul\"."))
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func calc(w http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
		w.Write([]byte("\nCalc only handles POST requests."))

	} else {

		if req.Header.Get("Content-Type") != "" {
			value, _ := header.ParseValueAndParams(req.Header, "Content-Type")
			if value != "application/json" {
				msg := "Content-Type header is not application/json"
				http.Error(w, msg, http.StatusUnsupportedMediaType)
				return
			}
		}

		var postBody OpList

		err := json.NewDecoder(req.Body).Decode(&postBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		switch postBody.Operation {

		case "sum":
			s := 0
			for _, i := range postBody.Numbers {
				s += i
			}
			result := fmt.Sprintf("The sum of %v is %d", postBody.Numbers, s)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(result))

		case "mul":
			s := 1
			for _, i := range postBody.Numbers {
				s *= i
			}
			result := fmt.Sprintf("The product of %v is %d", postBody.Numbers, s)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(result))

		default:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
			w.Write([]byte("\nCalc only handles sum and mul requests."))
		}
	}
}

func main() {
	fmt.Println("Starting web calculator!")
	http.HandleFunc("/", index)
	http.HandleFunc("/help", help)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/calc", calc)
	http.ListenAndServe(":8090", nil)
}
