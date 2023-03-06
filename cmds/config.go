package cmds

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"tool/mods/config"
	"tool/mods/db"
	"tool/mods/util"
)

func DumpDefaultConfig() {
	content, err := config.GeneDefaultConfig()
	if err != nil {
		fmt.Println("failed to generate default config")
	} else {
		fmt.Println(string(content))
	}
}

func SyncDB() (err error) {
	if !db.DBEnable() {
		logrus.Info("no enable sql")
		return
	}
	err = db.Connect()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to migrate:%v", err))
	}
	err = db.Migrate()
	if err != nil {
		return
	}

	err = db.InitData()
	return
}

func AutoMigrate() (err error) {
	if !db.DBEnable() {
		return
	}
	err = db.Connect()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to migrate:%v", err))
	}
	return db.Migrate()
}

func InitDefaultServerConfigFile(configFile string) string {
	if configFile != "" {
		return configFile
	}

	execPath, _ := os.Getwd()
	configDir := filepath.Join(execPath, "config")
	logDir := filepath.Join(execPath, "log")
	file := filepath.Join(configDir, "agent.json")

	if !util.PathExists(configDir) {
		_ = os.MkdirAll(configDir, os.ModePerm)
	}

	if !util.PathExists(logDir) {
		_ = os.MkdirAll(logDir, os.ModePerm)
	}

	if util.PathExists(file) {
		return file
	}

	content, _ := config.GeneDefaultConfig()
	_ = os.WriteFile(file, content, os.ModePerm)
	return file
}
