package main

import (
	"fmt"
	"strings"
	"os"
)

const (
	FKEY =iota //keyword 
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
	keys :=[12]string{
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
		"proc",
		"return"}

     for _ , v:= range keys {
	 if val == v {
	    return true
	 }
     }
     return false
     
}

//identifier

func check_identifier(val string) bool{
	arr := [] rune (val)
	keys :=   [54]rune {
		'a',
		'b',
		'c',
		'd',
		'e',
		'f',
		'g',
		'h',
		'i',
		'j',
		'k',
		'l',
		'm',
		'n',
		'o',
		'p',
		'q',
		'r',
		's',
		't',
		'u',
		'v',
		'w',
		'x',
		'y',
		'z',
		'A',
		'B',
		'C',
		'D',
		'E',
		'F',
		'G',
		'H',
		'I',
		'J',
		'K',
		'H',
		'L',
		'M',
		'N',
		'O',
		'P',
		'Q',
		'R',
		'S',
		'T',
		'U',
		'V',
		'W',
		'X',
		'Y',
		'Z',
		'_'}

	
	for _ , i := range keys{
		if(arr[0] == i){
			return true
		}
	}

	return false	

	
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
     keys:=[17]string{"!","%","*","-","+","/","|","&","&&","||",">","<",".",":",",","="}
   
     for _ ,v := range keys {
     	 if val==v {
	    return true
	  }
     }
     return false
}


//binds all functions 
func type_checker(val string) int{

	if check_Keyword(val){
		return FKEY
	}else if  check_literal(val){
		return FLIT
	}else if  check_operator(val){
		return FOP
	}else if  check_identifier(val){
		return FID
	}else if  val == "("{
		return FST
	}else if  val == ")"{
		return FEN
	}else if val == "{" {
		return FSTS
	}else if  val == "}"{
		return FENS
	}else{
		fmt.Println("There is something wronng with your code")
		os.Exit(1)
	}
	

	return -1

}
