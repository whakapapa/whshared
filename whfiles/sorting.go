package whfiles

// provide sorting interfaces for struct tFullFile


// sort by name
type sortName []tFullFile

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
type sortPath	[]tFullFile

func (f sortPath)	Len() int {
	return len(f)
}

func (f sortPath)	Swap(x, y int) {
	f[x], f[y] = f[y], f[x]
}

func (f sortPath)	Less(x, y int) bool {
	return f[x].Name < f[y].Name
}


func setHomeDir() string {
	buddy, err := user.Current()
	if err != nil {
		log.Println(err)
	}
	return buddy.HomeDir
}
