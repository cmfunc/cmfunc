package example

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "love" {
		talk = "I love you too"
	} else {
		talk = "hello"
	}
	return
}

func PrintStarInterface(i *interface{})  {
	fmt.Printf("%#v", i)
}
