package whfiles

// provide sorting interfaces for struct T_FullFile


// sort by name
type T_SortByName []T_FullFile

func (f T_SortByName)	Len() int {
	return len(f)
}

func (f T_SortByName)	Swap(x, y int) {
	f[x], f[y] = f[y], f[x]
}

func (f T_SortByName)	Less(x, y int) bool {
	return f[x].Name < f[y].Name
}

// path
type T_SortByPath	[]T_FullFile

func (f T_SortByPath)	Len() int {
	return len(f)
}

func (f T_SortByPath)	Swap(x, y int) {
	f[x], f[y] = f[y], f[x]
}

func (f T_SortByPath)	Less(x, y int) bool {
	return f[x].Name < f[y].Name
}
