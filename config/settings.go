package config

import (
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/bitxeno/go-docker-skeleton/internal/cfg"
	"github.com/bitxeno/go-docker-skeleton/internal/log"
	"github.com/bitxeno/go-docker-skeleton/internal/utils"
	"github.com/creasty/defaults"
)

var Settings SettingsConfiguration
var saveTimer *time.Timer = time.NewTimer(math.MaxInt64)

type SettingsConfiguration struct {
	Test string `koanf:"test" default:"test"`
}

func loadSettings() error {
	// load from settings.json file
	conf := cfg.New()
	conf.SetPath(filepath.Join(cfg.Server.WorkDir, "settings.json"))
	conf.Load()

	// set default value
	if err := defaults.Set(&Settings); err != nil {
		return err
	}
	if err := conf.BindStruct("", &Settings); err != nil {
		return err
	}

	saveTimer.Stop()
	go startSaveSettingsJob(conf.Path())
	return nil
}

func SaveSettings() {
	saveTimer.Reset(100 * time.Millisecond)
}

func startSaveSettingsJob(settingsPath string) {
	go func() {
		for {
			<-saveTimer.C
			log.Infof("Start to save settings... %s", settingsPath)

			if settingsPath == "" {
				log.Info("Setting path is empty.")
				continue
			}

			data := utils.ToIndentJSON(Settings)
			if err := ioutil.WriteFile(settingsPath, data, os.ModePerm); err != nil {
				log.Err(err).Msg("Save settings error.")
			} else {
				log.Infof("Save settings success. %s", settingsPath)
			}
		}
	}()
}
