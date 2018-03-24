package flag2

type flagProps struct {
	Dest  string
	Short string
	Long  string
	Desc  string
}

type boolFlag struct {
	Props flagProps
	Value bool
}

type stringFlag struct {
	Props flagProps
	Value string
}

type intFlag struct {
	Props flagProps
	Value int
}

type floatFlag struct {
	Props flagProps
	Value float64
}

type Options map[string]interface{}

// this will carry the potential flags a program can have
// now we can use the same receiver for adding
type FlagStruct struct {
	dests   []string
	bools   map[string]boolFlag
	strings map[string]stringFlag
	ints    map[string]intFlag
	floats  map[string]floatFlag
}
