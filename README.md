# envkeygo

Integrate [EnvKey](https://www.envkey.com) with your Go projects to keep api keys, credentials, and other configuration securely and automatically in sync for developers and servers.

This repo is mirrored in two locations:

- [A subdirectory of EnvKey's v2 monorepo](https://github.com/envkey/v2/tree/main/public/sdks/languages-and-frameworks/go/envkeygo).
- [envkeygo module repo](https://github.com/envkey/envkeygo)

# v1

For docs on the v1 version of this package, go to the [latest v1 tag of the module repo](https://github.com/envkey/envkeygo/tree/v1.2.5). Using v2 requires an EnvKey v2 organization (it won't work with ENVKEYs generated in a v1 org).

[Here's a guide on migrating from v1 to v2.](https://docs-v2.envkey.com/docs/migrating-from-v1)

## envkey-source

In EnvKey v2, using the [envkey-source](https://docs-v2.envkey.com/docs/envkey-source) executable from the command line offers additional functionality, like automatic reloads, that aren't available to EnvKey's language-specific SDKs. Consider using it instead of this library if it fits your use case.

## Installation

```bash
go get github.com/envkey/envkeygo/v2
```

## Usage

To load an EnvKey environment, follow the [integration quickstart](https://docs-v2.envkey.com/docs/integration-quickstart), but stop before integrating with envkey-source (which you won't be doing).

Now load your EnvKey configuration in `main.go`:

```go
// main.go
import (
  "os"
  _ "github.com/envkey/envkeygo/v2"
)

// assuming you have GITHUB_TOKEN set in EnvKey
token := os.Getenv("GITHUB_TOKEN") // this will stay in sync
```

### Overriding Vars

envkeygo will not overwrite existing environment variables or additional variables set in the `.env` file you loaded your `ENVKEY` from. This can be convenient for customizing environments that otherwise share the same configuration. You can also use [branches or local overrides](https://docs-v2.envkey.com/docs/branches-and-local-overrides) for this purpose.

### Working Offline

envkeygo can cache your encrypted config in development so that you can still use it while offline. Your config will still be available (though possibly not up-to-date) the next time you lose your internet connection. If you do have a connection available, envkeygo will always load the latest config. Your cached encrypted config is stored in `$HOME/.envkey/cache`

To turn on caching, set a `ENVKEY_SHOULD_CACHE=1` environment variable when running your program (_not_ in your EnvKey config):

```bash
ENVKEY_SHOULD_CACHE=1 ./your-program
```

## x509 error / ca-certificates

On a stripped down OS like Alpine Linux, you may get an `x509: certificate signed by unknown authority` error when envkeygo attempts to load your config. envkey-source (which envkeygo wraps) tries to handle this by including its own set of trusted CAs via [gocertifi](https://github.com/certifi/gocertifi), but if you're getting this error anyway, you can fix it by ensuring that the `ca-certificates` dependency is installed. On Alpine you'll want to run:

```
apk add --no-cache ca-certificates
```

## Further Reading

For more on EnvKey in general:

Read the [docs](https://docs-v2.envkey.com).

Read the [integration quickstart](https://docs-v2.envkey.com/docs/integration-quickstart.html).

Read the [security and cryptography overview](https://docs-v2.envkey.com/docs/security).

## Need help? Have questions, feedback, or ideas?

Post an [issue](https://github.com/envkey/envkey/issues), start a [discussion](https://github.com/envkey/envkey/dicussions), or email us: [support@envkey.com](mailto:support@envkey.com).
