// @File:  build_scripts.go
// @Time:  2022/1/9 20:21
// @Author: ClassmateLin
// @Email: classmatelin.site@gmail.com
// @Site: https://www.classmatelin.top
// @Description: 编译脚本

package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"scripts/com"
	"scripts/global"
	"strings"
)

func main() {
	var scriptList []string
	com.GetGoFileList(global.DefaultScriptPath, &scriptList)

	absRootPath, err := filepath.Abs(".")
	if err != nil {
		return
	}

	fmt.Printf("开始编译%d个脚本...\n", len(scriptList))
	for _, scriptPath := range scriptList {
		_, filename := filepath.Split(scriptPath)
		output := strings.Split(filename, ".")[0] + ".bin"

		cmd := exec.Command("go", "build", "-o", filepath.Join(absRootPath, output), scriptPath)

		err := cmd.Run()
		if err != nil {
			fmt.Println("Execute Command failed:" + err.Error())
			return
		}
	}
	fmt.Println("脚本编译完成...")
}
