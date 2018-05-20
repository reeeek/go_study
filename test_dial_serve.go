package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	res, _ := http.ReadResponse(bufio.NewReader(conn), nil)
	fmt.Printf("status : %v\n", res.Status)
	fmt.Printf("header : %v\n", res.Header)
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
