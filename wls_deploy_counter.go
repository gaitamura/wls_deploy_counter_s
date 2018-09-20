package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func worker(host []string) <-chan bool{
	wk := make(chan bool)

	for _, s := range host {
		go func(s string) {
			var fp *os.File
			var err error

			if len(os.Args) < 2 {
				fp = os.Stdin
			} else {
				fp, err = os.Open(os.Args[1])
				if err != nil {
					panic(err)
				}
					defer fp.Close()
			}

			counter := 0
			reg := regexp.MustCompile(`正常に完了.*配布.*` + s)
			scanner := bufio.NewScanner(fp)
			for scanner.Scan() {
				if reg.MatchString(scanner.Text()) {
					counter += 1
				}
			}
			if err := scanner.Err(); err != nil {
				panic(err)
			}
			fmt.Println(s, counter)
			wk <- true
			}(s)
	}
	return wk
}

func main() {
	host := []string{"managed_server1", "managed_server2", "managed_server3"}
	wk := worker(host)

	for i := 0; i < len(host); i++ {
		if <-wk {
			fmt.Printf("")
		}
	}
}
