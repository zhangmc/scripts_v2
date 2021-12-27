// @File:  file.go
// @Time:  2022/1/9 20:28
// @Author: ClassmateLin
// @Email: classmatelin.site@gmail.com
// @Site: https://www.classmatelin.top
// @Description:

package com

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// GetGoFileList
// @description: 获取指定目录下的go文件列表
// @param:  baseDir
// @param:  fileList
func GetGoFileList(baseDir string, fileList *[]string) {
	absBaseDir, err := filepath.Abs(baseDir)

	if err != nil {
		fmt.Println(err)
		return
	}
	itemList, err := ioutil.ReadDir(absBaseDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range itemList {
		if item.IsDir() {
			GetGoFileList(filepath.Join(absBaseDir, item.Name()), fileList)
			continue
		} else {
			if strings.HasSuffix(item.Name(), ".go") {
				*fileList = append(*fileList, filepath.Join(absBaseDir, item.Name()))
			}

		}
	}
}
