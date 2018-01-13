// flag2 allows the use of more traditional unix-style flags in Go programs.
package flag2

import (
	"fmt"
	//"os"
	//"regexp"
	"strings"
)

// New instantiates memory and returns something usable
// by the client application.
func New() FlagStruct {
	f := FlagStruct{}
	f.bools = make(map[string]boolFlag)
	f.strings = make(map[string]stringFlag)
	f.ints = make(map[string]intFlag)
	f.floats = make(map[string]floatFlag)
	return f
}

// validateFlag checks to see if it can possibly be
// added to the list of flags. If it can, return a
// flagProps struct with details of type non-specific info.
func (f *FlagStruct) validateFlag(metavar, short, long, desc string) (flagProps, error) {
	var p flagProps

	if metavar == "" {
		return p, fmt.Errorf("Must have metavar identifier")
	} else if listContains(f.metavars, metavar) {
		return p, fmt.Errorf("[%s]: Flag already exists", metavar)
	} else if short == "" && long == "" {
		return p, fmt.Errorf("[%s]: Must have at least one flag identifier for use in program", metavar)
	} else if len(short) > 1 {
		return p, fmt.Errorf("[%s]: Short must only be one character", metavar)
	} else if strings.Contains(long, " ") {
		return p, fmt.Errorf("[%s]: no spaces allowed in long flags", metavar)
	}

	// seems like a valid flag
	p = flagProps{
		Metavar: metavar,
		Short:   short,
		Long:    long,
		Desc:    desc,
	}
	return p, nil
}

// AddBool adds a boolean flag to your command line arguments.
// This will default to false and if selected then change to true.
//
// parameters:
//     metavar (string): program's way of accessing flag value
//     short (string): single character flag (empty quotes if none)
//     long (string): long name flag (empty quotes if none)
//     desc (string): help description
//     val (bool): default value (defaults to false)
func (f *FlagStruct) AddBool(metavar, short, long, desc string) error {

	props, err := f.validateFlag(metavar, short, long, desc)
	if err != nil {
		return err
	}

	// valid flag
	f.metavars = append(f.metavars, metavar)
	f.bools[metavar] = boolFlag{
		Props: props,
		Value: false,
	}
	return nil
}

// AddString adds a string flag to your command line arguments.
// if you don't want to use a short or a long, provide an empty string.
//
// parameters:
//     metavar (string): program's way of accessing flag value
//     short (string): single character flag (empty quotes if none)
//     long (string): long name flag (empty quotes if none)
//     desc (string): help description
//     val (string): default value (empty quotes if none)
func (f *FlagStruct) AddString(metavar, short, long, desc, val string) error {

	props, err := f.validateFlag(metavar, short, long, desc)
	if err != nil {
		return err
	}

	// valid flag
	f.metavars = append(f.metavars, metavar)
	f.strings[metavar] = stringFlag{
		Props: props,
		Value: val,
	}
	return nil
}

// AddInt adds a string flag to your command line arguments.
// if you don't want to use a short or a long, provide an empty string.
//
// parameters:
//     metavar (string): program's way of accessing flag value
//     short (string): single character flag (empty quotes if none)
//     long (string): long name flag (empty quotes if none)
//     desc (string): help description
//     val (int): default value
func (f *FlagStruct) AddInt(metavar, short, long, desc string, val int) error {

	props, err := f.validateFlag(metavar, short, long, desc)
	if err != nil {
		return err
	}

	// valid flag
	f.metavars = append(f.metavars, metavar)
	f.ints[metavar] = intFlag{
		Props: props,
		Value: val,
	}
	return nil
}

// AddFloat adds a string flag to your command line arguments.
// if you don't want to use a short or a long, provide an empty string.
//
// parameters:
//     metavar (string): program's way of accessing flag value
//     short (string): single character flag (empty quotes if none)
//     long (string): long name flag (empty quotes if none)
//     desc (string): help description
//     val (float): default value
func (f *FlagStruct) AddFloat(metavar, short, long, desc string, val float64) error {

	props, err := f.validateFlag(metavar, short, long, desc)
	if err != nil {
		return err
	}

	// valid flag
	f.metavars = append(f.metavars, metavar)
	f.floats[metavar] = floatFlag{
		Props: props,
		Value: val,
	}
	return nil
}

