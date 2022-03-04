package controller

import "testing"

func testSomeLibrart(custom int) *Messages {
	return &Messages{
		SomeLibrary: func() (httpstatus int) {
			return custom
		},
	}
}

func TestA(t *testing.T) {
	m := WriteHello()
	m.HelloWorld()

	m2 := testSomeLibrart(5000)
	m2.HelloWorld()
}
