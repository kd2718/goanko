package main

import (
  "fmt"
  "github.com/fernandosanchezjr/anko/vm"
  "github.com/kd2718/goanko/goscripts"
)

var env = vm.NewEnv()

var blaster = func(name string) {
  fmt.Println(name, "is great")
}

func main() {
  env.Define("println", fmt.Println)
  out, err := goscripts.Anko_inner_def_call(env)

  if err != nil {
    fmt.Println("error", err)
  } else {
    fmt.Println("ok", out)
  }

}
