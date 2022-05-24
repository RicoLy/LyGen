package service

import (
	"LyGen/constant"
	"LyGen/tools"
	"LyGen/types"
	"fmt"
	"regexp"
	"strings"
)

type ParseService struct {
}

var ParseSrv = new(ParseService)

func (p *ParseService)ParseMessages(fileInfo string) []*types.Message {
	messages := make([]*types.Message, 0)
	messageReg := regexp.MustCompile(constant.RegExpMessage)
	results := messageReg.FindAllStringSubmatch(fileInfo, -1)
	for _, result := range results {
		if len(result) == 4 {
			message := new(types.Message)
			message.Meta = result[0]
			message.Comment = strings.Trim(result[1], " \n")
			message.Name = result[2]
			message.ElementInfos = p.ParseElementInfos(result[3])
			messages = append(messages, message)
		}
	}
	return messages
}

func (p *ParseService)ParseElementInfos(elementStr string) []*types.ElementInfo {
	elementInfos := make([]*types.ElementInfo, 0)
	elementReg := regexp.MustCompile(constant.RegExpElementInfo)
	results := elementReg.FindAllStringSubmatch(elementStr, -1)
	//typeNameReg := regexp.MustCompile("")
	for _, result := range results {
		element := new(types.ElementInfo)
		element.Meta = result[0]
		element.Tags = result[1]
		typeNames := strings.Split(strings.Trim(result[2], " "), " ")
		if len(typeNames) == 3 && typeNames[0] == "repeated" {
			element.Type = fmt.Sprintf("[]*%s", typeNames[1])
			element.Name = typeNames[2]
		} else {
			element.Type = typeNames[0]
			element.Name = typeNames[1]
		}
		if golangType := constant.ProtoTypeToGoType[element.Type]; golangType != "" {
			element.Type = constant.ProtoTypeToGoType[element.Type]
		}

		elementInfos = append(elementInfos, element)
	}

	return elementInfos
}

func (p *ParseService)ParseServices(fileInfo string) []*types.Service {
	services := make([]*types.Service, 0)
	serviceReg := regexp.MustCompile(constant.RegExpService)
	groupAndPrefixReg := regexp.MustCompile(constant.RegExpAnoInfo)
	results := serviceReg.FindAllStringSubmatch(fileInfo, -1)
	for _, result := range results {
		if len(result) == 4 {
			service := new(types.Service)
			service.Meta = result[0]
			service.Comment = strings.Trim(tools.FindTopStr(result[1], " @"), "/ ")
			service.Name = result[2]
			gpResults := groupAndPrefixReg.FindAllStringSubmatch(result[1], -1)
			for _, gpResult := range gpResults {
				if len(gpResult) == 3 {
					if gpResult[1] == "group" {
						service.Group = gpResult[2]
					}
					if gpResult[1] == "prefix" {
						service.Prefix = gpResult[2]
					}
				}
			}
			methods := p.ParseMethods(result[3])
			service.Methods = make([]*types.Method, 0)
			for _, method := range methods {
				method.Group = service.Group
				service.Methods = append(service.Methods, method)
			}
			services = append(services, service)
		}
	}

	return services
}

func (p *ParseService)ParseMethods(methodStr string) []*types.Method {
	methods := make([]*types.Method, 0)
	methodReg := regexp.MustCompile(constant.RegExpMethod)
	extraInfoReg := regexp.MustCompile(constant.RegExpAnoInfo)
	results := methodReg.FindAllStringSubmatch(methodStr, -1)
	for _, result := range results {
		if len(result) == 5 {
			method := new(types.Method)
			method.Mata = result[0]
			method.Comment = strings.Trim(tools.FindTopStr(result[1], " @"), "/ ")
			extraResults := extraInfoReg.FindAllStringSubmatch(result[1], -1)
			for _, extraResult := range extraResults {
				if extraResult[1] == "method" {
					method.MethodType = extraResult[2]
				}
				if extraResult[1] == "path" {
					if method.MethodType == "get" {
						pathInfo := strings.Split(extraResult[2], "/:")
						method.PathParams = make([]string, 0)
						for i, s := range pathInfo {

							if i == 0 {
								method.Path = strings.TrimRight(s, "/")
							} else {
								method.PathParams = append(method.PathParams, s)
							}
						}
					} else {
						method.Path = extraResult[2]
					}
				}
				if extraResult[1] == "middleware" {
					method.MiddleWares = strings.Split(extraResult[2], "|")
				}
			}
			method.Name = result[2]
			method.Request = result[3]
			method.Response = result[4]
			methods = append(methods, method)
		}
	}
	//fmt.Println("================================")
	//fmt.Println("methodStr:", methodStr)
	return methods
}