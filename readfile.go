package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type ReadFile interface {
	Read(filename string)
}

type Rfi *os.File

func main() {
	filename := "slices.go"

	fileinfo, err := os.Stat(filename)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%v", fileinfo)
	}

	fi, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}

	Rfi = fi
	Rfi.ReadFile()

	fmt.Printf("ioutil start time %v \n", time.Now())
	fileByte, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("ioutil end time%v\n", time.Now())
	fmt.Println(string(fileByte))
}

func (r *Rfi) ReadFile() {
	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
}
