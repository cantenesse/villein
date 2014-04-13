package main

import (
  "path/filepath"
  "os"
  "flag"
  "fmt"
  //"regexp"
  "io/ioutil"
  "encoding/json"
  "strings"
  "io"
)

func visit(path string, f os.FileInfo, err error) error {
  fmt.Printf("Visited: %s\n", path)
  return nil
} 

//func is_pid(directory string) {
//  r, err := regexp.Compile(`[0-9]+`)
//  
//  if err != nil {
//    fmt.Printf("Problem compiling regexp")
//  }
//
// if r.MatchString(directory) == true {
//   return true
// } else {
//   return false
// }
//

func main() {
  type Config struct {
    Process, Script, Type string
  }

  // read the json file into a string
  f, ferr := ioutil.ReadFile("conf/villein.conf")
  if ferr != nil {
    //CJA: figure out logging
    fmt.Printf("cannot read config")
  }

  config := string(f)

  dec := json.NewDecoder(strings.NewReader(config))
  for {
    var c Config
    if err := dec.Decode(&c); err ==io.EOF {
      break
    } else if err != nil {
      fmt.Printf("can't decode json")
    }
    fmt.Printf("%s: %s\n", c.Process, c.Script)
  }

  // get command line arg
  flag.Parse()
  
  root := flag.Arg(0)
  err := filepath.Walk(root, visit)
  fmt.Printf("filepath.Walk() returned %v\n", err)
}