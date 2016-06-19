Flag2
---
A more traditional flag library for the go programming language

What?
---

A more triditional flag library for the Go programming language. I also have a long history with Python, so the implimentation code looks similar to Python's argparse class.

Why?
---

I did not like how the flag library that comes with Go parses command line flags.

Differences
---

* You can define full word flags with the `--` prefix. You can define single character flags with the `-` prefix.

* Example of a full word flag: `--help`

* Example of a single character flag: `-h`

* Single character strings can be grouped, but only for boolean types: `-abcd` is essentially `-a, -b, -c, -d`
  * This only works for boolean type flags

Getting started
---

To install: `go get github.com/ProfOak/flag2`

```
package main

import (
    "os"
    "fmt"
    "github.com/ProfOak/flag2"
)

func main() {
    f := flag2.NewFlag()

    // short flag, long flag, description, default argument
    f.AddString("n", "name", "this flag wants a name as input", "billy")
    f.AddBool("b", "bool", "this flag will store true", false)

    // a help flag is added during the parse step
    options, args := f.Parse(os.Args)

    // A usage method is provided, with details about each flag

    // unfortunate side effect of interfaces
    if options["help"] == true {
        f.Usage()
    }

    fmt.Println()
    if options["name"] != nil {
        fmt.Println("The name is:", options["name"])
    }

    fmt.Println()
    fmt.Println("===== FINAL RESULTS =====")
    fmt.Println("Options:", options)
    fmt.Println("Args:", args)
}

```

The result of running this program:

```
go run main.go -b -n ProfOak Extra args

--- Bools ---
-b, --bool      this flag will store true
-h, --help      Display this message and exit

--- Strings ---
-n, --name      this flag wants a name as input

Name is: ProfOak

===== FINAL RESULTS =====
Options: map[help:true bool:true name:ProfOak]
Args: [Extra args]

```

Reference ./test/test.go for a more detailed example.
