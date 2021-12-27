// @File:  base.go
// @Time:  2022/1/5 12:24 PM
// @Author: ClassmateLin
// @Email: classmatelin.site@gmail.com
// @Site: https://www.classmatelin.top
// @Description:
// @Cron: * */1 * * *

package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const DefaultConfigPath = "config"

/**
type Config struct {
	Notify Notify `json:"notify"`
}

type Notify struct {
	ServerJ    string `json:"server_j"`
	PushPlus   string `json:"push_plus"`
	TgBotToken string `json:"tg_bot_token"`
	TgUserId   string `json:"tg_user_id"`
}

var C *Config
*/

var VP *viper.Viper

// init
// @description: 初始化配置
func init() {

	VP = viper.New()

	VP.AutomaticEnv()

	VP.AddConfigPath(DefaultConfigPath)
	VP.SetConfigType(`json`)
	VP.SetConfigName(`config`)

	err := VP.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: #{err}\n"))
	}

	VP.WatchConfig()

	VP.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
	})

}
