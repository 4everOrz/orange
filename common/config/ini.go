package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var Global *ini.File

func init() {
	if err := ReadFile(); err != nil {
		fmt.Println("read config file failed")
		os.Exit(1)
	}
}

//重载
func ReLoad() {
	if err := ReadFile(); err != nil {
		fmt.Println("read config file failed")
		os.Exit(1)
	}
}

//读取config file
func ReadFile() (err error) {
	Global, err = ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		//	os.Exit(1)
	}
	return err
}

//读取字符串
func GetString(section, key string) (string, error) {
	if Global != nil {
		return Global.Section(section).Key(key).String(), nil
	} else {
		return "", errors.New("cfg is nil")
	}
}

//读取字符串,异常时返回默认值
func GetStringDefault(section, key, defaultv string) (string, error) {
	if Global != nil {
		return Global.Section(section).Key(key).MustString(defaultv), nil
	} else {
		return "", errors.New("cfg is nil")
	}
}

//读取整型
func GetInt(section, key string) (int, error) {
	if Global != nil {
		value, err := Global.Section(section).Key(key).Int()
		return value, err
	} else {
		return 0, errors.New("cfg is nil")
	}
}

//读取整型,异常时返回默认值
func GetIntDefault(section, key string, defaultv int) (int, error) {
	if Global != nil {
		value := Global.Section(section).Key(key).MustInt(defaultv)
		return value, nil
	} else {
		return 0, errors.New("cfg is nil")
	}
}

//读取整型64
func GetInt64(section, key string) (int64, error) {
	if Global != nil {
		value, err := Global.Section(section).Key(key).Int64()
		return value, err
	} else {
		return 0, errors.New("cfg is nil")
	}
}

//读取整型64,异常时返回默认值
func GetInt64Default(section, key string, defaultv int64) (int64, error) {
	if Global != nil {
		value := Global.Section(section).Key(key).MustInt64(defaultv)
		return value, nil
	} else {
		return 0, errors.New("cfg is nil")
	}
}

//读取布尔
func GetBool(section, key string) (bool, error) {
	if Global != nil {
		value, err := Global.Section(section).Key(key).Bool()
		return value, err
	} else {
		return false, errors.New("cfg is nil")
	}
}

//读取布尔,异常时返回默认值
func GetBoolDefault(section, key string, defaultv bool) (bool, error) {
	if Global != nil {
		value := Global.Section(section).Key(key).MustBool(defaultv)
		return value, nil
	} else {
		return false, errors.New("cfg is nil")
	}
}

//修改字段并保存到当前文件
func UpdateSave(section, key, value string) error {
	if Global != nil {
		Global.Section(section).Key(key).SetValue(value)
		err := Global.SaveTo("config.ini")
		return err
	} else {
		return errors.New("cfg is nil")
	}
}

/*
	cfg, err := ini.Load("my.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	// 典型读取操作，默认分区可以使用空字符串表示
	cfg.Section("").Key("app_mode").String()
	cfg.Section("paths").Key("data").String()
	// 我们可以做一些候选值限制的操作
	cfg.Section("server").Key("protocol").In("http", []string{"http", "https"})
	// 如果读取的值不在候选列表内，则会回退使用提供的默认
	cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"})
	// 自动类型转换
	cfg.Section("server").Key("http_port").MustInt(9999)
	cfg.Section("server").Key("enforce_domain").MustBool(false)
	// 修改某个值然后进行保存
	cfg.Section("").Key("app_mode").SetValue("production")
	cfg.SaveTo("my.ini")
*/
