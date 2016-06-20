package main

import (
	"fmt"
	"github.com/ProfOak/flag2"
	//"os"
)

func main() {

	f := flag2.NewFlag()

	// short flag, long flag, description, default value
	f.AddBool("s", "short", "Test short flag", false)
	f.AddBool("l", "long", "Test long flag", false)

	// you can have multi short flags for bools
	f.AddBool("x", "example", "Test for the multi short flag", false)
	f.AddBool("v", "never", "This flag will never happen in this test", false)

	f.AddString("n", "name", "Store a person's name", "John")
	f.AddString("a", "age", "Store a person's age", "42")

	// you cannot have multi short flags for strings
	f.AddString("z", "zero", "Failing test for -xyz", "bad")

	// test to enter two instances of the same flag
	if err := f.AddString("n", "name", "Store a person's name", "John"); err != nil {
		// if it already exists, then f does not change
		// use an error if you want it to be explicit
		fmt.Println()
		fmt.Println("===== EXAMPLE ERROR =====")
		fmt.Println(err)
	}

	// running Prase will add a help flag
	// it will display the list of flags and what they do
	//options, args := f.Parse(os.Args)

	test_args := []string{
		"-s",       // single short arg
		"-xyz",     // multiple short arg
		"--long",   // single long arg
		"-a", "12", // short string arg (age)
		"--name", "billy", // long string arg
		"--", "--never", // --never will go to args because -- denotes the end of options
		"foot", "loose", // loose arguments (not options)

	}
	options, args := f.Parse(test_args)

	fmt.Println()
	fmt.Println("===== COMMAND LINE ARGUMENTS =====")
	fmt.Println(test_args)

	fmt.Println()
	fmt.Println("===== ALL FLAG KEYS =====")
	fmt.Println(f.FlagKeys())

	fmt.Println()
	fmt.Println("===== USAGE =====")

	// unfortunate side-effect of interfaces
	if options["help"] == true {
		f.Usage()
	}

	fmt.Println()
	if options["name"] != nil {
		fmt.Println("Name is:", options["name"])
	}

	fmt.Println()
	fmt.Println("===== FINAL RESULTS =====")
	fmt.Println("Options:", options)
	fmt.Println("Args:", args)

}
