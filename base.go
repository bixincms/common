package common

func SliceStringUnique(s []string) (unique_silce_string []string) {
	_s := map[string]struct{}{}
	for _, v := range s {
		if _, isok := _s[v]; isok {
			unique_silce_string = append(unique_silce_string, v)
		}
	}
	return
}
