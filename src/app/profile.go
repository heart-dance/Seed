package app

import (
	"os"
)

func CheckProfile(profile string) error {
	if err := CheckProfileDir(profile); err != nil {
		return err
	}
	if err := CheckProfileLogDir(profile + "/log"); err != nil {
		return err
	}
	return nil
}

func CheckProfileDir(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path, os.ModePerm)
			return nil
		} else {
			return nil
		}
	}
	return nil
}

func CheckProfileLogDir(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path, os.ModePerm)
			return nil
		} else {
			return nil
		}
	}
	return nil
}
