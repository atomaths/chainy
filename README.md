# Chainy :link: :link: :link: :neckbeard:

## What?

    os.Stdin ->  func(string) string) -> func(string) string -> .... -> func(string) string -> os.Stdout

## How?
   
    $ go get github.com/minikomi/chainy

## Then?

   Make this guy:

### chainytest.go

```go
package main

import (
    "github.com/minikomi/chainy"
)

func through(s string) string {
    return "* " + s
}

func through2(s string) string {
    return "- " + s
}

func main() {
    chainy.MakeChain(through, through2, through2, through)
}
```

## And? 

        minikomi•~» cat chainytest.go | go run chainytest.go | sort | uniq -c | go run chainytest.go
        * - - *       4 * - - * 
        * - - *       1 * - - * )
        * - - *       3 * - - * }
        * - - *       1 * - - *   "chainy"
        * - - *       1 * - - *         chainy.MakeChain(through, through2, through2, through)
        * - - *       1 * - - * func main() {
        * - - *       1 * - - * func through2(s string) string {
        * - - *       1 * - - * func through(s string) string {
        * - - *       1 * - - * import (
        * - - *       1 * - - * package main
        * - - *       1 * - - *         return "- " + s
        * - - *       1 * - - *         return "* " + s

:bowtie:
