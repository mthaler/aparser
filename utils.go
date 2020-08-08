package aparser

func delete(a []interface{}, i int) []interface{} {
	return append(a[:i], a[i+1:]...)
}

func deleteRange(a []interface{}, i, c int) []interface{} {
	return append(a[:i], a[i+c:]...)
}
