package path

import "os"

type RootPath string

func NewRootPath() RootPath {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	rp := RootPath(pwd + "/")
	return rp
}
