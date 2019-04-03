package main

import (
	"os"
	"strings"
)

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

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func GetParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func CreateGoPackageParentPath(path string) (string, error) {
	// 获取父目录
	parent := GetParentDirectory(path)
	// fmt.Println(parent)
	exist, err := PathExists(parent)
	if err != nil {
		return "", err
	}
	if exist {
		return "", nil
	} else {
		err := os.MkdirAll(parent, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	return parent, nil
}

func LinkGoPackagePath(oldName, newName string) error {
	err := os.Symlink(oldName, newName)
	if err != nil {
		return err
	}
	return nil
}

func CreateGoPathBin(path string) error {
	exist, err := PathExists(path)
	if err != nil {
		return err
	}
	if exist {
		return nil
	} else {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
