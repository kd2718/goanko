package goscripts


import (
	"fmt"
  "github.com/fernandosanchezjr/anko/vm"
	"github.com/fernandosanchezjr/anko/parser"
)
var blast = func(){
	fmt.Println("I will blast you!!")
}

var blast2 = func(temp string){
	fmt.Println("this is a blast!", temp)
}

func basic_exe(env *vm.Env, script string) (string, error){
	success, err := env.Execute(script)

	return fmt.Sprint(success), err
}

func Anko_print_script(env *vm.Env) (string, error) {
	script := `
	println("hello world... I am Kory")
	`

	return basic_exe(env, script)
}


func Anko_call_func(env *vm.Env) (string, error){


	env.Define("blast", blast)
	script := `
	blast()
	println("finished!!!")
	`

	return basic_exe(env, script)
}

func Anko_def_call(env *vm.Env) (string, error){
	script := `
	var blast = func(){
	println("you got totally blasted!@#JLLK")
	}
	blast()
	`
	return basic_exe(env, script)
}

func Anko_print_gvar(env *vm.Env) (string, error){
	script := `
	println("this is my test", test)
	`
	env.Define("test", "kory")

	return basic_exe(env, script)
}

func Anko_pass_value(env *vm.Env) (string, error){
	script := `
	blast2(test)
	println("all done")
	`
	name := "Kory Donati"
	env.Define("test", name)
	env.Define("blast2", blast2)

	return basic_exe(env, script)
}

func Anko_inner_def_call(env *vm.Env) (string, error) {
	script := `
	println(x,y)
	var out = func(x, y){
	println("you picked", x+y)
	}
	out(x, y)
	`
	pscript, err := parser.ParseSrc(script)
	if err != nil {
		fmt.Println("error parsing", err)
		return "", err
	}
	env.Define("x", 5)
	env.Define("y", 20)
	out, err:= vm.Run(pscript, env)

	return fmt.Sprint(out), err
}