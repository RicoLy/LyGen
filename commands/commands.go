package commands

import (
	"LyGen/constant"
	"LyGen/db"
	"LyGen/logic"
	"LyGen/tools"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Commands struct {
}

func (c Commands) Start() error {

	br := bufio.NewReader(os.Stdin)
	handlers := c.Handlers()
	for {
		fmt.Print("-> # ")
		line, _, _ := br.ReadLine()
		if len(line) == 0 {
			continue
		}
		tokens := strings.Split(string(line), " ")
		if handler, ok := handlers[strings.ToLower(tokens[0])]; ok {
			ret := handler(tokens...)
			if ret != 0 {
				break
			}
		} else {
			fmt.Println("Unknown command>>", tokens[0])
		}
	}
	return nil
}

func (c *Commands) Handlers() map[string]func(args ...string) int {
	return map[string]func(args ...string) int{
		"0": c.customDir,
		"1": c.markDown,
	}
}

func (c *Commands) customDir(_ ...string) int {
	fmt.Print("Please set the build directory>")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	if string(line) != "" {
		path, err := tools.GenerateDir(string(line))
		if err == nil {
			constant.CustomDir = path
			fmt.Println("Directory success:", path)
		} else {
			log.Println("Set directory failed>>", err)
		}
	}
	return 0
}

// markDown 自定义生成目录
func (c *Commands) markDown(args ...string) int {
	fmt.Println("Preparing to generate the markdown document...")
	c.customDir()
	err := logic.Lg.CreateMarkdown()
	if err != nil {
		log.Println("MarkDown>>", err)
	}
	return 0
}

// ConnectMysql 初始化数据库
func (c Commands) ConnectMysql(_ ...string) int {

	if db.MysqlRepo.DB == nil {
		for {
			if constant.MysqlDataSource == "" {
				fmt.Print("Please set mysql dataSource>")
				line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
				constant.MysqlDataSource = string(line)
			}

			err := db.MysqlRepo.InitMysqlDB(constant.MysqlDataSource)
			if err != nil || db.MysqlRepo.DB == nil {
				fmt.Printf("connect mysql err: %v", err)
				constant.MysqlDataSource = ""
				continue
			}
			break
		}
	}

	return 0
}
