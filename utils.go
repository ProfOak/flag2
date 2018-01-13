package flag2

func listContains(strings []string, elem string) bool {
	for _, s := range strings {
		if s == elem {
			return true
		}
	}
	return false
}

/*
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
*/
