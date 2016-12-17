package main

import (
  "fmt"
  "github.com/fernandosanchezjr/anko/vm"
)

var env = vm.NewEnv()


func main() {
  env.Define("blast", func(name string) {
    println(name, "You just got BLASTED!!!")
  })
  fmt.Println("hello")
  who := "Kory Donati"
  exe_str := fmt.Sprint("blast(\"", who, "\")")
  env.Execute(exe_str)

}
