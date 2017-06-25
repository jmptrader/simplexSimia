package repl

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "simplexSimia/lexer"
  "simplexSimia/token"
)

func Start(in io.Reader, out io.Writer) {

  hostname,err := os.Hostname()
    if err == nil {
    }

  PROMPT := ("@" + hostname + ":\\>>")

  scanner := bufio.NewScanner(in)

  for {
    fmt.Printf(PROMPT)
    scanned := scanner.Scan()
    if !scanned {
      return
    }

    line := scanner.Text()
    l := lexer.New(line)

    for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
      fmt.Printf("%+v\n", tok)
    }
  }
}
