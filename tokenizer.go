package main

import (
	//"fmt"
       "strings"
)

const (
	FKEY=iota+1 //keyword 
	FID //identifier
	FLIT //literal
	FOP //operator
	FST //start
	FEN //end
	FSTS //start scope
	FENS //end scope
)

//keyword n identifier

func check_Keyword(val string) bool{
	keys :=[11]string{
		"i32",
		"i64",
		"f32",
		"f64",
		"string",
		"char",
		"if",
		"else",
		"while",
		"struct",
		"proc"
	}
     var bVal=false
     for _ , v:= range keys {
     	 if val==keys[v] {
	    bVal=true
	 }
     }
     return bVal
     
}

//literal

func check_literal(val string) bool{
     keys :=[10]string{"0","1","2","3","4","5","6","7","8","9"}
     var bVal=false
     str_split:=strings.Split(val,"")
     len_substr:=len(str_split)
     len_key:=len(keys)
     count:=0
     for i:=0;i<len_substr;i++ {
     	 for j:=0;j<len_key;j++ {
	     if(str_split[i]==keys[j])	{
	     	count++			
	     }
	 }
     }
     if count==len_substr {
     	bVal=true
     }
     return bVal
     
}

//special symbol

func check_operator(val string) bool{
     keys:=[14]string{"!","%","*","-","+","/","|","&","&&","||",">","<","."}
     var bVal=false
     for _ ,v:=range keys {
     	 if val==keys[v] {
	    bVal=true
	  }
     }
     return bVal
}


//binds all functions 
func type_checker(val string) int{

	switch{
		case check_keyword(val):
		return FKEY
		case check_literal(val):
		return FLIT
		case ckeck_operator(val):
		return FOP
		case val == "(":
		return FST
		case val == ")":
		return FEN
		case val == "{":
		return FSTS
		case val == "}":
		return FENS
	}

}
