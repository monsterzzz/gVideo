package fileUtil

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type FileName []os.FileInfo

func (f FileName) Len() int { return len(f) }

func (f FileName) Less(i, j int) bool {
	f1Index, _ := strconv.Atoi(strings.Split(strings.ReplaceAll(f[i].Name(), ".tmp", ""), "_")[1])
	f2Index, _ := strconv.Atoi(strings.Split(strings.ReplaceAll(f[j].Name(), ".tmp", ""), "_")[1])
	return f1Index < f2Index
}

func (f FileName) string() string {
	var s []string
	for _, v := range f {
		s = append(s, v.Name())
	}
	return strings.Join(s, "")
}

func (f FileName) Swap(i, j int) {

	f[i], f[j] = f[j], f[i]

}

func MergeFiles(dirPath string, fileType string) {
	var fNList FileName
	fNList, e := ioutil.ReadDir(dirPath)
	sort.Sort(fNList)
	if e != nil {
		fmt.Println(e)
	}
	fString := strings.Split(dirPath, "/")
	md5 := fString[len(fString)-1]
	endFile, _ := os.Create("./videoRepo/fullFile/" + md5 + "." + fileType)
	defer endFile.Close()
	for _, f := range fNList {
		src, err := os.Open(dirPath + "/" + f.Name())
		fmt.Println(f.Name())
		if err != nil {
			fmt.Println(err)
		}
		ipt, _ := ioutil.ReadAll(src)
		endFile.Write(ipt)
		src.Close()
	}
	err := os.RemoveAll(dirPath)
	if err != nil {
		fmt.Println(err)
	}
}
