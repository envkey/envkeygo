# envkey-go

Integrate [EnvKey](https://www.envkey.com) with your Go projects to keep configuration securely and automatically in sync for developers and servers.

## Installation

```bash
go get github.com/envkey/envkeygo
```

## Usage

First, generate an `ENVKEY` in the [EnvKey App](https://github.com/envkey/envkey-app). Then set `ENVKEY=...`, either in a gitignored `.env` file in the root of your project (in development) or in an environment variable (on servers).

Then load your EnvKey configuration in your `main.go` file with `envkeygo.Load`. It accepts a single argument, `shouldCache`, which determines whether your encrypted config is cached locally (in $HOME/.envkey/cache) so that it's available for offline work if you lose your internet connection. In general, you should cache in development and not on servers.

```go
// main.go

import (
  "os"
  "github.com/envkey/envkeygo"
)

shouldCache := os.Getenv("IS_LOCAL") == "true" // determine this however you want
envkeygo.Load(shouldCache) // panics if ENVKEY is missing or invalid

// assuming you have GITHUB_TOKEN set in EnvKey
token := os.Getenv("GITHUB_TOKEN") // this will stay in sync
```

### Overriding Vars

envkeygo will not overwrite existing environment variables or additional variables set in a `.env` file. This can be convenient for customizing environments that otherwise share the same configuration. You can read more about this topic in the EnvKey [docs](https://docs.envkey.com/overriding-envkey-variables.html).

## Further Reading

For more on EnvKey in general:

Read the [docs](https://docs.envkey.com).

Read the [integration quickstart](https://docs.envkey.com/integration-quickstart.html).

Read the [security and cryptography overview](https://security.envkey.com).

## Need help? Have questions, feedback, or ideas?

Post an [issue](https://github.com/envkey/envkeygo/issues) or email us: [support@envkey.com](mailto:support@envkey.com).


