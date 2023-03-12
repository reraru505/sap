package main

import (
	"fmt"
)

type token struct{
	token_name string
	token_type int
	token_operator_name string
} 

type typesystem struct{

	
}

func tok_name(val int) string{
     switch val{
     	    case 0:
			return "KEYWORD"
	    case 1:
			return "LITERAL"
	    case 2:
			return "START"
	    case 3:
			return "END"
	    case 4:
			return "START SCOPE"
	    case 5:
			return "END SCOPE"
	    case 6:
			return "ASSIGNMENT OPERATOR"
	    case 7:
			return "EXCLAMATION"
	    case 8:
			return "MODULUS"
	    case 9:
			return "STAR"
	    case 10:
			return "SUBTRACTION"
	    case 11:
			return "ADDITION"
	    case 12:
			return "DIVISION"
	    case 13:
			return "SINGLELINE"
	    case 14:
			return "DOUBLE LINE"
	    case 15:
			return "AND"
	    case 16:
			return "AND AND"
	    case 17:
			return "GREATER THAN"
	    case 18:
			return "LESS THAN"
	    case 19:
			return "STOP"
	    case 20:
			return "COLON"
	    case 21:
			return "COMMA"
	    default:
			return "IDENTIFIER"
			
     }
}
//concatinate strings
func flatten_string(val []string)[]string{
     var buffer []string
     var buffer2 string
     var jval int
     for i:=0;i<len(val);i++{
     	 if val[i]==string(0x22){
	    buffer=append(buffer,val[i])
	    count:=0
	    for j:=i+1;j<len(val);j++{
	    	 if val[j]==string(0x22) && count!=1{
		      count=1
		      jval=j
		    }
		 if count==1{
		    continue
		 }
		}
	    
	    fmt.Println("Val of jval: ",jval)
	    for k:=i+1;k<jval;k++{
	    	buffer2=buffer2+val[k]
	    }
	    buffer=append(buffer,buffer2)
		buffer2=""

	    i=jval
	 }
	 buffer=append(buffer,val[i])
	 
	 }
	 return buffer
     
}
func tokenarray_get(arr []string) []token{
	var buffer []token
	arr2:=flatten_string(arr)
	for  _ , i := range arr2{
		var tok token
		fmt.Println(i)
		tok.token_name = i
		tok.token_type = type_checker(i)
		tok.token_operator_name=tok_name(tok.token_type)
		buffer = append(buffer,tok)
	}
	return buffer
	
}



