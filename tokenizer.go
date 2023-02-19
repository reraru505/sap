package main

import (
       "fmt"
       "strings"
)

const (
      keyword=iota+1
      identifier
      literal
      special_symbol
      start
      end
      startscope
      endscope
)
//keyword n identifier

func check_Keyword(val string) bool{
     keys :=[13]string{"break","case","continue","default","else","fallthrough","for","goto","if","range","return","select","switch"}
     var bVal=false
     for v:= range keys {
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

func check_special_symbol(val string) bool{
     keys:=[11]string{"!","#","$","*","-","+",","}
     var bVal=false
     for v:=range keys {
     	 if val==keys[v] {
	    bVal=true
	  }
     }
     return bVal
}

//start

func check_start(val string) bool{
     if(val=="(") {
     		  return true
     }else{
	return false
     }
}
//end
func check_end(val string) bool{
     if(val==")") {
     		  return true
     }else{
	return false
     }
}

//startscope

func check_startscope(val string) bool{
     if(val=="{") {
     		  return true
     }else{
	return false
     }
}

//endscope

func check_endscope(val string) bool{
     if(val=="}") {
     		  return true
     }else{
	return false
     }
}
