package relationalDB_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	relationalDB "ilhammhdd.com/pitih/internal/pkg/relational-db"
)

func TestNewSelectModeler(t *testing.T) {
	stubModeler := relationalDB.NewSelectModeler("SELECT s.* FROM "+stubTableName+" s", "s", stubMapScanDestFunc)
	expectedScanDests := []any{
		&stubModeler.Model.id, &stubModeler.Model.name, &stubModeler.Model.age, &stubModeler.Model.foreigner, &stubModeler.Model.stocksOwned,
	}

	assert.NotNil(t, stubModeler.Model)
	assert.Equal(t, expectedScanDests, stubModeler.ScanDests)
}

func TestNewSelectModeler_NoAlias(t *testing.T) {
	stubModeler := relationalDB.NewSelectModeler("SELECT * FROM "+stubTableName, "", stubMapScanDestFunc)

	assert.Empty(t, stubModeler.Model)
	assert.Empty(t, stubModeler.ScanDests)
}

func TestNewSelectModeler_LowerCase(t *testing.T) {
	stubModeler := relationalDB.NewSelectModeler("select s.* from "+stubTableName+" s", "s", stubMapScanDestFunc)
	expectedScanDests := []any{
		&stubModeler.Model.id, &stubModeler.Model.name, &stubModeler.Model.age, &stubModeler.Model.foreigner, &stubModeler.Model.stocksOwned,
	}

	assert.NotNil(t, stubModeler.Model)
	assert.Equal(t, expectedScanDests, stubModeler.ScanDests)
}

func TestNewSelectModeler_Join(t *testing.T) {
	query := "SELECT s1.*, s2.name, s2.age FROM table s1 JOIN table s2 ON s2.id = s1.id"
	stubModelerOne := relationalDB.NewSelectModeler(query, "s1", stubMapScanDestFunc)
	stubModelerTwo := relationalDB.NewSelectModeler(query, "s2", stubMapScanDestFunc)

	actualArgs := stubModelerOne.ScanDests
	actualArgs = append(actualArgs, stubModelerTwo.ScanDests...)

	expectedScanDests := []any{
		&stubModelerOne.Model.id,
		&stubModelerOne.Model.name,
		&stubModelerOne.Model.age,
		&stubModelerOne.Model.foreigner,
		&stubModelerOne.Model.stocksOwned,
		&stubModelerTwo.Model.name,
		&stubModelerTwo.Model.age,
	}

	assert.NotNil(t, stubModelerOne.Model)
	assert.NotNil(t, stubModelerTwo.Model)
	assert.Equal(t, expectedScanDests, actualArgs)
}

func TestNewSelectModeler_WrongQuery(t *testing.T) {
	stubModeler := relationalDB.NewSelectModeler("UPDATE SET s.name=? FROM stub_table s WHERE s.id=?", "s", stubMapScanDestFunc)

	assert.Nil(t, stubModeler)
}
