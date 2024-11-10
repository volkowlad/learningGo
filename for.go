package main

func Map(strs []string, mapFunc func(s string) string) []string {
	newstrs := make([]string, len(strs))
	for i, s := range strs {
		newstrs[i] = mapFunc(s)
	}

	return newstrs
}
