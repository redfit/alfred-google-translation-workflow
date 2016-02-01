package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/spf13/viper"
	"os"
)

func cmdSetup(c *cli.Context) {
	token := c.Args().First()
	loadConfig()
	viper.Set("accessToken", token)

	err := saveConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Token saved")
}
