package main

import(
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"strconv"
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
	A string
	X string
	Y string

	//stack
	stack [1024*1024]string
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
	cmpf     bool
	greater  bool
	lesser   bool
	overflow bool
	negative bool
	
}

//@do this before all <step one>

func (m * machine) proccollector() {

	for j , i := range m.code{
		mb := []rune(i)
		if(mb[0] == '.'){
			var name []rune
			for _ , k := range mb {
				if k != '.'{
					name = append(name,k)
				}
			}
			var buffer proc
			buffer.name = string(name)
			buffer.index = j

			m.procarr = append(m.procarr,buffer)
		}
	}
}


func (m * machine) labelcollector() {

	for j , i := range m.code{
		mb := []rune(i)
		if(mb[0] == '#'){
			var name []rune
			for _ , k := range mb {
				if k != '#'{
					name = append(name,k)
				}
			}
			var buffer label
			buffer.name = string(name)
			buffer.index = j

			m.labelarr = append(m.labelarr,buffer)
		}
	}
}

func (m *machine) lda(val string){
	if val == "[X]"{
		m.A = m.X
	}else if val == "[Y]"{
		m.A = m.Y
	}else{
		m.A = val
	}

	m.codeindex++
}


func (m *machine) ldx(val string){
	if val == "[A]"{
		m.X = m.A
	}else if val == "[Y]"{
		m.X = m.Y
	}else{
		m.X = val
	}
	m.codeindex++
}


func (m *machine) ldy(val string){
	if val == "[X]"{
		m.Y = m.X
	}else if val == "[A]"{
		m.Y = m.A
	}else{
		m.Y = val
	}
	m.codeindex++
}

func (m *machine) inc (val string){

	if val == "[X]"{
		i,_ := strconv.Atoi(m.X)
		i++
		m.X = strconv.Itoa(i)
	}else if val == "[A]"{
		i,_ := strconv.Atoi(m.A)
		i++
		m.A = strconv.Itoa(i)
	}else if val == "[Y]"{
		i,_ := strconv.Atoi(m.Y)
		//fmt.Println("<debug> value of Y = ",i)
		i++
		m.Y = strconv.Itoa(i)
	}else{
		fmt.Println("<debug> feture incomplete")
	}
	m.codeindex++
}



func (m *machine) push(val string){
	m.stack[m.top] = val
	m.top++
	m.codeindex++
}

func (m *machine) pop() string{
	if m.top == 0 {
		return "0"
	}
	m.top--
	return m.stack[m.top]
}

func (m *machine) call(val string){

	//@todo handle arguments
	m.bottom =  m.top
	m.push(m.A)
	m.push(m.X)
	m.push(m.Y)

	m.mainindex = m.codeindex-1
	
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
	m.bottom = m.top
	m.codeindex = m.mainindex 
		
}

func (m *machine) add(a string,b string){
	var av string
	var bv string
	if a == "[X]"{
		av = m.X
	}else if a == "[Y]"{
		av = m.Y
	}else if a == "[A]"{
		av = m.A
	}else {
		av = a
	}

	if b == "[X]"{
		bv = m.X
	}else if b == "[Y]"{
		bv = m.Y
	}else if b == "[A]"{
		bv = m.A
	}else {
		bv = b
	}

	c,_ := strconv.Atoi(av)
	

	d,_ := strconv.Atoi(bv)

	var e int

	// fmt.Println ("c ",c ,"and ","d ",d)
	
	e = c + d 

	m.A = strconv.Itoa(e)
	m.codeindex += 2
}

func (m *machine) sub(a string,b string){
	c,_ := strconv.Atoi(a)
	d,_ := strconv.Atoi(b)

	e := c - d
	m.A = strconv.Itoa(e)
	
	m.codeindex += 2
}

func (m *machine) mul(a string ,b string){
	c,_ := strconv.Atoi(a)
	d,_ := strconv.Atoi(b)

	e := c * d
	m.A = strconv.Itoa(e)

	m.codeindex += 2
}

func (m *machine) div(a string,b string){
	c,_ := strconv.Atoi(a)
	d,_ := strconv.Atoi(b)

	e := c / d
	m.A = strconv.Itoa(e)

	m.codeindex += 2
}

func (m *machine) mod(a string ,b string){
	c,_ := strconv.Atoi(a)
	d,_ := strconv.Atoi(b)

	e := c % d
	m.A = strconv.Itoa(e)

	m.codeindex += 2
}

func (m *machine) cmp(a string,b string){
	var av string
	var bv string
	
	if a == "[X]"{
		av = m.X
	}else if a == "[Y]"{
		av = m.Y
	}else if a == "[A]"{
		av = m.A
	}else {
		av = a
	}

	if b == "[X]"{
		bv = m.X
	}else if b == "[Y]"{
		bv = m.Y
	}else if b == "[A]"{
		bv = m.A
	}else{
		bv = b
	}

	i,_ := strconv.Atoi(av)
	j,_ := strconv.Atoi(bv)
	
	if j == i {
		m.cmpf = true
		m.greater = false
		m.lesser = false
	}else if  i > j {
		
		m.cmpf = false
		m.greater = true
		m.lesser = false
		
	}else if i < j {
		
		m.cmpf = false
		m.greater = false
		m.lesser = true
	}

	m.codeindex += 2

}


// JUMP IF EQUAL

func (m *machine) je(val string){

	if m.cmpf == true{
		for _ , i := range m.labelarr{
			if val == i.name{
				m.codeindex = i.index
				break
			}
		}
		m.codeindex++
	}else{
		m.codeindex += 2
	}

	
	
}

//JUMP IF NOT EQUAL 

func (m *machine) jne(val string){

	if m.cmpf == false{
		for _ , i := range m.labelarr{
			if val == i.name{
				m.codeindex = i.index
				break
			}
		}
		m.codeindex++
	}else{
		m.codeindex += 2
	}
	
	
	
}

// JUMP IF GREATER
func (m *machine) jig(val string){

	if m.greater == true{
		for _ , i := range m.labelarr{
			if val == i.name{
				m.codeindex = i.index
				break
			}
		}
		m.codeindex++
	}else{
		m.codeindex += 2
	}
	
	
	
}


//JUMP IF LESSER
func (m *machine) jil(val string){

	if m.lesser == true{
		for _ , i := range m.labelarr{
			if val == i.name{
				m.codeindex = i.index
				break
			}
		}
		m.codeindex++
	}else{
		m.codeindex += 2
	}
	
	
	
}





//@ this is a debug function and must be replaced
func (m *machine) print(val string){
	if val == "[A]" {
		fmt.Print(m.A)
		
	}else if val == "[X]" {
		fmt.Print(m.X)
		
	}else if val == "[Y]" {
		fmt.Print(m.Y)
		
	}else if val == "<newline>" {
		fmt.Println()
		
	}else{
		fmt.Print(val)
	}
	
	m.codeindex += 1
}
//@ remove this !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

//@helpers or rectifier

func islabel(val string) bool{        //yes this also includes procs
	i := []rune(val)

	if i[0] == '#' || i[0] == '.'{
		return true
	}

	return false
}
