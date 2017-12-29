package main

import (
	"github.com/LK4D4/vndr/build"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	ignoreFile    = "vndr.ignore"
	ignoreFolders = "ignoreFolders"
	ignorePaths   = "ignorePaths"
)

func readIgnoreFile() (ignoreFolders []string, ignorePaths []string) {
	_, err := os.Stat(ignoreFile)
	if err != nil {
		log.Println("Missing vndr.ignore file skip this step")
		return
	}
	fileBytes, err := ioutil.ReadFile(ignoreFile)
	if err != nil {
		log.Printf("Read %s file error %v", ignoreFile, err)
		return
	}
	fileContent := string(fileBytes)
	fileLines := strings.Split(fileContent, ":")
	if len(fileLines) != 3 {
		return
	}
	for _, value := range strings.Split(fileLines[1], "\n") {
		if strings.Contains(value, "-") {
			tmpStr := strings.Replace(value, "-", "", 1)
			tmpStr = strings.TrimSpace(tmpStr)
			ignoreFolders = append(ignoreFolders, tmpStr)
		}
	}

	for _, value := range strings.Split(fileLines[2], "\n") {
		if strings.Contains(value, "-") {
			tmpStr := strings.Replace(value, "-", "", 1)
			tmpStr = strings.TrimSpace(tmpStr)
			ignorePaths = append(ignorePaths, tmpStr)
		}
	}

	return
}

func removeIgnoreImports(p *build.Package) *build.Package {
	_, ignorePaths := readIgnoreFile()
	var tmpImportsSlice []string
	tmpImportsFlag := false
	for _, tmpImport := range p.Imports {
		for _, tmpIgnorePath := range ignorePaths {
			if strings.HasPrefix(tmpImport, tmpIgnorePath) {
				tmpImportsFlag = true
				break
			}
		}
		if tmpImportsFlag {
			log.Printf("Ignore import %s", tmpImport)
			tmpImportsFlag = false
			continue
		}
		tmpImportsSlice = append(tmpImportsSlice, tmpImport)
	}
	p.Imports = tmpImportsSlice
	return p
}

func isIgnorePath(path string) bool {
	_, ignorePaths := readIgnoreFile()
	for _, value := range ignorePaths {
		if strings.Contains(path, value) {
			return true
		}
	}
	return false
}

func addWhilteList(whiteList regexpSlice) regexpSlice {
	_, ignorePaths := readIgnoreFile()
	for _, value := range ignorePaths {
		whiteList.Set(value)
	}
	return whiteList
}