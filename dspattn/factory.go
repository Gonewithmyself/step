package dspattn

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

type extConfig struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func parseConfig(fname string, data []byte) (*extConfig, error) {
	ext := filepath.Ext(fname)
	var (
		conf extConfig
		er   error
	)
	switch ext {
	case "json":
		er = json.Unmarshal(data, &conf)
	case "ini":
	case "yml", "yaml":

	default:
		er = fmt.Errorf("unknown ext %s", ext)
	}

	return &conf, er
}
