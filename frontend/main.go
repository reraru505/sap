package main

import(
	"fmt"
)


func main(){

	var token_process tokenizer
	token_process.init()
	token_process.fileload("example.txt")
	token_process.spacer()
	token_process.tokenizer()

	for _ ,i := range token_process.tokens{
		fmt.Println(i)
	}

	fmt.Println("End")
	b := tokenarray_get(token_process.tokens)
	
	
	blen := len(b)
	for i := 0 ; i < blen ; i++{
		fmt.Println("< ",b[i].token_name ," : " ,b[i].token_type ,":",b[i].token_operator_name," >")
	}
}
