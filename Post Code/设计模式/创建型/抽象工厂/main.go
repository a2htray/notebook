// 抽象工厂模式是一种创建型设计模式， 它能创建一系列相关的对象， 而无需指定其具体类。
// 抽象工厂模式建议为系列中的每件产品明确声明接口

package main

import (
	"fmt"
	"os"
)

type IConfigurationFactory interface {
	CreateConfiguration(file string) IConfiguration
}

type JSONConfigurationFactory struct {}

func (j *JSONConfigurationFactory) CreateConfiguration(file string) IConfiguration {
	return new(JSONConfiguration)
}

type YAMLConfigurationFactory struct {}

func (y *YAMLConfigurationFactory) CreateConfiguration(file string) IConfiguration {
	return new(YAMLConfiguration)
}

type IConfiguration interface {
	ParseInt64(key string) (int64, error)
	ParseString(key string) (string, error)
}

type JSONConfiguration struct {}

func (j *JSONConfiguration) ParseInt64(key string) (int64, error) {
	fmt.Println("使用 JSONConfiguration")
	// 实际代码
	return 0, nil
}

func (j *JSONConfiguration) ParseString(key string) (string, error) {
	fmt.Println("使用 JSONConfiguration")
	// 实际代码
	return "", nil
}


type YAMLConfiguration struct {}

func (y *YAMLConfiguration) ParseInt64(key string) (int64, error) {
	fmt.Println("使用 YAMLConfiguration")
	// 实际代码
	return 0, nil
}

func (y *YAMLConfiguration) ParseString(key string) (string, error) {
	fmt.Println("使用 YAMLConfiguration")
	// 实际代码
	return "", nil
}

func main() {
	var factory IConfigurationFactory
	var configuration IConfiguration

	switch os.Args[1] {
	case "json":
		factory = new(JSONConfigurationFactory)
	case "yaml":
		factory = new(YAMLConfigurationFactory)
	default:
		factory = new(JSONConfigurationFactory)
	}

	configuration = factory.CreateConfiguration(os.Args[1])

	_, _ = configuration.ParseInt64("test")
}

// 抄袭来源
// https://refactoringguru.cn/design-patterns/abstract-factory
