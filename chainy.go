package chainy

import (
	"bufio"
	"io"
	"os"
)

func MakeChain(funcs ...func(string) string) {
	r := bufio.NewReader(os.Stdin)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		for _, f := range funcs {
			line = f(line)
		}
		io.WriteString(os.Stdout, line)
	}
}
