package envkeygo

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/envkey/envkey-fetch/fetch"
	"github.com/joho/godotenv"
)

func Load() error {
	godotenv.Load()
	envkey := os.Getenv("ENVKEY")

	if envkey == "" {
		return errors.New("Missing ENVKEY")
	}

	res := fetch.Fetch(envkey, fetch.FetchOptions{true, ""})

	if strings.HasPrefix(res, "error:") {
		return errors.New(strings.Split(res, "error:")[1])
	}

	var resMap map[string]string
	err := json.Unmarshal([]byte(res), &resMap)

	if err != nil {
		return errors.New("There was a problem parsing EnvKey's response")
	}

	for k, v := range resMap {
		if os.Getenv(k) == "" {
			os.Setenv(k, v)
		}
	}

	return nil
}
