package biz

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	parser := &cronParser{}
	dur := parser.Parser("0 21 * * *")
	fmt.Println(dur)
}
