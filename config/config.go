package config

import (
	"encoding/json"
)


type ConfigPlugin interface {
	ProcessConfig(configMap map[string]json.RawMessage) (string, error)
}
