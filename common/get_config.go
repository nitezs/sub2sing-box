package common

import (
	"log/slog"
	"net/http"
	"os"
	"strings"
)

func GetConfig(config string) ([]byte, error) {
	var configData []byte
	var err error
	if strings.HasPrefix(config, "http://") || strings.HasPrefix(config, "https://") {
		res, err := http.Get(config)
		if err != nil {
			slog.Error("get config failed", err)
			return nil, err
		}
		defer res.Body.Close()
		_, err = res.Body.Read(configData)
		if err != nil {
			slog.Error("read config failed", err)
			return nil, err
		}
	} else {
		configData, err = os.ReadFile(config)
		if err != nil {
			slog.Error("read config failed", err)
			return nil, err
		}
	}
	return configData, nil
}
