// @File:  update_config.go
// @Time:  2022/1/5 1:10 PM
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
	"reflect"
)

const publishedConfigFile = "config/config.json"
const defaultConfigFile = ".config.json"

// updateConfig
// @description: 更新配置文件
func updateConfig() {
	defaultVp := viper.New()
	defaultVp.SetConfigFile(defaultConfigFile)

	if err := defaultVp.ReadInConfig(); err != nil {
		fmt.Printf("读取配置文件:%s出错, %s", defaultConfigFile, err)
		return
	}

	publishedVp := viper.New()
	publishedVp.SetConfigFile(publishedConfigFile)

	if err := publishedVp.ReadInConfig(); err != nil {
		fmt.Printf("读取配置文件:%s出错, %s", publishedConfigFile, err)
		return
	}

	result := mergeConfig(publishedVp.AllSettings(), defaultVp.AllSettings())

	data, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(publishedConfigFile, data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("成功合并配置文件, [%s, %s]", publishedConfigFile, defaultConfigFile)
}

// mergeConfig
// @description: 合并配置文件
// @param:  publishedConfig 已发布的配置文件
// @param:  defaultConfig 默认的配置文件
// @return: map[string]interface{}
func mergeConfig(publishedConfig map[string]interface{}, defaultConfig map[string]interface{}) map[string]interface{} {
	for key, val := range defaultConfig {
		if reflect.TypeOf(val) == reflect.TypeOf(map[string]interface{}{}) {
			m, ok := publishedConfig[key]
			if !ok {
				publishedConfig[key] = val
			} else {
				val = mergeConfig(m.(map[string]interface{}), val.(map[string]interface{}))
				publishedConfig[key] = val
			}
			continue
		}
		_, ok := publishedConfig[key]
		if ok {
			continue
		} else {
			publishedConfig[key] = val
		}
	}
	return publishedConfig
}

func main() {
	updateConfig()
}
