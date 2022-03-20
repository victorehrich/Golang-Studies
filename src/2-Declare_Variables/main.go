package main

import (
	"fmt"
	"strconv"
)

//podemos declarar a variável aqui tbm, mas temos que
//escrever a declaração completa
var z int = 42

//podemos fazer blocos de variáveis tbm
var (
	nome  = "davi"
	idade = 15
)

//variáveis minusculas são variáveis visiveis no package que elas estão
//variáveis maiusculas são visíveis globalmente

func main() {
	//variáveis dentro do bloco so são visíveis no bloco
	var j int
	j = 27
	fmt.Println(j)

	var i int = 42
	fmt.Println(i)

	b := 42
	fmt.Println(b)

	//printa o valor da variável e o tipo dela
	n := 99
	fmt.Printf("%v, %T", n, n)
	fmt.Println()

	//com esse tipo de declaração, podemos ter int como no exemplo
	//acima, ou float64, como no abaixo
	m := 99.
	fmt.Printf("%v, %T", m, m)
	fmt.Println()

	//shadowing é quando temos uma variável no package level
	//e a mesma variável redeclarada na função. O valor dela será
	//o valor a partir do momento que ela é declarada na função
	fmt.Println(z)
	var z int = 9
	fmt.Println(z)

	//variáveis locais devem ser usada, se não, o código apontará um erro

	//Conversão de variáveis (cast)

	var q int = 42
	fmt.Printf("%v, %T", q, q)
	fmt.Println()
	var w float32
	w = float32(q)
	fmt.Printf("%v, %T", w, w)
	fmt.Println()

	//perda de informação
	var e float32 = 42.5
	fmt.Printf("%v, %T", e, e)
	fmt.Println()
	var r int
	r = int(e)
	fmt.Printf("%v, %T", r, r)
	fmt.Println()

	//converções para string são utilizando o packagr "strconv"
	var t int = 42
	fmt.Printf("%v, %T", t, t)
	fmt.Println()
	var y string
	y = strconv.Itoa(t)
	fmt.Printf("%v, %T", y, y)
	fmt.Println()
}
