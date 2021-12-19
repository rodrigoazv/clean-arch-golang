package main

import "fmt"

type ContaCorrente struct {
	user    string
	agency  int
	accout  int
	balance float64
}

func main() {
	cc := ContaCorrente{user: "Rodrigo", agency: 123, accout: 123, balance: 123.4}

	fmt.Println(cc)
}
