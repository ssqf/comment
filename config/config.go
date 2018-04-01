package config

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var configFile = "./app.conf"
var conf = make(map[string]string)

var (
	ErrNotExsit                = errors.New("coifg itme isn't exsit")
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

		line = bytes.TrimSpace(line)
		if bytes.HasPrefix(line, []byte{'#'}) {
			continue
		}

		if len(line) == 0 {
			continue
		}

		keyValue := bytes.SplitN(line, []byte{'='}, 2)
		if len(keyValue) != 2 {
			log.Printf("%s is not a correct configuration item", string(line))
			continue
		}

		key := strings.Trim(string(keyValue[0]), " \"'\t")
		value := strings.Trim(string(keyValue[1]), " \"'\t")

		conf[key] = value
	}
}
