package main

import(
	"strings"
	"io/ioutil"
	"os"
)

func fileload(src string) []string{

	raw,err := ioutil.ReadFile(src)
	if err != nil{
		os.Exit(1)
	}

	fil := string(raw)
	retval := string.Fields(fil)
	return retval
}

