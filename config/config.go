package config

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var configFile = "./app.conf"
var conf = make(map[string]string)

var (
	ErrNotExsit                = errors.New("config itme isn't exsit")
	ErrCanNotConvertInt        = errors.New("string can't convert int")
	ErrCanNotConvertIntStrings = errors.New("string can't convert []string")
)

//GetInt 获取配置key的值为int
func GetInt(key string) (int, error) {
	v, ok := conf[key]
	if !ok {
		return 0, ErrNotExsit
	}

	value, err := strconv.Atoi(v)
	if err != nil {
		return 0, ErrCanNotConvertInt
	}
	return value, nil
}

//GetString 获取配置key的值为string
func GetString(key string) (string, error) {
	v, ok := conf[key]
	if !ok {
		return "", ErrNotExsit
	}

	return v, nil
}

//GetStrings 获取配置key的值为[]string
func GetStrings(key string) ([]string, error) {
	v, ok := conf[key]
	if !ok {
		return nil, ErrNotExsit
	}

	value := strings.Split(v, ",")
	for i, s := range value {
		value[i] = strings.TrimSpace(s)
	}
	return value, nil
}

// GetBool 获取key的布尔值
func GetBool(key string) bool {
	v, ok := conf[key]
	if !ok {
		return false
	}
	return isTrue(v)
}

var trueStrings = []string{"t", "true", "1"}

func isTrue(str string) bool {
	str = strings.ToLower(str)
	for _, v := range trueStrings {
		if str == v {
			return true
		}
	}
	return false
}

func init() {
	confContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Println("config file[./app.conf] is not exist")
		return
	}

	br := bufio.NewReader(bytes.NewReader(confContent))
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			break
		}

		//#开头为注释忽略
		line = bytes.TrimSpace(line)
		if bytes.HasPrefix(line, []byte{'#'}) {
			continue
		}

		//空行忽略
		if len(line) == 0 {
			continue
		}

		//将参数用=分割
		keyValue := bytes.SplitN(line, []byte{'='}, 2)
		if len(keyValue) != 2 {
			log.Printf("%s is not a correct configuration item", string(line))
			continue
		}

		//去除首位的空格、双引号，单引号、tab
		key := strings.Trim(string(keyValue[0]), " \"'\t")
		value := strings.Trim(string(keyValue[1]), " \"'\t")

		conf[key] = value
	}
	fmt.Println(conf)
}
