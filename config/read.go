package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func ReadFile() (*Config, error) {
	conf := &Config{}

	err := conf.SetDefault()
	if err != nil {
		return nil, err
	}

	filename, _ := filepath.Abs(FileName)
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, create it with default config
			err = conf.WriteDefaultFile(filename)
			if err != nil {
				return nil, err
			}
			return conf, nil
		}
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(conf)
	if err != nil {
		return nil, err
	}

	// fmt.Println("config:", conf)

	return conf, nil
}
