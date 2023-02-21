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

type proc struct {
	name string
	index int
}

type machine struct{

	A interface{}
	x interface{}
	Y interface{}
	stack []interface{}
	top int
	bottom int
	code []string
	codeindex int
	procarr []proc
	mainindex int

}

//@do this before all <step one>

func (m * machines) proccollector() {

	for j , i := range m.code{
		m := []rune(i)
		if(m[i] == '.'){
			var buffer proc
			buffer.name = i
			buffer.index = j

			m.procarr = append(m.procarr,buffer)
		}
	}
}


func (m *machine) push(val rune){
	m.stack = append(m.stack,val)
	m.top++
}

func (m *machine) pop() interface{
	if top == 0 {
		return 0
	}
	a := m.stack[top]
	top--
	return a
}

func (m *machine) call(val string){

	//@todo handle arguments
	m.bottom =  m.top
	m.push(m.A)
	m.push(m.X)
	m.push(m.y)
	m.push(codeindex)

	m.mainindex = m.codeindex
	for _ , i := range m.procarr{
		if val == i.name{
			m.codeindex = i.index
			break
		}
	}
	
}

func (m *machine) ret(){
	
	m.Y = m.pop()
	m.X = m.pop()
	m.A = m.pop()

	m.codeindex = m.mainindex
		
}


