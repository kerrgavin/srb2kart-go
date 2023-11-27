package config

import (
	"fmt"
	"strings"
)

type ConfigHelper struct {
	strings.Builder
}

func (ConfigHelper *ConfigHelper) Comment(comment string) {
	ConfigHelper.WriteString(fmt.Sprintf("#%s\n", comment))
}

func (ConfigHelper *ConfigHelper) ConfigString(serverConfig string, jsonConfig *string) error {
	if jsonConfig != nil {
		ConfigHelper.WriteString(fmt.Sprintf("%s \"%s\"\n", serverConfig, *jsonConfig))
	}
	return nil
}

func (ConfigHelper *ConfigHelper) ConfigBool(serverConfig string, jsonConfig *bool) error {
	if jsonConfig != nil {
		ConfigHelper.WriteString(fmt.Sprintf("%s \"%s\"\n", serverConfig, boolToConfig(*jsonConfig)))
	}
	return nil
}

func (ConfigHelper *ConfigHelper) ConfigInt(serverConfig string, jsonConfig *int) error {
	if jsonConfig != nil {
		ConfigHelper.WriteString(fmt.Sprintf("%s \"%d\"\n", serverConfig, *jsonConfig))
	}
	return nil
}

func boolToConfig(value bool) string {
	if value {
		return "On"
	}
	return "Off"
}
