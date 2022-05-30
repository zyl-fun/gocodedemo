package main

import (
	"testing"
)

type bar struct {
	foo int
}

func useforI() {
	var s = []bar{{0}, {1}, {2}}
	for i := range s {
		s[i].foo += 10000
	}
	//fmt.Println(s)
}

func BenchmarkUseForI(b *testing.B) {
	for n := 0; n < b.N; n++ {
		useforI()
	}
}

func modifyPtr() {
	s := []*bar{&bar{0}, &bar{1}, &bar{2}}
	for _, item := range s {
		item.foo += 10000
	}
}

func BenchmarkUsePtr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		modifyPtr()
	}
}
