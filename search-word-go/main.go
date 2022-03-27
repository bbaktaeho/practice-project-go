package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	name  string
	lines []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("At least 2 arguments are required. ex) 'search-word-go word filename1 filename2 ...'")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	findInfos := []FindInfo{}
	fmt.Println("찾으려는 단어:", word)
	fmt.Println()

	for _, path := range files {
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}

	for _, findInfo := range findInfos {
		fmt.Println(findInfo.name)
		fmt.Println("-----------------------------------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("-----------------------------------------------------")
		fmt.Println()
	}
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	fileList, err := GetFileList(path)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. err:", err)
		return findInfos
	}

	for _, name := range fileList {
		findInfos = append(findInfos, FindWordInFile(word, name))
	}
	return findInfos
}

func FindWordInFile(word, name string) FindInfo {
	findInfo := FindInfo{name, []LineInfo{}}
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. err:", err)
		return findInfo
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	return findInfo
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

// func Exists(name string) (bool, error) {
// 	_, err := os.Stat(name)
// 	if os.IsNotExist(err) {
// 		return false, err
// 	}
// 	return err == nil, nil
// }

// func PrintAllFiles(files []string) {
// 	fmt.Println("[찾으려는 파일 리스트]")

// 	for _, name := range files {
// 		ok, err := Exists(name)
// 		if !ok {
// 			fmt.Printf("%s 파일은 존재하지 않습니다.\n", name)
// 			fmt.Println("err:", err)
// 			continue
// 		}
// 		fmt.Println(name)
// 	}
// }

// func PrintFile(name string) {
// 	file, err := os.Open(name)
// 	if err != nil {
// 		fmt.Printf("%s 파일은 존재하지 않습니다.\n", name)
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}
// }
