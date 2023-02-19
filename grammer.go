package main

import (
	"fmt"
)

type token struct{
	token_name string
	token_type int 
}

func tokenarray_get(arr []string) []token{
	var buffer []token
	
	for  _ , i := range arr{
		var tok token
		fmt.Println(i)
		tok.token_name = i
		tok.token_type = type_checker(i)
		buffer = append(buffer,tok)
	}

	return buffer
	
}



