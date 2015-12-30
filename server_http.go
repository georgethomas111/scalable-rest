package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
)

type Server struct {
	Port     int
	Protocol string
}

func NewServer(port int, proto string) (server *Server) {
	server = &Server{
		Port:     port,
		Protocol: proto,
	}
	return
}

func (ser *Server) HandleConnection(conn net.Conn) {
	// Close the connection what ever be the case.
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	req, err := http.ReadRequest(reader)
	if err != nil {
		// Let the world know
		fmt.Println("HandleConnection Http", err)
	} else {
		//	resp, err := http.ReadResponse(reader, req)
		if err != nil {
			fmt.Println("HandleConnection ReadResp", err)
		} else {
			json := GetJson(req)
			resp := &http.Response{
				Status:     "200 OK",
				StatusCode: 200,
				Proto:      req.Proto,
				Body:       ioutil.NopCloser(strings.NewReader(json)),
			}
			fmt.Println("Json for req:", json)
			err := resp.Write(writer)
			if err != nil {
				fmt.Println("Error writing", err)
			}
		}
	}
}

func (ser *Server) Listen() {
	strPort := ":" + strconv.Itoa(ser.Port)
	listener, err := net.Listen(ser.Protocol, strPort)
	if err != nil {
		// Nothing else can be done about it. Sorry.
		panic(err)
	} else {
		for {
			// Server is waiting. Now let the action begin.
			fmt.Println("Server waiting to Accept")
			conn, err := listener.Accept()
			if err != nil {
				// Lets log what the error is.
				fmt.Println(err)
			} else {
				fmt.Println("Processing connection", conn)
				ser.HandleConnection(conn)
			}
		}
	}
}
