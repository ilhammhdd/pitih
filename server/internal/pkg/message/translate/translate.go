package translate

import (
	"ilhammhdd.com/pitih/internal/pkg/localization"
	messagePkg "ilhammhdd.com/pitih/internal/pkg/message"
)

func Translate(locale localization.Localization, message string, args ...any) string {
	switch message {
	case messagePkg.Hello:
		return locale.Translate(messagePkg.Hello, args...)
	default:
		return locale.Translate(message)
	}
}
