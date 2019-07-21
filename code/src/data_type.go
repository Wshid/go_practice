package main
// import "fmt"

func main(){
	// `(back quote), raw string literal, 있는 그대로 저장
	// 복수라인의 문자열 표현시 사용
	// escaping이 되지 않음
	rawLiteral := `firstline\n second
	third
	`
	interpretedLiteral := "firstline\nsecond\nthird"

	var a = 4
	b := 4
	println(a, b)

	println(rawLiteral)
	println(interpretedLiteral)
}
/**
  * =와 :=의 의미 : https://stackoverflow.com/questions/17891226/difference-between-and-operators-in-go
  * 아래 두 식은 같다.
  * var foo int = 10
  * foo := 10
  * */