package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	music := viper.GetBool("music")

	if music {
		viper.SetDefault("music.path", "")
		musicPath := viper.GetString("music.path")
		if musicPath == "" {
			fmt.Println("Music path is not set")
			return
		}

	} else {

	}
}
