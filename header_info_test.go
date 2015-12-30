package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

// Test the header with mock requests.
// Expect the request client http as response
func TestHeader(t *testing.T) {

	req, err := http.NewRequest("GET", "http://dontcare.com", nil)
	if err != nil {
		panic(err)
	} else {
		req.Header.Set("Referer", "I love you")
		fmt.Println(GetJson(req))
		head := &Header{}
		err := json.Unmarshal([]byte(GetJson(req)), head)

		if err == nil {
			fmt.Println("Marshalled Object", head.Referer)
		} else {
			panic(err)
		}
	}

}
