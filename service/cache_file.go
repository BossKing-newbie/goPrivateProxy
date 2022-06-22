package service

import (
	"go_private_proxy/constant"
	"io/ioutil"
	"log"
	"os"
)

type FileCacheInfo struct {
	Id       int             `json:"id"`
	Path     string          `json:"path"`
	Size     int64           `json:"size"` //单位：字节
	ModTime  string          `json:"modTime"`
	IsDir    bool            `json:"isDir"`
	ParentId int             `json:"parentId"`
	Child    []FileCacheInfo `json:"child"`
}

/*生成文件树*/
func ListFileCache(dirName string, parentId int) []FileCacheInfo {
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	var fileCacheInfoList []FileCacheInfo
	pId := parentId * 10
	for _, fi := range fileInfos {
		filename := dirName + string(os.PathSeparator) + fi.Name()
		c := FileCacheInfo{
			ParentId: parentId,
			Path:     filename,
			Size:     fi.Size(),
			ModTime:  constant.GetFormatTime(fi.ModTime()),
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
