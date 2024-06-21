package app

import (
	"io/ioutil"
	"math"
	"os"
	"time"

	"github.com/bitxeno/go-docker-skeleton/internal/log"
	"github.com/bitxeno/go-docker-skeleton/internal/utils"
)

var (
	Settings *SettingsConfiguration
)
var saveTimer *time.Timer = time.NewTimer(math.MaxInt64)

// Settings holds any kind of configuration that can modify by admin web ui.
type SettingsConfiguration struct {
	Test string `koanf:"test" default:"test"`
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
