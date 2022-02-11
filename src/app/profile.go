package app

import (
	"fmt"
	"os"
)

func CheckProfile(profile string) error {
	if err := CheckProfileDir(profile); err != nil {
		return err
	}
	if err := CheckProfileConfigDir(profile + "/config"); err != nil {
		return err
	}
	return nil
}

func CheckProfileDir(path string) error {
	f, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path, os.ModePerm)
			return nil
		} else {
			return nil
		}
	}
	fmt.Println(f.Name())
	return nil
}

func CheckProfileConfigDir(path string) error {
	f, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path, os.ModePerm)
			return nil
		} else {
			return nil
		}
	}
	fmt.Println(f.Name())
	return nil
}
