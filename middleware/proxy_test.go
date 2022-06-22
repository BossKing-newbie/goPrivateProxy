package middleware

import (
	"fmt"
	"go_private_proxy/service"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestGetProxy(t *testing.T) {
	dirname := "D:\\Go_Workspace\\Demo"
	for _, info := range ListFileCache(dirname, 1) {
		fmt.Println(info)
	}
}
func TestWalk(t *testing.T) {
	err := filepath.WalkDir("D:\\Go_Workspace\\modules", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// handle possible path err, just in case...
			return err
		}
		fmt.Println("skip", path)
		if d.IsDir() {
			fmt.Println("skip", path)
			return fs.SkipDir
		}
		// ... process entry
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
func TestDir(t *testing.T) {
	root := "D:\\Go_Workspace\\modules"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Println("dir:", path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

}
func funcWalkFile(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	return files, err
}
func GetDir(dir string) {
	f, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		if file.IsDir() {
			GetDir(dir + "//" + file.Name())
		}
	}
}
func readDirFile(dir string) []string {
	var allFiles []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			readDirFile(dir + "/" + file.Name())
		} else {
			allFiles = append(allFiles, dir+"/"+file.Name())
		}
	}
	return allFiles
}
func listFiles(dirname string, level int) {
	// level用来记录当前递归的层次
	// 生成有层次感的空格
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|   " + s
	}
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		filename := dirname + fmt.Sprintf("%c", os.PathSeparator) + fi.Name()
		fmt.Printf("%s%s\n", s, filename)
		if fi.IsDir() {
			//继续遍历fi这个目录
			listFiles(filename, level+1)
		}
	}
}
func ListFileCache(dirName string, parentId int) []service.FileCacheInfo {
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	var fileCacheInfoList []service.FileCacheInfo
	pId := parentId * 10
	for _, fi := range fileInfos {
		filename := dirName + fmt.Sprintf("%c", os.PathSeparator) + fi.Name()
		c := service.FileCacheInfo{
			ParentId: parentId,
			Path:     filename,
			Size:     fi.Size(),
			ModTime:  fi.ModTime(),
			Id:       pId,
			IsDir:    false,
		}
		if fi.IsDir() {
			//继续遍历目录
			c.IsDir = true
			c.Child = ListFileCache(filename, c.Id)
		} else {
			c.IsDir = false
		}
		pId++
		fileCacheInfoList = append(fileCacheInfoList, c)
	}
	return fileCacheInfoList
}
