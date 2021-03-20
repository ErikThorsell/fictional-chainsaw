package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(index)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, wanted %v", status, http.StatusOK)
	}

	expected := "This is a calculator!\n"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned wrong status code: got %v, wanted %v", rr.Body.String(), expected)
	}

}

func TestSum(t *testing.T) {

	post_req := []byte(`{"operation": "sum", "numbers": [1,2,3,4,5]}`)
	req, err := http.NewRequest("POST", "/calc", bytes.NewBuffer(post_req))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(calc)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, wanted %v", status, http.StatusOK)
	}

	expected := "The sum of [1 2 3 4 5] is 15"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned wrong status code: got %v, wanted %v", rr.Body.String(), expected)
	}

}

func TestMul(t *testing.T) {

	post_req := []byte(`{"operation": "mul", "numbers": [1,2,3,4,5]}`)
	req, err := http.NewRequest("POST", "/calc", bytes.NewBuffer(post_req))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(calc)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, wanted %v", status, http.StatusOK)
	}

	expected := "The product of [1 2 3 4 5] is 120"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned wrong status code: got %v, wanted %v", rr.Body.String(), expected)
	}

}
