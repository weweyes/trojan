package trojan

import (
	"encoding/base64"
	"fmt"
	"trojan/core"
	"trojan/util"
)

var clientPath = "/root/config.json"

// GenClientJson 生成客户端json
func GenClientJson() {
	fmt.Println()
	var user core.User
	domain, err := core.GetValue("domain")
	if err != nil {
		fmt.Println(util.Yellow("无域名记录, 生成的配置文件需手填域名字段(ssl.sni)"))
		domain = ""
	}
	mysql := core.GetMysql()
	userList := mysql.GetData()
	if userList == nil {
		fmt.Println("连接mysql失败!")
		return
	}
	if len(userList) == 1 {
		user = *userList[0]
	} else {
		UserList()
		choice := util.LoopInput("请选择要生成配置文件的用户序号: ", userList, true)
		if choice < 0 {
			return
		}
		user = *userList[choice-1]
	}
	pass, err := base64.StdEncoding.DecodeString(user.Password)
	if err != nil {
		fmt.Println(util.Red("Base64解码失败: " + err.Error()))
		return
	}
	if !core.WriteClient(string(pass), domain, clientPath) {
		fmt.Println(util.Red("生成配置文件失败!"))
	} else {
		fmt.Println("成功生成配置文件: " + util.Green(clientPath))
	}
}
