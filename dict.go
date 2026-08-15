package main

import (
	"math/rand"
)

type Word struct {
	Latin   string
	Cyril   string
	Russian string
}

var Dict = []Word{
	{Latin: "salam", Cyril: "салам", Russian: "привет"},
	{Latin: "kitap", Cyril: "китап", Russian: "книга"},
	{Latin: "bash", Cyril: "баш", Russian: "голова"},}
}
func RandomWord() Word {
	return Dict[rand.Intn(len(Dict))]
}
