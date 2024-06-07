package benching_test

import (
	"task19/funcs"
	"testing"
)


func BenchmarkReverse(b *testing.B) {
	var data [10]string = [10]string{"Python", "Ruby", "ASM", "Go", "Pascal", "C#", "Java", "Javascript", "Rust", "Nim"}
	
	for i := range b.N {
		funcs.ReverseString(data[i%10])
	}
}

func BenchmarkReverseV2(b *testing.B) {
	var data [10]string = [10]string{"Python", "Ruby", "ASM", "Go", "Pascal", "C#", "Java", "Javascript", "Rust", "Nim"}
	
	for i := range b.N {
		funcs.ReverseStringV2(&data[i%10])
	}
}


