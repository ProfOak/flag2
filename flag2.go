package flag2

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func NewFlag() FlagStruct {
	// kind of like a constructor
	f := FlagStruct{}
	f.Bools = make(map[string]boolFlag)
	f.Strings = make(map[string]stringFlag)
	return f
}

func (f *FlagStruct) AddBool(short string, long string, desc string, val bool) error {

	// check if it doesn't exist
	if f.Bools[long] == (boolFlag{}) {
		f.Bools[long] = boolFlag{short, long, desc, val}
	} else {
		// not empty, and already exists == no good
		return fmt.Errorf("[FlagError] Flag already exists in list of bools %s", long)
	}

	return nil
}

func (f *FlagStruct) AddString(short string, long string, desc string, val string) error {

	// check if it doesn't exist
	if f.Strings[long] == (stringFlag{}) {
		f.Strings[long] = stringFlag{short, long, desc, val}
	} else {
		// not empty, and already exists == no good
		return fmt.Errorf("[FlagError] Flag already exists in list of strings: %s", long)
	}

	return nil
}

func (f FlagStruct) FlagKeys() []string {
	// keys = list of long flag names

	var keys []string

	for _, s := range f.Strings {
		keys = append(keys, s.Long)
	}

	for _, s := range f.Bools {
		keys = append(keys, s.Long)
	}

	return keys
}

func (f FlagStruct) Parse(argv []string) (Options, []string) {

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

// ===== HELPER FUNCTIONS =====

func contains(strings []string, elem string) bool {
	for _, s := range strings {
		if s == elem {
			return true
		}
	}
	return false
}

func getShortLongdict(f FlagStruct) map[string]string {
	dict := make(map[string]string, 0)

	for _, flag := range f.Strings {
		dict[flag.Short] = flag.Long
	}
	for _, flag := range f.Bools {
		dict[flag.Short] = flag.Long
	}
	return dict
}
