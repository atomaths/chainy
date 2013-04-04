package main

import (
	"chainy"
)

func through(s string) string {
	return "* " + s
}

func main() {
	chainy.MakeChain(through, through, through, through)
}
