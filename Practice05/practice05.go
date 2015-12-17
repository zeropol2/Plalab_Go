package main

import "fmt"
import "math/cmplx"

/*
* 뉴턴의 방법 참고
* https://en.wikipedia.org/wiki/Newton%27s_method
* http://ntalbs.github.io/2014/07/25/newtons-method/
* http://nosyu.pe.kr/1169
* f(x) = x^3 - X 에서 시작.
* xn+1 = xn - (xn^3 - X) / 3xn^2
* 초기값은 1로
* 실제 2의 3제곱근은 1.259921049894873
* 5의 3제곱근은 1.709975946676697
*
*/
func Cbrt(x complex128) complex128 {
    var result complex128 = 1
    for i:=0; i < 10; i++ {
        result = result - ((cmplx.Pow(result, 3) - x) / (3 * cmplx.Pow(result, 2)))
    }

    return result
}

func main() {
    fmt.Println(Cbrt(2))
}

