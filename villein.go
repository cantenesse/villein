package main

import (
  "path/filepath"
  "os"
  "flag"
  "fmt"
  //"regexp"
  "io/ioutil"
  "encoding/json"
//  "strings"
//  "io"
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

  // read the json file into a string
  f, ferr := ioutil.ReadFile("conf/villein.json")
  if ferr != nil {
    //CJA: figure out logging
    fmt.Printf("cannot read config")
  }

  //config := string(f)

  c := make(map[string]interface{})

  e := json.Unmarshal(f, &c)

  if e != nil {
    fmt.Printf("unable decode json")
  }

  k := make([]string, len(c))

  i := 0

  for s, _ := range c {
    k[i] = s
    i++
  }

  fmt.Printf("%#v\n", k)
 
  // get command line arg
  flag.Parse()
  
  root := flag.Arg(0)
  err := filepath.Walk(root, visit)
  fmt.Printf("filepath.Walk() returned %v\n", err)
}