package localization

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	_ "ilhammhdd.com/pitih/internal/pkg/translation"
)

const (
	enUSLangTag = "en-US"
	idIDLangTag = "id-ID"
)

var DefaultLangTag = enUSLangTag

var localesMap = map[string]Localization{
	DefaultLangTag: {printer: message.NewPrinter(language.MustParse(DefaultLangTag))},
	idIDLangTag:    {printer: message.NewPrinter(language.MustParse(idIDLangTag))},
}

type Localization struct {
	printer *message.Printer
}

func (l Localization) Translate(key message.Reference, args ...interface{}) string {
	return l.printer.Sprintf(key, args...)
}

func IsLangTagExists(langTag string) bool {
	_, ok := localesMap[langTag]
	return ok
}

// pass empty string to use default Localization
func MustGet(langTag string) Localization {
	if langTag == "" {
		return localesMap[DefaultLangTag]
	}

	locale, ok := localesMap[langTag]
	if ok {
		return locale
	}

	return localesMap[DefaultLangTag]
}
