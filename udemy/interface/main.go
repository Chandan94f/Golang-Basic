package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	readWriteInterace()

	assignment1()

	assignment2()

}

// ---------- case 1 --------------

func readWriteInterace() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// fmt.Println(resp)

	// bs := []byte{}
	// n, err := resp.Read(b)

	// // to print all the body code - html
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	// different way

	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {

	// fmt.Println(string(bs))

	fmt.Println("len of bs: ", len(bs))

	return len(bs), nil

}

// ---------- case 2 --------------

func assignment1() {
	// exmaple - shape

	s := square{
		sideLength: 2.4,
	}

	t := triangle{
		base:   1.2,
		height: 1.6,
	}

	printArea(s)
	printArea(t)
}

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func (s square) getArea() float64 {
	fmt.Println("Area of Square is :")
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	fmt.Println("Area of Trianle is :")
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())

}

// ---------- case 3 --------------

func assignment2() {

	if len(os.Args) < 2 {
		fmt.Println("Missing text file")
		os.Exit(2)
	}

	fmt.Println("two files")
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(3)
	}

	io.Copy(os.Stdout, f)

}
