package flag2

import (
	"testing"
)

// proper flag values
const (
	metavar     = "meta"
	short       = "s"
	long        = "long"
	desc        = "help description"
	boolValue   = true
	stringValue = "string"
	intValue    = 1234
	floatValue  = 1234.5678
)

func TestNew(t *testing.T) {
	var f FlagStruct

	if f.bools != nil {
		t.Error("Bools should be nil before initializing")
	} else if f.strings != nil {
		t.Error("Strings should be nil before initializing")
	} else if f.ints != nil {
		t.Error("Ints should be nil before initializing")
	} else if f.bools != nil {
		t.Error("floats should be nil before initializing")
	}

	f = New()

	if f.bools == nil {
		t.Error("Bools should not be after before initializing")
	} else if f.strings == nil {
		t.Error("Strings should not be nil after initializing")
	} else if f.ints == nil {
		t.Error("Ints should not be nil after initializing")
	} else if f.floats == nil {
		t.Error("Floats should not be nil after initializing")
	}
}

func TestValidateFlag(t *testing.T) {
	var (
		f     FlagStruct
		p     flagProps
		props flagProps
		err   error
	)
	f = New()
	props = flagProps{metavar, short, long, desc}
	p, err = f.validateFlag(metavar, short, long, desc)

	// check for valid flag to fail
	if err != nil ||
		p != props ||
		p.Metavar != metavar ||
		p.Short != short ||
		p.Long != long ||
		p.Desc != desc {

		t.Error("Real flag is invalid")
	}

	// check for duplicity
	f.bools["metavar"] = boolFlag{
		Props: props,
		Value: true,
	}
	f.metavars = append(f.metavars, metavar)
	_, err = f.validateFlag(metavar, short, long, desc)
	if err == nil {
		t.Error("Allowed adding multiple metavars")
	}

	// check for other bad flag structures
	bad_flags := [][]string{
		{"", "s", "long", "help"},            // bad metavar
		{"meta", "", "", "help"},             // no flags
		{"meta", "notshort", "long", "help"}, // short too long
		{"meta", "s", "l o n g", "help"},     // spaces in flag name
	}
	for _, flag_set := range bad_flags {
		f = New()
		_, err = f.validateFlag(flag_set[0], flag_set[1], flag_set[2], flag_set[3])
		if err == nil {
			t.Errorf("%V: Invalid flag passed validation", flag_set)
		}
	}
}

func TestAddBool(t *testing.T) {
	f := New()
	f.AddBool(metavar, short, long, desc)
	if f.metavars[0] != metavar {
		t.Error("Metavar slice not being set")
	} else if _, exists := f.bools[metavar]; !exists {
		t.Error("Flag not added to bools map")
	}
}

func TestAddString(t *testing.T) {
	f := New()
	f.AddString(metavar, short, long, desc, stringValue)
	if f.metavars[0] != metavar {
		t.Error("Metavar slice not being set")
	} else if _, exists := f.strings[metavar]; !exists {
		t.Error("Flag not added to strings map")
	}
}

func TestAddInt(t *testing.T) {
	f := New()
	f.AddInt(metavar, short, long, desc, intValue)
	if f.metavars[0] != metavar {
		t.Error("Metavar slice not being set")
	} else if _, exists := f.ints[metavar]; !exists {
		t.Error("Flag not added to ints map")
	}
}

func TestAddFloat(t *testing.T) {
	f := New()
	f.AddFloat(metavar, short, long, desc, floatValue)
	if f.metavars[0] != metavar {
		t.Error("Metavar slice not being set")
	} else if _, exists := f.floats[metavar]; !exists {
		t.Error("Flag not added to floats map")
	}
}

func TestContains(t *testing.T) {
	values := []string{"a", "b", "c"}
	for _, val := range values {
		if !listContains(values, val) {
			t.Error("Value not found: %s", val)
		}
	}
}
