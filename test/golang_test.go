package test

import (
	"fmt"
	"testing"
	"github.com/oyamo/telegram-compiler/src"
)

const (
	test_code = `package main
	import "fmt"
	func main() {
		fmt.Println("Hello, 世界")
	}`
)

func TestGolang(t *testing.T) {
	res, err := src.Golang(test_code)
	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Error("Empty response")
	}

	if res.StatusCode != 200 {
		t.Error("Invalid status code " +fmt.Sprint(res.StatusCode))
	}

	t.Log(res.Body)
}
