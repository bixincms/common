package common

import (
	"errors"
	"os"
	"path/filepath"
)

//如果文件夹不存在则直接创建
func PathIsNoDirCreate(path string) error {
	dir_path := filepath.Dir(path)
	info, e := os.Stat(dir_path)
	if e == nil {
		//判断是否是文件
		if !info.IsDir() {
			return errors.New("目标路径存在,但不是文件夹")
		}
		return nil
	}
	//创建文件夹
	e = os.MkdirAll(dir_path, 0666)
	return e
}

//检查路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
