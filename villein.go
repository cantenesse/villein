package main

import (
  "path/filepath"
  "os"
  "flag"
  "fmt"
  _ "regexp"
  "io/ioutil"
  "encoding/json"
  _ "strings"
  _ "io"
)

type AppList struct {
  Applications []App
}

type App struct {
  Name string
  StartCommand map[string]string `json:"start_command"`
  ProcName string `json:"procname"`
}

func NewAppList(source []byte) *AppList {
  al := new(AppList)
  al.Applications = make([]App, 0, 0)
  al.FromJson(source)
  return al
}

func (al *AppList) FromJson(source []byte) {
  c := make(map[string]App)

  e := json.Unmarshal(source, &c)

  if e != nil {
    fmt.Printf("unable decode json")
  }

   for key, value := range c {
    value.Name = key
    al.Applications = append(al.Applications, value)
  }

}

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

  ap := NewAppList(f)
  fmt.Println(ap)
  //config := string(f)

 
  // get command line arg
  flag.Parse()
  
  root := flag.Arg(0)
  err := filepath.Walk(root, visit)
  fmt.Printf("filepath.Walk() returned %v\n", err)
}