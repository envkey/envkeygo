package envkeygo

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/envkey/envkey-fetch/fetch"
	"github.com/joho/godotenv"
)

func Load() {
	godotenv.Load()
	envkey := os.Getenv("ENVKEY")

	if envkey == "" {
		err := errors.New("Missing ENVKEY")
		panic(err)
	}

	res := fetch.Fetch(envkey, fetch.FetchOptions{true, ""})

	if strings.HasPrefix(res, "error:") {
		err := errors.New(strings.Split(res, "error:")[1])
		panic(err)
	}

	var resMap map[string]string
	err := json.Unmarshal([]byte(res), &resMap)

	if err != nil {
		panic(errors.New("There was a problem parsing EnvKey's response"))
	}

	for k, v := range resMap {
		if os.Getenv(k) == "" {
			os.Setenv(k, v)
		}
	}
}
