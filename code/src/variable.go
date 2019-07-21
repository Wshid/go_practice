package main

func main(){
	var a int = 4
	var b int = 5
	const c string = "Hi"
	// var b float32 = 6.0
	println(a+b)
	println(c)

	// 타입 자동 추론
	const auto_typing = "Hello"

	// 상수 여러개 동시 선언
	const {
		multi = "abcde"
		ple = "cdede"
		va = "ab"
	}

	// 순차 상수 부여, 0,1,2, ...
	const {
		Apple = itoa
		Grape
		Orange
	}
}