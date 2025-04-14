package main

import (

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/v2/frame/g"

	"errors"
	
	"vocabulary/internal/cmd"
)

func main() {
	var err error
	g.I18n().SetLanguage("zh-CN")

	err = connectDB()
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}

func connectDB() error{
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("数据库连接失败")
	}
	return nil
}
