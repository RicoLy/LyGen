package commands

import (
	"LyGen/constant"
	"LyGen/db"
	"LyGen/logic"
	"LyGen/service"
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
	c.Help()
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

	return db.MysqlRepo.Close()
}

func (c *Commands) Handlers() map[string]func(args ...string) int {
	return map[string]func(args ...string) int{
		"0":     c.customDir,
		"1":     c.markDown,
		"2":     c.GenerateEntry,
		"3":     c.GenerateCURD,
		"fiber": c.GenerateFiberProject,
		"cl":    c.Clean,
		"q":     c.Quit,
		"h":     c.Help,
	}
}

// customDir 设置文件生成目录
func (c *Commands) customDir(_ ...string) int {
	fmt.Print("Please set the build directory>")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	if string(line) != "" {
		path, err := tools.GenerateDir(string(line))
		if err == nil {
			constant.CustomDir = path
			path = strings.TrimRight(path, constant.DS)
			constant.Project = tools.FindLastStr(path, constant.DS)
			fmt.Println("Directory success:", path)
		} else {
			log.Println("Set directory failed>>", err)
		}
	}
	return 0
}

// customProtoPath 设置proto文件路径
func (c *Commands) customProtoPath(_ ...string) int {
	fmt.Print("Please set the proto file path>")
	for  {
		line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		if string(line) != "" {
			constant.ProtoPath = string(line)
			break
		}
		fmt.Print("Proto file path is null please reset>")
	}

	return 0
}

// markDown 自定义生成目录
func (c *Commands) markDown(_ ...string) int {
	c.customDir()
	c.ConnectMysql()
	c._setDatabaseName()
	fmt.Println("Preparing to generate the markdown document...")
	err := logic.Lg.CreateMarkdown()
	if err != nil {
		log.Println("MarkDown>>", err)
	}
	return 0
}

// ConnectMysql 初始化数据库
func (c *Commands) ConnectMysql(_ ...string) int {

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

// GenerateEntry 生成结构体
func (c *Commands) GenerateEntry(args ...string) int {
	formats := c._setFormat()
	c.customDir()
	c.ConnectMysql()
	c._setDatabaseName()
	err := logic.Lg.CreateEntity(formats)
	if err != nil {
		log.Println("GenerateEntry>>", err.Error())
		return 0
	}
	go tools.Gofmt(constant.CustomDir) //格式化代码
	return 0
}

// GenerateCURD 生成CRUD代码
func (c *Commands) GenerateCURD(_ ...string) int {
	formats := c._setFormat()
	c.customDir()
	c.ConnectMysql()
	c._setDatabaseName()
	err := logic.Lg.CreateCURD(formats)
	if err != nil {
		log.Println("GenerateCURD>>", err.Error())
	}
	go tools.Gofmt(constant.CustomDir)
	return 0
}

// GenerateFiberProject 生成fiber项目
func (c *Commands) GenerateFiberProject(_ ...string) int {
	c.customDir()
	c.customProtoPath()
	fmt.Println("GenerateFiberProject|constant.Project", constant.Project)
	if err := logic.Lg.GenerateFiberProject(); err != nil {
		log.Println("GenerateFiberProject>>", err.Error())
	}
	go tools.Gofmt(constant.CustomDir)
	return 0
}

// _setFormat set struct format
func (c *Commands) _setFormat() []string {
	fmt.Print("Do you need to set the format of the structure?(Yes|No)>")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	switch strings.ToLower(string(line)) {
	case "yes", "y":
		fmt.Print("Set the mapping name of the structure, separated by a comma (example :json,gorm)>")
		input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		if string(input) != "" {
			formatList := tools.CheckCharDoSpecialArr(string(input), ',', `[\w\,\-]+`)
			if len(formatList) > 0 {
				fmt.Printf("Set format success: %v\n", formatList)
				return formatList
			}
		}
		fmt.Println("Set failed")
	}

	return nil
}

//help list
func (c *Commands) Help(args ...string) int {
	for _, row := range constant.CmdHelp {
		s := fmt.Sprintf("%s %s\n", "NO:"+row.No, row.Msg)
		fmt.Print(s)
	}
	return 0
}

func (c Commands) _setDatabaseName(args ...string) int {
	if service.DB.DBName == "" {
		for {
			fmt.Print("Please set database name>")
			line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
			if string(line) != "" {
				service.DB.DBName = string(line)
				break
			} else {
				log.Println("database name is empty>>")
			}
		}
	}
	return 0
}

// Quit 退出
func (c *Commands) Quit(_ ...string) int {
	return 1
}

// Clean 清屏
func (c *Commands) Clean(_ ...string) int {
	tools.Clean()
	return 0
}
