package config

import (
	"os"
)

const (
	defaultCheckDepsYmlPath = "checkdeps.yml"
)

func GetCheckDepsYmlPath(envPath string) string {
	p := os.Getenv(envPath)
	if p == "" {
		p = defaultCheckDepsYmlPath
	}
	return p
}
