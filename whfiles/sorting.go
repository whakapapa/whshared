package whfiles

// provide sorting interfaces for struct TfullFile


// sort by name
type sortName []TfullFile

func (f sortName)	Len() int {
	return len(f)
}

func (f sortName)	Swap(x, y int) {
	f[x], f[y] = f[y], f[x]
}

func (f sortName)	Less(x, y int) bool {
	return f[x].Name < f[y].Name
}

// path
type sortPath	[]TfullFile

func (f sortPath)	Len() int {
	return len(f)
}

func (f sortPath)	Swap(x, y int) {
	f[x], f[y] = f[y], f[x]
}

func (f sortPath)	Less(x, y int) bool {
	return f[x].Name < f[y].Name
}
