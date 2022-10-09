package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	MaxDeep    = 0
	uuid       = 0
	MaxLeafLen = 2
)

type FileTree struct {
	ID        string      `json:"id"`
	Label     string      `json:"lab"`
	Children  []*FileTree `json:"children"`
	Collapsed bool        `json:"collapsed"`
	Dir       bool        `json:"dir"`
}

func newFileTree(label string) *FileTree {
	id := fmt.Sprintf("id_%d", uuid)
	uuid += 1
	return &FileTree{id, label, []*FileTree{}, true, true}
}
func main() {
	var fileName string
	flag.StringVar(&fileName, "file", "sample.txt", "請輸入檔名")
	flag.Parse()
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fileTree := newFileTree("root")
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		textArr := strings.Split(strings.TrimLeft(scanner.Text(), "./"), "/")
		RecurParse(fileTree, textArr, 0)
	}
	SaveFile(fileTree)
	fmt.Printf("done\n ")
}
func SaveFile(ft *FileTree) {
	sb := strings.Builder{}
	sb.WriteString("data=")
	barr, _ := json.Marshal(ft)
	sb.Write(barr)
	os.WriteFile("web/out.json", []byte(sb.String()), 0644)
}
func ParseTree(input []byte) FileTree {
	var ft FileTree
	err := json.Unmarshal(input, &ft)
	if err != nil {
		panic("fail to parse")
	}
	return ft

}
func RecurParse(root *FileTree, path []string, deep int) {
	if len(path) == 0 {
		return
	}
	nRoot := contains(root.Children, path[0])
	if nRoot == nil {
		nRoot = newFileTree(path[0])
		if !strings.HasSuffix(path[0], "/") && len(path) == 1 && len(root.Children) > MaxLeafLen {
			return
		}
		root.Children = append(root.Children, nRoot)
	}
	if deep > MaxDeep {
		nRoot.Collapsed = true
	}
	if len(path) > 1 && len(path[1]) > 0 {
		ndeep := deep + 1
		RecurParse(nRoot, path[1:], ndeep)
	} else {
		nRoot.Dir = false
	}

}
func contains(arr []*FileTree, elem string) *FileTree {
	for _, item := range arr {
		if item.Label == elem {
			return item
		}
	}
	return nil
}
