package common

import (
	"os"
	"path/filepath"
)

/* __________________________________________________ */

const (
	ConfigDirectory = ".gophkeeper"
	ConfigFilename  = "config"

	ConfigLocationPropertyName = "Config"
	ConfigLocationEnvName      = "GOPHKEEPER__CONFIG"
	ConfigLocationArgName      = "cfg"
)

//goland:noinspection GoNameStartsWithPackageName
func BuildConfigPath() (string, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(home, ConfigDirectory, ConfigFilename)
	return path, nil

}

//goland:noinspection GoNameStartsWithPackageName
func MustBuildConfigPath() string {
	path, err := BuildConfigPath()
	if err != nil {
		panic(err)
	}
	return path
}

/* __________________________________________________ */

const (
	AddressPropertyName = "Address"
	AddressEnvName      = "GOPHKEEPER__ADDRESS"
)

/* __________________________________________________ */
