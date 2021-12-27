// @File:  update_crontab.go
// @Time:  2022/1/5 1:11 PM
// @Author: ClassmateLin
// @Email: classmatelin.site@gmail.com
// @Site: https://www.classmatelin.top
// @Description:
// @Cron: * */1 * * *
package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"regexp"
	"scripts/com"
	"scripts/global"
	"strings"
)

const defaultCrontabFile = ".crontab.json"
const publishedCrontabFile = "config/crontab.json"

// updateDefaultCrontab
// @description: 更新默认的crontab
// @param:  scriptList
func updateDefaultCrontabConf(scriptList []string) {

	crontabMap := map[string]string{}

	r, err := regexp.Compile("@Cron:(.*)")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, script := range scriptList {
		_, name := filepath.Split(script)
		name = strings.Split(name, ".")[0]
		content, err := ioutil.ReadFile(script)
		if err != nil {
			continue
		}
		res := r.FindString(string(content))
		if res == "" {
			continue
		}
		data := strings.Split(res, ":")
		crontab := data[1]
		crontabMap[name] = strings.TrimSpace(crontab)
	}

	data, err := json.MarshalIndent(crontabMap, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(defaultCrontabFile, data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// updatePublishedCrontab
// @description: 更新在使用的crontab.json
// @param:  scriptList
func updatePublishedCrontabConf() {
	defaultVp := viper.New()
	defaultVp.SetConfigFile(defaultCrontabFile)

	if err := defaultVp.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}

	defaultCrontab := defaultVp.AllSettings()

	publishedVp := viper.New()
	publishedVp.SetConfigFile(publishedCrontabFile)

	if err := publishedVp.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}

	publishedCrontab := publishedVp.AllSettings()

	for key, val := range defaultCrontab {
		_, ok := publishedCrontab[key]
		if ok {
			continue
		} else {
			publishedCrontab[key] = val
		}
	}

	data, err := json.MarshalIndent(publishedCrontab, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(publishedCrontabFile, data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// updateCrontab
// @description: 更新crontab定时任务
func updateCrontab() {
	cronVp := viper.New()
	cronVp.SetConfigFile(publishedCrontabFile)

	if err := cronVp.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}

	cronMap := cronVp.AllSettings()

	absPath, err := filepath.Abs(".")
	if err != nil {
		return
	}

	crontabList := []string{
		"30 11 * * * docker-entrypoint",
	}

	for name, cron := range cronMap {
		if cron == false { // 已关闭脚本
			continue
		}
		itemList := strings.Split(strings.TrimSpace(cron.(string)), " ")
		if len(itemList) > 5 { // 错误的配置
			continue
		}

		crontabString := fmt.Sprintf("%s cd %s && %s.bin", cron, absPath, filepath.Join(absPath, name))
		crontabList = append(crontabList, crontabString)
	}

	crontabText := "SHELL=/bin/sh\n\n"
	for _, crontab := range crontabList {
		crontabText += crontab + "\n\n"
	}
	fmt.Println(crontabText)
	err = ioutil.WriteFile("/tmp/crontab", []byte(crontabText), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	cmd := exec.Command("crontab", "/tmp/crontab")

	err = cmd.Run()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return
	}
}

func main() {
	var scriptList []string
	com.GetGoFileList(global.DefaultScriptPath, &scriptList)
	updateDefaultCrontabConf(scriptList)
	updatePublishedCrontabConf()
	updateCrontab()
}
