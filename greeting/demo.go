package main

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type demo struct {
	Lat, Long float64
}

var m map[string]demo

func compute(x, y float64, fn_name func(float64, float64) float64) int {
	return int(fn_name(x, y))
}

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return MyError{
		time.Now(),
		"it didn't work",
	}
}

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(b []byte) (int, error) {
	a := 'A'
	for i := 0; i < len(b); i++ {
		b[i] = byte(a)
	}
	return len(b), nil
}
func main() {

	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	hypot := func(x, y float64) float64 {
		return x + y
	}
	fmt.Println(hypot(5, 12))
	fmt.Printf("Kishor")
	fmt.Println(compute(6, 3, hypot))

	// Create a tic-tac-toe board.
	board := make([][]string, 3)
	board[0] = make([]string, 3, 3)
	board[1] = make([]string, 3, 3)
	board[2] = make([]string, 3, 3)
	// The players take turns.
	board[0][0] = "X"
	board[0][1] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	fmt.Println("Is:" + board[2][0] + ": ads")
	for i := 0; i < len(board); i++ {
		//fmt.Println(board[i])
		//fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
