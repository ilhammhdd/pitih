package relationalDB

import "strings"

const InsertClause = "INSERT INTO"
const SelectClause = "SELECT"
const ValuesClause = "VALUES"
const WhereClause = "WHERE"
const SetClause = "SET"
const FromClause = "FROM"

func getClauseIndex(query, clause string) (int, bool) {
	clauseIdx := strings.Index(query, clause)
	if clauseIdx == -1 {
		clauseIdx = strings.Index(query, strings.ToLower(clause))
		if clauseIdx == -1 {
			return -1, false
		}
	}
	return clauseIdx, true
}
