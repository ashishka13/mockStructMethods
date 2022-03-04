package controller

import (
	"fmt"
	"log"
)

type Messages struct {
	SomeLibrary func() (httpstatus int)
}

func (m Messages) HelloWorld() {
	log.Println("-=-=-=-=-=-=-==-")
	fmt.Println("-=-=-=-=-=-=-==-")

	a := m.SomeLibrary()
	fmt.Println(a)
}

func WriteHello() *Messages {
	return &Messages{
		SomeLibrary,
	}
}

func SomeLibrary() (httpstatus int) {
	httpstatus = 200
	return
}
