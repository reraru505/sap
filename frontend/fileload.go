package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type tokenizer struct{
	rawfile string // raw file 
	file []rune //rune version of the file loaded
	spacedfile []rune // file with space seperating special symbols
	tokens []string // list of tokens 
	LUT_special []rune // special symbol for look up table
}

//@step 1
func (l *tokenizer) init(){
	
	l.LUT_special = []rune{
			',',
			'<',
			'>',
			'.',
			'_',
			'(',
			')',
			';',
			'$',
			':',
			'%',
			'[',
			']',
			'#',
			'?',
			'\'',
			'&',
			'{',
			'}',
			'"',
			'^',
			'!',
			'*',
			'/',
			'|',
			'-',
			'\\',
			'~',
		'+'}
	
}

//@step 2
func (l *tokenizer) fileload(src string) {
	var err error
	raw, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println("Error reading file:", err.Error())
		os.Exit(1);
	}
	l.rawfile = string(raw)

	l.file = []rune (l.rawfile)

	
	
}

//@step 3
func (l *tokenizer) spacer(){
	for _ , i := range l.file {
		if l.is_special(i){
			l.spacedfile = append(l.spacedfile,' ')
			l.spacedfile = append(l.spacedfile,i)
			l.spacedfile = append(l.spacedfile,' ')
			
		}else{
			l.spacedfile = append(l.spacedfile,i)
		}

	}

}

//@step 4
func (l *tokenizer) tokenizer(){

	spacedstring := string(l.spacedfile)

	l.tokens = strings.Fields(spacedstring)

}



//@helper
func (l *tokenizer) is_special(c rune) bool {
	for _ , i := range l.LUT_special{
		if c == i {
			return true 
		}
	}

	return false
}


//@gives token as []stirng 

func get_tokens(src string) []string {
	var token_process tokenizer
	token_process.init()
	token_process.fileload(src)
	token_process.spacer()
	token_process.tokenizer()
	return token_process.tokens
	
}