/*

// Parse will return a list of flags and arguments based on argv
// TODO: get rid of parameters and just use argvs directly
func (f FlagStruct) ParseArgs(argv []string) (Options, []string) {

	var (
		long_str_flags  []string
		long_bool_flags []string

		// this will hold the flags and their vals
		options map[string]interface{}

		// this will hold the arguments without flags
		args []string

		// contains the trainslation from SHORT to LONG
		longDict map[string]string

		// don't reuse these flags in the args
		used bool
	)

	// remove filename from the front of the os.Args array
	if len(argv) > 0 && argv[0] == os.Args[0] {
		if len(argv) == 1 {
			// array has no flags to parse
			// i.e. only arguments
			argv = []string{}
		} else {
			// chop off os.Argv[0] (file name)
			argv = argv[1:]
		}
	}

	options = make(Options)

	// add help flag to FlagStruct at parse time
	f.AddBool("h", "help", "Display this message", false)
	options["help"] = false

	// collect keys for parsing
	for _, s := range f.Strings {
		long_str_flags = append(long_str_flags, s.Long)
	}
	for _, b := range f.Bools {
		long_bool_flags = append(long_bool_flags, b.Long)
	}

	// translation from short to long
	// map[short]long
	longDict = getShortLongdict(f)

	// regular expressions to verify flag
	short_regex, _ := regexp.Compile("^-([0-9a-zA-Z]+)$")
	long_regex, _ := regexp.Compile("^--([0-9a-zA-Z]+[0-9a-zA-Z-]*)$")
	short_equal, _ := regexp.Compile("^-([0-9a-zA-Z])=(.+)$")
	long_equal, _ := regexp.Compile("^--([0-9a-zA-Z]+[0-9a-zA-Z-]*)=(.+)$")

	for i := 0; i < len(argv); i++ {

		// -- denotes the end of flag options
		if argv[i] == "--" && i < len(argv)-2 {
			args = append(args, argv[i+1:]...)

			// at this point everything is allocated
			// since there are no more options
			// and all the arguments are stored
			i = len(argv)
			continue
		}

		// ========== SHORTS ==========
		s := short_equal.FindStringSubmatch(argv[i])
		if len(s) == 3 {
			options[longDict[s[1]]] = s[2]
			used = true
		}

		s = short_regex.FindStringSubmatch(argv[i])
		if len(s) == 2 { // group matched!

			// split into "characters"
			short_flags := strings.Split(s[1], "")
			for _, s := range short_flags {

				current_long := longDict[s]
				if contains(long_bool_flags, current_long) {
					options[current_long] = true
					used = true

				} else if len(short_flags) == 1 && i < len(argv)-1 &&
					contains(long_str_flags, current_long) {
					// next elemnet is our string data
					i++
					// can't i++ as a part of this statement
					options[current_long] = argv[i]
					used = true
				}
			}
		}

		// ========== LONGS ==========
		l := long_equal.FindStringSubmatch(argv[i])
		// --example=string
		if len(l) == 3 {
			options[l[1]] = l[2]
			used = true
		}

		l = long_regex.FindStringSubmatch(argv[i])
		// --example string
		if len(l) == 2 { // group matched!
			current_long := l[1]

			// default = store true
			if contains(long_bool_flags, current_long) {
				options[current_long] = true
				used = true

			} else if contains(long_str_flags, current_long) {

				// Warning: greedy
				// if next elemnt is a flag, it will be treated as an argument
				// because this is what the string stores
				if i <= len(argv)-2 {
					// next elemnt is our string dat
					i++
					// can't i++ as a part of this statement
					options[current_long] = argv[i]
					used = true
				}
			}
		}

		// does not belong to a flag group
		// add to arguments
		if !used {
			args = append(args, argv[i])
			used = false
		}

	}
	// at this point command line arguments are parsed
	// now assign the defaults
	for _, j := range f.Bools {
		if options[j.Long] == nil {
			options[j.Long] = j.Default
		}
	}

	for _, j := range f.Strings {
		if options[j.Long] == nil {
			options[j.Long] = j.Default
		}
	}

	return options, args
}

func (f FlagStruct) Usage() {
	fmt.Println()
	fmt.Println("--- Bools ---")
	for _, val := range f.Bools {
		fmt.Printf("-%s, --%-10s\t%s\n", val.Short, val.Long, val.Desc)
	}

	fmt.Println()

	// only display them if they exist
	// no need to waste space
	if len(f.Strings) > 0 {
		fmt.Println("--- Strings ---")
		for _, val := range f.Strings {
			fmt.Printf("-%s, --%-10s\t%s\n", val.Short, val.Long, val.Desc)
		}
	}
}

*/
