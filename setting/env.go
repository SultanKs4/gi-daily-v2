package setting

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
)

type env struct {
	Cookie           string
	TelegramBotToken string
	TelegramOwnerId  string
}

var Environment = &env{}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return fmt.Sprint(filepath.Dir(d), "/.env")
}
