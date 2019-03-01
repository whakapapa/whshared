package whfiles

// provide sorting interfaces for struct TfullFile


// sort by name
type SortName []TfullFile

func (f SortName)	Len() int {
	return len(f)
}

func (f SortName)	Swap(x, y int) {
	f[x], f[y] = f[y], f[x]
}

func (f SortName)	Less(x, y int) bool {
	return f[x].Name < f[y].Name
}

// path
type SortPath	[]TfullFile

func (f SortPath)	Len() int {
	return len(f)
}

func (f SortPath)	Swap(x, y int) {
	f[x], f[y] = f[y], f[x]
}

func (f SortPath)	Less(x, y int) bool {
	return f[x].Name < f[y].Name
}
