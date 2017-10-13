package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	var s, e, l int    
	var d string
	flag.IntVar(&s, "s", 1, "start page, default 1")
	flag.IntVar(&e, "e", 1, "end page, deafault 1")
	flag.IntVar(&l, "l", 72, "how many lines in single page, deafault 72")
	flag.StringVar(&d, "d", "", "pipe or not, default nil")
	flag.Parse()
	fileName := flag.Args()[0]
	f, err := os.Open(fileName)
	if err != nil {
		f.Close()
		os.Exit(1)
	}

	begin := (s - 1) * l
	end := e * l
	count := 1
	scanner := bufio.NewScanner(f)
	cmdinfo := ""
	for scanner.Scan() {
		line := scanner.Text()
		if count > begin && count <= end {
			if d == "" {                    // without d param, then print to stdout
				fmt.Println(line)
			} else {
				cmdinfo = cmdinfo + line + "\n"
			}
		}
	}
	if d != "" {                       // with d param,  pipe the info read from file to a command
		cmdArg := "-d" + d
		cmd := exec.Command("lp", cmdArg)
		cmdin, cmderr := cmd.StdinPipe()
		if cmderr != nil {
			cmdin.Close()
			os.Exit(1)
		}
		io.WriteString(cmdin, cmdinfo)
	}
}
