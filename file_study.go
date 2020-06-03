package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main2() {
	fmt.Printf("--------")
	var b bytes.Buffer
	b.Write([]byte("Hellow"))

	fmt.Fprintf(&b, "World!")

	b.WriteTo(os.Stdout)

	fmt.Println(os.Args)

	r, err := http.Get("http://www.baidu.com")
	if err == nil {
		file, err2 := os.Create("./test.txt")
		if err2 == nil {
			defer file.Close()
			dest := io.MultiWriter(file, os.Stdout)
			io.Copy(dest,r.Body)
			//fmt.Println(r.Body)
			fmt.Println(dest)
		}
	}
}
