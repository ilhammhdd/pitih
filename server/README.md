# Introduction

this is a monorepo of server and website for Pitih

## Tools and library
1. error tracing with [tracerr](https://github.com/ztrue/tracerr) and wrapped in [errors](./server/internal/pkg/errors/)
1. dependency injection using [wire](https://github.com/google/wire)
1. translation using [go text](golang.org/x/text) following guidance from [here](https://www.alexedwards.net/blog/i18n-managing-translations)
1. configuration solution using [viper](https://github.com/spf13/viper)

## Resolve Dependency
1. go to `server` directory
1. run: `go mod download`
1. to vendor the dependencies, run: `go mod vendor`

## Dependency Injection
install wire: `go install github.com/google/wire/cmd/wire@latest`\
generate wire file:
1. go to `dependency` directory
1. run: `$(go env GOPATH)/bin/wire`

for instruction on how to use wire please refer to their [user guide](https://github.com/google/wire/blob/main/docs/guide.md) and [tutorial](https://github.com/google/wire/blob/main/_tutorial/README.md)

## Translation and Localisation
install gotext: `go install golang.org/x/text/cmd/gotext@latest`\
gotext use the message itself as the translation ID

### Adding new message
1. add it in `const` inside [message.go](server/internal/pkg/message/message.go)
1. add a case to translate it in [translate.go](server/internal/pkg/message/translate/translate.go)
> this is necessary because gotext will only generate `out.gotext.json` for a direct call to message.Printer.Printf(), Fprintf() and Sprintf(), and `Localization.Translate` did it for us
1. go to `server` directory
1. run: `./script/generate-translation.sh`
1. run: `./script/cp-generated-translation.sh`
1. for each language tag directory in `./server/internal/pkg/translation/locales`, do:
    1. open `messages.gotext.json`
    1. fill every empty value of `translation` key
    1. for handling plural translation fill `translation` like below:
    ```
    "translation": {
        "select": {
            "feature": "plural",
            "arg": "<ArgumentName>",
            "cases": {
                "=0": {
                    "msg": "message to display if `arg` is zero"
                },
                "=1": {
                    "msg": "message to display if `arg` is one"
                },
                "other": {
                    "msg": "message to display if `arg` is more than one"
                }
            }
        }
    }
    ```
1. run: `./script/generate-translation.sh` again

### Using message translation
1. add `import "ilhammhdd.com/pitih/internal/pkg/message/translate"`
1. call `translate.Translate` and pass localization and message, and arguments if any