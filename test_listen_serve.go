package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			errors.New(fmt.Sprintln("errorrrrr"))
		}
		go func() {
			fmt.Println("%v\n", conn.RemoteAddr())
			defer conn.Close()
		}()
		res := http.Response{
			StatusCode: 200,
			Body: ioutil.NopCloser(
				strings.NewReader("hello\n"),
			),
		}
		res.Write(conn)
	}
}
