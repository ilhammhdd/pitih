package relationalDB

import (
	"strings"
)

type KeyReplacement struct {
	Key         string
	Replacement string
}

func FormatQuery(queryFormat string, keyReplacements []*KeyReplacement) string {
	for idx := range keyReplacements {
		queryFormat = strings.ReplaceAll(queryFormat, keyReplacements[idx].Key, keyReplacements[idx].Replacement)
	}
	return queryFormat
}

func FormatQueryV2(queryFormat string, keyReplacements map[string]string) string {
	for key, value := range keyReplacements {
		queryFormat = strings.ReplaceAll(queryFormat, "{"+key+"}", value)
	}
	return queryFormat
}

func FormatQueryV3(queryFormat []string, keyReplacements map[string]string) string {
	queryFormatJoined := strings.Join(queryFormat, " ")
	for key, value := range keyReplacements {
		queryFormatJoined = strings.ReplaceAll(queryFormatJoined, "{"+key+"}", value)
	}
	return queryFormatJoined
}
