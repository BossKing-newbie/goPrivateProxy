package service

import (
	"go_private_proxy/constant"
	"strings"
)

type VersionStatistic struct {
	Module        string `json:"module"`
	Version       string `json:"version"`
	DownloadCount int64  `json:"downloadCount"`
}

func GetVersionList() []VersionStatistic {
	versions := constant.GetConcurrentMap().Keys()
	result := make([]VersionStatistic, len(versions))
	for _, version := range versions {
		if len(version) == 0 || !strings.Contains(version, ":") {
			continue
		}
		mod := strings.Split(version, ":")[0]
		ver := strings.Split(version, ":")[1]
		count, _ := constant.GetConcurrentMap().Get(version)
		v := VersionStatistic{
			Module:        mod,
			Version:       ver,
			DownloadCount: count.(int64),
		}
		result = append(result, v)
	}
	return result
}
func GetModList() []VersionStatistic {
	versionList := GetVersionList()
	vMap := make(map[string]int64, 10)
	result := make([]VersionStatistic, len(versionList))
	for _, v := range versionList {
		if len(v.Module) == 0 {
			continue
		}
		if count, ok := vMap[v.Module]; ok {
			vMap[v.Module] = count + 1
		} else {
			vMap[v.Module] = 1
		}
	}
	for k, v := range vMap {
		r := VersionStatistic{
			Module:        k,
			DownloadCount: v,
		}
		result = append(result, r)
	}
	return result
}
