package main

import (
	//"fmt"
)

func interpreter(code []string){

	var vm machine

	vm.code = code
	
	vm.proccollector()
	vm.labelcollector()

	for _ , i := range vm.procarr{
		if "main" == i.name{
			vm.codeindex = i.index
			break
		}
	}

	for{
		
		if islabel(vm.code[vm.codeindex]){
			vm.codeindex++
			
		}else if vm.code[vm.codeindex] == "sys_exit"{
			break;
			
		}else if vm.code[vm.codeindex] == "push"{
			vm.push(vm.code[vm.codeindex+1])
			vm.codeindex++
			
		}else if vm.code[vm.codeindex] == "pop"{
			vm.pop()
			vm.codeindex++
			
		}else if vm.code[vm.codeindex] == "call"{
			vm.call(vm.code[vm.codeindex+1])
			//vm.codeindex++
			
		}else if vm.code[vm.codeindex] == "ret"{
			vm.ret()
			//fmt.Println("<debug> returned at ",vm.code[vm.codeindex])
			//vm.codeindex++
		}else if vm.code[vm.codeindex] == "lda"{
			vm.lda(vm.code[vm.codeindex+1])
			vm.codeindex++
			
		}else if vm.code[vm.codeindex] == "ldx"{
			vm.ldx(vm.code[vm.codeindex+1])
			vm.codeindex++
			
		}else if vm.code[vm.codeindex] == "ldy"{
			vm.ldy(vm.code[vm.codeindex+1])
			vm.codeindex++
			
		}else if vm.code[vm.codeindex] == "inc"{
			vm.inc(vm.code[vm.codeindex+1])
			vm.codeindex++
			
		}else if vm.code[vm.codeindex] == "add"{
			vm.add(vm.code[vm.codeindex+1],vm.code[vm.codeindex+2])
			vm.codeindex++
			
		}else if vm.code[vm.codeindex] == "sub"{
			vm.sub(vm.code[vm.codeindex+1],vm.code[vm.codeindex+2])
			vm.codeindex++
			
		}else if vm.code[vm.codeindex] == "mul"{
			vm.mul(vm.code[vm.codeindex+1],vm.code[vm.codeindex+2])
			vm.codeindex++
			
		}else if vm.code[vm.codeindex] == "div"{
			vm.div(vm.code[vm.codeindex+1],vm.code[vm.codeindex+2])
			vm.codeindex++
			
		}else if vm.code[vm.codeindex] == "mod"{
			vm.mod(vm.code[vm.codeindex+1],vm.code[vm.codeindex+2])
			vm.codeindex++
			
		}else if vm.code[vm.codeindex] == "cmp"{
			vm.cmp(vm.code[vm.codeindex+1],vm.code[vm.codeindex+2])
			vm.codeindex++
			
		}else if vm.code[vm.codeindex] == "je"{
			vm.je(vm.code[vm.codeindex+1])
			
		}else if vm.code[vm.codeindex] == "jne"{
			vm.jne(vm.code[vm.codeindex+1])
			
		}else if vm.code[vm.codeindex] == "jig"{
			vm.jig(vm.code[vm.codeindex+1])
			
		}else if vm.code[vm.codeindex] == "jil"{
			vm.jil(vm.code[vm.codeindex+1])
			
		}else if vm.code[vm.codeindex] == "print"{
			vm.print(vm.code[vm.codeindex+1])
			vm.codeindex++
		}

		
	}
}
