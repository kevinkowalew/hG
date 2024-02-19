package main

import (
	"bufio"
	"fmt"
	"hg/list"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		args = append(args, scanner.Text())
	}

	if scanner.Err() != nil {
		panic("Failed to read input from STDIN: " + scanner.Err().Error())
	}

	l := list.NewSingleSelectList(args, 20)
	line, err := l.Run()
	if err != nil {
		fmt.Println("SingleSelectList exited with error: " + err.Error())
		return
	}

	p := strings.Split(line, " ")
	args = []string{}
	if len(p) > 1 {
		args = p[1:]
	}

	o, err := exec.Command(
		p[0],
		args...,
	).Output()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Print(string(o))
	}
}
