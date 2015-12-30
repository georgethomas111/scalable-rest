package main

import (
	"encoding/json"
	"net/http"
)

type Header struct {
	Referer string `json:"referer"`
}

func NewHeader(req *http.Request) (head *Header) {
	head = &Header{
		Referer: req.Referer(),
	}
	return
}

func GetJson(req *http.Request) string {

	head := NewHeader(req)
	//FIXME: ignore err for now.
	data, _ := json.Marshal(head)
	return string(data)
}
