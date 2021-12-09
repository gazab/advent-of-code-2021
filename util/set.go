package util

func Intersection(s1 map[string]bool, s2 map[string]bool) map[string]bool {
	s_intersection := map[string]bool{}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1 // better to iterate over a shorter set
	}
	for k := range s1 {
		if s2[k] {
			s_intersection[k] = true
		}
	}

	return s_intersection
}

func Union(s1 map[string]bool, s2 map[string]bool) map[string]bool {
	s_union := map[string]bool{}
	for k := range s1 {
		s_union[k] = true
	}
	for k := range s2 {
		s_union[k] = true
	}

	return s_union
}

func Difference(s1 map[string]bool, s2 map[string]bool) map[string]bool {
	s_difference := map[string]bool{}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1 // better to iterate over a shorter set
	}
	for k := range s1 {
		if !s2[k] {
			s_difference[k] = true
		}
	}

	return s_difference
}
