package flag2

type boolFlag struct {
	Short string
	Long  string
	Desc  string
	Val   bool
}

type stringFlag struct {
	Short string
	Long  string
	Desc  string
	Val   string
}

type Options map[string]interface{}

// this will carry the potential flags a program can have
// now we can use the same receiver for adding
type FlagStruct struct {
	Bools   map[string]boolFlag
	Strings map[string]stringFlag
}
