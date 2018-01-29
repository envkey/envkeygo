package loader

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/envkey/envkey-fetch/fetch"
	"github.com/joho/godotenv"
)

func Load(shouldCache bool) {
	godotenv.Load()
	envkey := os.Getenv("ENVKEY")

	if envkey == "" {
		panic(errors.New("missing ENVKEY"))
	}

	res := fetch.Fetch(envkey, fetch.FetchOptions{shouldCache, "", "envkeygo", ""})

	if strings.HasPrefix(res, "error:") {
		panic(errors.New(strings.Split(res, "error:")[1]))
	}

	var resMap map[string]string
	err := json.Unmarshal([]byte(res), &resMap)

	if err != nil {
		panic(errors.New("problem parsing EnvKey's response"))
	}

	for k, v := range resMap {
		if os.Getenv(k) == "" {
			os.Setenv(k, v)
		}
	}
}
