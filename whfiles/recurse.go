package whfiles


import (
	"os"
	"os/user"
	"log"
	"sort"
	"path/filepath"
)


type TfullFile struct {
	Name		string
	Path		string
	IsDir		bool
}


func setHomeDir() string {
	buddy, err := user.Current()
	if err != nil {
		log.Println(err)
	}
	return buddy.HomeDir
}


// deliver full directory content in TfullFile struct
// to be evaluated and stripped in caller
func ReadDirContent(dirPath string) ([]TfullFile) {

	currentDir, err := os.Open(dirPath)
	if err != nil {
		log.Println(err)
	}
	defer currentDir.Close()

	// catch all items from the current directory
	allContent, err := currentDir.Readdir(0)
	if err != nil {
		log.Println(err)
	}

	// transfer into TfullFile struct
	var allItems []TfullFile
	for i := range allContent {
		// translate into own TfullFile struct
		var tmpItem TfullFile

		tmpItem.Path = dirPath
		tmpItem.Name = allContent[i].Name()

		switch {
		case allContent[i].IsDir():
			tmpItem.IsDir = true
		default:
			tmpItem.IsDir = false
		}
		allItems = append(allItems, tmpItem)
	}
	return allItems
}


func CatalogByPattern(allItems []TfullFile, regPattern string) ([]TfullFile, []string) {

	var resultCatalog []TfullFile
	var parseDirs []string

	// filter results based on regPattern
	for i := range allItems {
		keepItem, err := filepath.Match(regPattern, allItems[i].Name)
		if err != nil {
			log.Println(err)
		}
		if keepItem {
			resultCatalog = append(resultCatalog, allItems[i])
		}
	}

	// return all newly found dirs for further parsing
	for i := range allItems {
		if allItems[i].IsDir {
			parseDirs = append(parseDirs, allItems[i].Name)
		}
	}
	return resultCatalog, parseDirs
}


func BuildFullCatalog(dirPath string, kinds int, recurse bool, regPattern string) ([]TfullFile, int, int) {
	// kinds are for now: 0: dirs, 1: files, 2: both
	var fullList []TfullFile
	var remainingDirs []string

	// check if item exists in fs
	// fallback user home directory
	ref, err := os.Stat(dirPath)
	if err != nil {
		log.Println(err)
		log.Println("path or file does not exist, using home directory instead")
		dirPath = setHomeDir() + "/"
		log.Println(dirPath)
		// reopen the connection
		ref, _ = os.Stat(dirPath)
	}

	// if item exists, but is file not dir
	// construct base dir from string
	if ref.IsDir() == false {
		log.Println(dirPath, "is a file, constructing parent directory")
		// find the last forward slash
		pos := len(dirPath)
		for pos > 0 && dirPath[pos-1:pos] != "/" {
			pos -= 1
		}
		// strip pathfile
		tmpPath := dirPath[0:pos]
		log.Println("constructed path is", tmpPath)
		dirPath = tmpPath
	}

	remainingDirs = append(remainingDirs, dirPath)

	for len(remainingDirs) > 0 {
		newItems, newDirs := CatalogByPattern(ReadDirContent(remainingDirs[0]), regPattern)

		for i := range newItems {
			// add new items to fullList
			fullList = append(fullList, newItems[i])
		}

		if recurse {

			// extend remainingDirs for further parsing
			for i := range newDirs {
				parseDir := remainingDirs[0] + newDirs[i] + "/"
				remainingDirs = append(remainingDirs, parseDir)
			}

			// now strip last directory from remainingDirs
			remainingDirs = remainingDirs[1:]
		}
	}

	// now keep only wanted items (files and/or directories)
	var wantedItems []TfullFile

	switch {
	case kinds == 2:
		wantedItems = fullList
	default:
		for i := range fullList {
			switch {
				// directories only
			case kinds == 0 && fullList[i].IsDir:
				wantedItems = append(wantedItems, fullList[i])
				// files only
			case kinds == 1 && fullList[i].IsDir == false:
				wantedItems = append(wantedItems, fullList[i])
			}
		}
	}


	// finally sort by name and path
	sort.Sort(SortName(wantedItems))
	sort.Sort(SortPath(wantedItems))
	return wantedItems, cDirs, cFiles
}


func CountFiles(cat []TfullFile) int {
	// provide simple stats
	var cFiles int
	for i := range wantedItems {
		if cat[i].IsDir == false {
			cFiles += 1
		}
	}
	return cFiles
}


func CountDirs(cat []TfullFile) int {
	// provide simple stats
	var cDirs int
	for i := range wantedItems {
		if cat[i].IsDir {
			cDirs += 1
		}
	}
	return cDirs
}
