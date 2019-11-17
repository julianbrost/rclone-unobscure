package main

import (
	"fmt"
	"github.com/rclone/rclone/fs/config"
	"github.com/rclone/rclone/fs/config/obscure"
)

func main() {
	for _, section := range config.FileSections() {
		pass := config.FileGet(section, "pass")
		fmt.Println(section, obscure.MustReveal(pass))
	}
}
