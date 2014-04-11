package main

import (
  "path/filepath"
  "os"
  "flag"
  "fmt"
  "regexp"
)

func visit(path string, f os.FileInfo, err error) error {
  fmt.Printf("Visited: %s\n", path)
  return nil
} 

func is_pid(directory string) {
  r, err := regexp.Compile(`[0-9]+`)
  
  if err != nil {
    fmt.Printf("Problem compiling regexp")
  }

  if r.MatchString(directory) == true {
    return true
  } else {
    return false
  }
}

func main() {
  flag.Parse()
  
  root := flag.Arg(0)
  err := filepath.Walk(root, visit)
  fmt.Printf("filepath.Walk() returned %v\n", err)
}