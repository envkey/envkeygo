package envkeygo_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/envkey/envkeygo"
)

const VALID_ENVKEY = "Emzt4BE7C23QtsC7gb1z-3NvfNiG1Boy6XH2o-env-staging.envkey.com"
const INVALID_ENVKEY = "Emzt4BE7C23QtsC7gb1z-3NvfNiG1Boy6XH2oinvalid-env-staging.envkey.com"

func TestLoadMissing(t *testing.T) {
	os.Clearenv()
	assert.NotNil(t, envkeygo.Load())
}

func TestLoadValid(t *testing.T) {
	os.Clearenv()
	os.Setenv("ENVKEY", VALID_ENVKEY)
	err := envkeygo.Load()
	assert.Nil(t, err)
	assert.Equal(t, "it", os.Getenv("TEST"))
	assert.Equal(t, "works!", os.Getenv("TEST_2"))
}

func TestLoadInvalid(t *testing.T) {
	os.Clearenv()
	os.Setenv("ENVKEY", INVALID_ENVKEY)
	assert.NotNil(t, envkeygo.Load())
}

func TestLoadOverrides(t *testing.T) {
	os.Clearenv()
	os.Setenv("ENVKEY", VALID_ENVKEY)
	os.Setenv("TEST_2", "override")
	err := envkeygo.Load()
	assert.Nil(t, err)
	assert.Equal(t, "it", os.Getenv("TEST"))
	assert.Equal(t, "override", os.Getenv("TEST_2"))
}
