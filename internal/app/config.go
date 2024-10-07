package app

import "github.com/bitxeno/go-docker-skeleton/internal/db"

var (
	Config *Configuration
)

// configuration holds any kind of configuration that comes from the outside world and
// is necessary for running the application.
type Configuration struct {
	Log struct {
		Level      string `koanf:"level" default:"info"`
		TimeFormat string `koanf:"time_format" default:"2006-01-02 15:04:05.000"`
		LogFile    string `koanf:"log_file"`
		AccessLog  string `koanf:"access_log"`
	} `koanf:"log" json:"log"`

	Server struct {
		ListenAddr string `koanf:"listen_addr" default:"0.0.0.0"`
		Port       int    `koanf:"port" default:"9000"`
		DataDir    string `koanf:"data_dir"`
	} `koanf:"server" json:"server"`

	Db db.Config `koanf:"db" json:"db"`
}
