package whfiles


import (
	"os"
	"os/user"
	"log"
	"sort"
	"path/filepath"
	"time"
)


type T_FullFile struct {
	Name		string
	Path		string
	IsDir		bool
	Size		int64
	Mod		time.Time
}


func CheckError (e err) {
	if e != nil {
		log.Println(e)
	}
}


func setHomeDir() string {
	buddy, err := user.Current()
	CheckError(err)

	return buddy.HomeDir
}


// deliver full directory content in T_FullFile struct
// to be evaluated and stripped in caller
func ReadDirContent(dirPath string) ([]T_FullFile) {

	currentDir, err := os.Open(dirPath)
	CheckError(err)
	defer currentDir.Close()

	// catch all items from the current directory
	allContent, err := currentDir.Readdir(0)
	CheckError(err)

	// transfer into T_FullFile struct
	var allItems []T_FullFile
	for _, v := range allContent {
		// translate into own T_FullFile struct
		var tmpItem T_FullFile

		tmpItem.Path = dirPath
		tmpItem.Name = v.Name()

		switch {
		case v.IsDir():
			tmpItem.IsDir = true
		default:
			tmpItem.IsDir = false
		}

		tmpItem.Size = v.Size()
		tmpItem.Mod = v.ModTime()
		
		allItems = append(allItems, tmpItem)
	}
	return allItems
}


func CatalogByPattern(allItems []T_FullFile, regPattern string) ([]T_FullFile, []string) {

	var resultCatalog []T_FullFile
	var parseDirs []string

	// filter results based on regPattern
	for i := range allItems {
		keepItem, err := filepath.Match(regPattern, allItems[i].Name)
		CheckError(err)
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


func BuildFullCatalog(dirPath string, kinds int, recurse bool, regPattern string) []T_FullFile {
	// kinds are for now: 0: dirs, 1: files, 2: both
	var fullList []T_FullFile
	var remainingDirs []string

	// check if item exists in fs
	// fallback user home directory
	ref, err := os.Stat(dirPath)
	CheckError(err)

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
	var wantedItems []T_FullFile

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
	sort.Sort(T_SortByName(wantedItems))
	sort.Sort(T_SortByPath(wantedItems))
	return wantedItems
}


func CountFiles(cat []T_FullFile) int {
	// provide simple stats
	var cFiles int
	for i := range cat {
		if cat[i].IsDir == false {
			cFiles += 1
		}
	}
	return cFiles
}


func CountDirs(cat []T_FullFile) int {
	// provide simple stats
	var cDirs int
	for i := range cat {
		if cat[i].IsDir {
			cDirs += 1
		}
	}
	return cDirs
}
