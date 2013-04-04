package chainy

import (
	"bufio"
	"io"
	"os"
)

func MakeChain(funcs ...func(string) string) {
	r := bufio.NewReader(os.Stdin)
	final := make(chan error)
	outCh := make(chan string)
	go (func() {
		for {
			line, err := r.ReadString('\n')
			if err != nil {
				final <- err
				return
			}
			for _, f := range funcs {
				line = f(line)
			}
			outCh <- line
		}
	})()

	for {
		select {
		case s := <-outCh:
			io.WriteString(os.Stdout, s)
		case err := <-final:
			if err == io.EOF {
				return
			}
			panic(err)
		}
	}
}
