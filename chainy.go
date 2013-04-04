package chainy

import (
	"bufio"
	"io"
	"os"
)

func makeLink(inCh <-chan string, outCh chan<- string, f func(string) string) {
	for {
		s := <-inCh
		outCh <- f(s)
	}
}

func MakeChain(funcs ...func(string) string) {
	var inCh, outCh chan string

	stdInCh := make(chan string)
	inCh = stdInCh

	for _, f := range funcs {
		outCh = make(chan string)
		go makeLink(inCh, outCh, f)
		inCh = outCh
	}

	r := bufio.NewReader(os.Stdin)
	go (func() {
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				os.Exit(0)
			}
			stdInCh <- line
		}
	})()

	for {
		io.WriteString(os.Stdout, <-outCh)
	}
}
