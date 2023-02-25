package main

import(
	"fmt"
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
	retval := strings.Fields(fil)
	return retval
}

type proc struct {
	name string
	index int
}

type label struct {
	name string
	index int
}

type machine struct{

	//registers
	A interface{}
	X interface{}
	Y interface{}

	//stack
	stack []interface{}
	top int
	bottom int

	//code
	code []string
	codeindex int

	//interpreter stuff
	procarr []proc
	labelarr []label
	mainindex int

	
	
	//flags
	carry    bool
	cmpf      bool
	overflow bool
	negative bool
	
}

//@do this before all <step one>

func (m * machine) proccollector() {

	for j , i := range m.code{
		mb := []rune(i)
		if(mb[0] == '.'){
			var buffer proc
			buffer.name = i
			buffer.index = j

			m.procarr = append(m.procarr,buffer)
		}
	}
}


func (m * machine) labelcollector() {

	for j , i := range m.code{
		mb := []rune(i)
		if(mb[0] == '_'){
			var buffer label
			buffer.name = i
			buffer.index = j

			m.labelarr = append(m.labelarr,buffer)
		}
	}
}



func (m *machine) push(val interface{}){
	m.stack = append(m.stack,val)
	m.top++
	m.codeindex++
}

func (m *machine) pop() interface{} {
	if m.top == 0 {
		return 0
	}
	a := m.stack[m.top]
	m.top--
	return a
}

func (m *machine) call(val string){

	//@todo handle arguments
	m.bottom =  m.top
	m.push(m.A)
	m.push(m.X)
	m.push(m.Y)
	m.push(m.codeindex)

	m.mainindex = m.codeindex+1
	
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

func (m *machine) add(a interface{},b interface{}){
	c := a.(int)
	d := b.(int)

	m.A = c + d
	m.codeindex += 2
}

func (m *machine) sub(a interface{},b interface{}){
	c := a.(int)
	d := b.(int)

	m.A = c - d

	m.codeindex += 2
}

func (m *machine) mul(a interface{} ,b interface{}){
	c := a.(int)
	d := b.(int)

	m.A = c * d

	m.codeindex += 2
}

func (m *machine) div(a interface{} ,b interface{}){
	c := a.(int)
	d := b.(int)

	m.A = c / d

	m.codeindex += 2
}

func (m *machine) mod(a interface{} ,b interface{}){
	c := a.(int)
	d := b.(int)

	m.A = c % d

	m.codeindex += 2
}

func (m *machine) cmp(a interface{},b interface{}){
	if a == b {
		m.cmpf = true
	}else{
		m.cmpf = false
	}

}

func (m *machine) je(a interface{}){
	val := a.(string)

	if m.cmpf == true{
		for _ , i := range m.labelarr{
			if val == i.name{
				m.codeindex = i.index
				break
			}
		}
	}
	
}



//@ this is a debug function and must be replaced
func (m *machine) print(val interface{}){
	fmt.Print(val)
	m.codeindex += 1
}
//@ remove this !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

//@helpers or rectifier

func islabel(val string) bool{        //yes this also includes procs
	i := []rune(val)

	if i[0] == '_' || i[0] == '.'{
		return true
	}

	return false
}
