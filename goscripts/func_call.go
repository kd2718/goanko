package goscripts


import (
	"fmt"
  "github.com/fernandosanchezjr/anko/vm"
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

func Anko_inner_def_call(env *vm.Env) (string, error){
	script := `
	var x = 0
	println(x,y)
	var out = func(x int){
	println("you picked")
	}
	`
	//	var killer = func(x, y int){
	//println(x, y)
	//}
	//killer(a, b)
	env.Define("x", 5)
	env.Define("y", 20)

	return basic_exe(env, script)
}