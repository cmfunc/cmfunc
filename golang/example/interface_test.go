package example

import (
	"fmt"
	"testing"
)

func TestStudentSpeck(t *testing.T) {
	var peo People = &Student{}
	think := "love"
	fmt.Println(peo.Speak(think))
}

func TestPrintStarInterface(t *testing.T) {
	var a interface{} = nil
	var i *interface{} = &a
	PrintStarInterface(i)
}
