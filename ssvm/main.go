package main

import "os"

func main(){

	i := fileload(os.Args[1])

	interpreter(i)

}
