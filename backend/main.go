package main

import (
	"resume_analysis/api"

	"github.com/sirupsen/logrus"
)

func main() {
	err := api.Engine.Run()
	if err != nil {
		logrus.Fatalf("[main] engine run err:%+v", err)
	}
}
