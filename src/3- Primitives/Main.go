package main

import (
	"fmt"
)

func main() {
	var n bool = false
	fmt.Printf("%v %T", n, n)
	fmt.Println()

	n = 1 == 1
	m := 1 == 2
	fmt.Printf("%v %T", n, n)
	fmt.Println()
	fmt.Printf("%v %T", m, m)
	fmt.Println()

	//Booleans sempre são iniciadas como false
	var j bool
	fmt.Printf("%v %T", j, j)
	fmt.Println()

	//------------------------------//
	//int8 ------------- (-128 - 127)
	//int16 ------------- (-32768 - 32767)
	//int32 ------------- (-2147483648 - 2147483647)
	//int64 ------------- (-9223372036854775808 - 9223372036854775807)
	//Maior q isso, olhar na Math

	//Todos os tipos int, tem seus equivalentes em uint, com excessao do de 64bits
	//Porém, temos o tipo Byte

	a := 10
	b := 3
	fmt.Println(a + b) //13
	fmt.Println(a - b) //7
	fmt.Println(a * b) //30
	fmt.Println(a / b) //3
	fmt.Println(a % b) //1

	//a/b nos da resto 3, por conta do tipo que estamos usando
	// a e b sao int, então, o resultado nos dará um int
	// isso significa que desprezaremos a parte fracional, sem aproximar

	// não conseguimos fazer operaçoes com tipos diferentes também

	var c int = 10
	var d int8 = 3
	//precisamos fazer um cast para que isso funcione
	fmt.Println(c + int(d))

	//a = 10 -> 1010
	//b = 3  -> 0011

	//vai olhar quais bits estão setados (1) no primeiro e segundo numero
	fmt.Println(a & b)  //2  -> 0010 (operação &, ou seja, os dois com 1)
	fmt.Println(a | b)  //11 -> 1011 (operação |, ou seja, pelo menos um com 1)
	fmt.Println(a ^ b)  //9  -> 1001 (operação ^, pelo menos um com 1, mas não os dois)
	fmt.Println(a &^ b) //8 -> 0100 (operação &^, ou seja, nenhum pode ter o bit setado)

	//bit shifting
	t := 8
	fmt.Println(t << 3) //64 -> 8 * 2^3
	fmt.Println(t >> 3) //1  -> 8 / 2^3

	//float
	var p float32 = 3.14
	p = 2.1e14
	fmt.Println(p)

	i := 10.3
	o := 3.9
	fmt.Println(i + o)
	fmt.Println(i - o)
	fmt.Println(i * o)
	fmt.Println(i / o)

	//complexos
	//temos 2 tipos, complex64 e complex128
	var l complex64 = 1 + 2i
	var k complex64 = 3 + 9i
	fmt.Printf("%v %T", l, l)
	fmt.Println()
	fmt.Printf("%v %T", k, k)
	fmt.Println()
	fmt.Println(l + k)
	fmt.Println(l - k)
	fmt.Println(l * k)
	fmt.Println(l / k)

	//pega a parte real
	fmt.Printf("%v %T", real(k), real(k))
	fmt.Println()

	//pega a parte imaginaria
	fmt.Printf("%v %T", imag(k), imag(k))
	fmt.Println()

	//gerando um numero imaginario
	var g complex128 = complex(5, 12)
	fmt.Printf("%v %T", g, g)
	fmt.Println()

	//string
	str := "this is a string :)"
	fmt.Printf("%v %T", str, str)
	fmt.Println()

	//podemos considerar como um array de caracteres tbm
	fmt.Printf("%v %T", string(str[2]), string(str[2]))
	fmt.Println()

	str1 := "this is a string "
	str2 := "and this is another string :)"
	fmt.Printf("%v %T", str1+str2, str1+str2)
	fmt.Println()

	str3 := "this is a string"
	arrayBytes := []byte(str3)
	fmt.Printf("%v %T", arrayBytes, arrayBytes) // ascii values
	fmt.Println()

	//runes
	//runes são int32
	r := 'a'
	fmt.Printf("%v %T", r, r)

}
