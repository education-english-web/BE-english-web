package cronjob

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Environments []string `yaml:"environments"`
	Schedules    []string `yaml:"schedules"`
}

type JobMapping map[string]Config

func LoadCronJobConfig(filePath string) (JobMapping, error) {
	var config JobMapping

	// #nosec
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return config, nil
}
