package main

import (
  //"fmt"
  "os"
  //"os/user"
  "simplexSimia/repl"
)

func main() {
  // user, err := user.Current( )
  // if err != nil {
  //   panic(err)
  // }
  // fmt.Printf("hello %s! This is the sim programming language!\n", user.Username)
  repl.Start(os.Stdin, os.Stdout)
}
