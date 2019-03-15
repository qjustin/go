package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	f, err := os.Open("c://list.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, error := rd.ReadString('\n')
		if error != nil {
			break
		}

		if strings.Contains(line, "/") && !strings.Contains(line, "pica-cloud-configuration") && len(line) > 0 {

			fmt.Println(line)
		}
	}
}
