package recyclebin

// go-ini 的简单封装

import (
	"github.com/Wjinlei/golib"
	ini "gopkg.in/ini.v1"
)

type iniClass struct {
	File *ini.File
	path string
}

func NewIni() iniClass {
	return iniClass{}
}

// Create 创建一个配置文件
func (i iniClass) Create(filePath string) (iniClass, error) {
	// 创建ini文件,并跳过无法识别的数据行
	cfg := ini.Empty(ini.LoadOptions{
		SkipUnrecognizableLines: true,
	})

	filePathAbs, err1 := golib.GetAbsPath(filePath)
	if err1 != nil {
		return iniClass{}, err1
	}

	err2 := cfg.SaveTo(filePathAbs)
	if err2 != nil {
		return iniClass{}, err2
	}

	var configObject iniClass
	configObject.File = cfg
	configObject.path = filePathAbs

	return configObject, nil
}

// LoadFile 加载一个配置文件
func (i iniClass) LoadFile(filePath string) (iniClass, error) {
	// 加载ini文件,并跳过无法识别的数据行
	cfg, err := ini.LoadSources(ini.LoadOptions{
		SkipUnrecognizableLines: true,
	}, filePath)

	if err != nil {
		return iniClass{}, err
	}

	var configObject iniClass
	configObject.File = cfg
	configObject.path = filePath

	return configObject, nil
}

// Save 保存
func (i iniClass) Save() error {
	err := i.File.SaveTo(i.path)
	if err != nil {
		return err
	}
	return nil
}
