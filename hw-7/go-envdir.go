package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	path := os.Args[1]
	// executable := os.Args[2]
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	os.Clearenv()
	for _, file := range files {
		fpath := filepath.Join(path, file.Name())
		rfile, err := os.OpenFile(fpath, os.O_RDWR, 0644)
		if err != nil {
			log.Fatal(err)
		}
		buf, err := ioutil.ReadAll(rfile)
		if err != nil {
			log.Fatal(err)
		}
		os.Setenv(file.Name(), string(buf))
		rfile.Close()
	}
	goExec, err := exec.LookPath("go")
	if err != nil {
		log.Fatal(err)
	}
	commandbuf := make([]string, 1)
	commandbuf = append(commandbuf, goExec)
	//commandbuf = append(commandbuf, executable)
	for _, file := range files {
		commandbuf = append(commandbuf, file.Name())
	}

	fmt.Println(commandbuf)
	// cmdGoProg := &exec.Cmd{
	// 	Path:   goExec,
	// 	Args:   commandbuf,
	// 	Stdout: os.Stdout,
	// 	Stderr: os.Stdout,
	// }
	// if err = cmdGoProg.Run(); err != nil {
	// 	fmt.Println("Error:", err)
	// }
}
