package main

import (
	"encoding/json"
	"fmt"
	_ "io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	_ "strings"
)

type AppList struct {
	Applications []App
}

type App struct {
	Name         string
	StartCommand map[string]string `json:"start_command"`
	ProcName     string            `json:"procname"`
}

type WalkError struct {
	Message string
}

func (w WalkError) Error() string {
	return w.Message
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

func walkThing(command string) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fmt.Println("Checking", path)
		if is_pid(path) {
			fmt.Println("Visited:", path)
			if is_process_running(command, path) {
				return WalkError{"foo"}
			}
		}
		return nil
	}
}

func is_pid(directory string) bool {
	r, err := regexp.Compile(`[0-9]+$`)

	if err != nil {
		fmt.Printf("Problem compiling regexp")
	}

	if r.MatchString(directory) == true {
		return true
	}

	return false
}

func is_process_running(process, pid string) bool {
	cmdfile := fmt.Sprintf("%s/cmdline", pid)
	fmt.Println("Checking", cmdfile)
	f, ferr := ioutil.ReadFile(cmdfile)

	if ferr != nil {
		fmt.Println("opening of cmdline file failed")
		return false
	}

	return string(f) == process
}

func main() {
	procfs := "/tmp/proc"

	// read the json file into a string
	f, ferr := ioutil.ReadFile("conf/villein.json")
	if ferr != nil {
		//CJA: figure out logging
		fmt.Printf("cannot read config")
	}

	ap := NewAppList(f)
	fmt.Println(ap)

	// get command line arg

	err := filepath.Walk(procfs, walkThing("/home/melite/melite"))
	fmt.Printf("filepath.Walk() returned %v\n", err)
}
