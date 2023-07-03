package config

func Load() error {
	if err := loadCustom(); err != nil {
		return err
	}
	return loadSettings()
}
