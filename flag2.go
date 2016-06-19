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
		return fmt.Errorf("Atribute already exists in list of bools: %s", long)
	}

	return nil
}

func (f *FlagStruct) AddString(short string, long string, desc string, val string) error {

	// check if it doesn't exist
	if f.Strings[long] == (stringFlag{}) {
		f.Strings[long] = stringFlag{short, long, desc, val}
	} else {
		// not empty, and already exists == no good
		return fmt.Errorf("Atribute already exists in list of strings: %s", long)
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
	f.AddBool("h", "help", "Display this message and exit", true)
	options["help"] = true

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
	short_regex, _ := regexp.Compile("^-([\\w]*)+$")
	long_regex, _ := regexp.Compile("^--([^-]+)$")

	for i := 0; i < len(argv); i++ {

		// ========== SHORTS ==========
		s := short_regex.FindStringSubmatch(argv[i])
		if len(s) > 1 { // group matched!

			// split into "characters"
			short_flags := strings.Split(s[1], "")
			for _, s := range short_flags {

				current_long := longDict[s]
				if contains(long_bool_flags, current_long) {
					options[current_long] = true

				} else if len(short_flags) == 1 && i < len(argv)-1 &&
					contains(long_str_flags, current_long) {
					// next elemnt is our string dat
					i++
					// can't i++ as a part of this statement
					options[current_long] = argv[i]
				}
			}
		}

		// ========== LONGS ==========
		l := long_regex.FindStringSubmatch(argv[i])
		if len(l) > 1 { // group matched!
			current_long := l[1]

			// default = store true
			if contains(long_bool_flags, current_long) {
				options[current_long] = true

			} else if contains(long_str_flags, current_long) {
				// Warning: greedy
				// if next elemnt is a flag, it will be treated as an argument
				// because this is what the string stores
				if i <= len(argv)-1 {
					// next elemnt is our string dat
					i++
					// can't i++ as a part of this statement
					options[current_long] = argv[i]
				}
			}
		}

		// does not belong to a flag group
		// add to arguments
		if len(s) == 0 && len(l) == 0 {
			args = append(args, argv[i])
		}

	}
	return options, args
}

func (f FlagStruct) Usage() {
	fmt.Println()
	fmt.Println("--- Bools ---")
	for _, val := range f.Bools {
		fmt.Printf("-%s, --%s\t%s\n", val.Short, val.Long, val.Desc)
	}

	fmt.Println()
	fmt.Println("--- Strings ---")
	for _, val := range f.Strings {
		fmt.Printf("-%s, --%s\t%s\n", val.Short, val.Long, val.Desc)
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
