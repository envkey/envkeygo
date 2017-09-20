package autoload

import "github.com/envkey/envkeygo"

func init() {
	err := envkeygo.Load()
	if err != nil {
		panic(err)
	}
}
