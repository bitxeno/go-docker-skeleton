package app

import "github.com/bitxeno/go-docker-skeleton/internal/app/build"

var Version = AppVersion{
	Commit:    build.Commit,
	Version:   build.Version,
	BuildDate: build.BuildDate,
}

// Info holds build information
type AppVersion struct {
	Commit    string `json:"commit"`
	Version   string `json:"version"`
	BuildDate string `json:"build_date"`
}
