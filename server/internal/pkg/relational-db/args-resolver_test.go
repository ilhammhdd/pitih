package relationalDB_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	relationalDB "ilhammhdd.com/pitih/internal/pkg/relational-db"
)

var stubModelBudi = stubModel{
	id: 10, name: "budi", age: 23, foreigner: false, stocksOwned: 3.23,
}

func TestResolveInsertArgs(t *testing.T) {
	argsResolver := relationalDB.NewArgsResolver(stubMapScanDestFunc, stubAllColumns)

	actualArgs := argsResolver.ResolveInsertArgs(
		"INSERT INTO "+stubTableName+" VALUES "+relationalDB.NewPlaceHolder(len(stubAllColumns)),
		&stubModelBudi,
	)
	expectedArgs := []any{
		&stubModelBudi.id, &stubModelBudi.name, &stubModelBudi.age, &stubModelBudi.foreigner, &stubModelBudi.stocksOwned,
	}

	assert.Equal(t, expectedArgs, actualArgs)
}

func TestResolveInsertArgs_LowerCaseQuery(t *testing.T) {
	argsResolver := relationalDB.NewArgsResolver(stubMapScanDestFunc, stubAllColumns)

	actualArgs := argsResolver.ResolveInsertArgs(
		"insert into "+stubTableName+" values "+relationalDB.NewPlaceHolder(len(stubAllColumns)),
		&stubModelBudi,
	)
	expectedArgs := []any{
		&stubModelBudi.id, &stubModelBudi.name, &stubModelBudi.age, &stubModelBudi.foreigner, &stubModelBudi.stocksOwned,
	}

	assert.Equal(t, expectedArgs, actualArgs)
}

func TestResolveInsertArgs_MixedCase(t *testing.T) {
	argsResolver := relationalDB.NewArgsResolver(stubMapScanDestFunc, stubAllColumns)

	actualArgs := argsResolver.ResolveInsertArgs(
		"insert INTO "+stubTableName+" vaLUEs "+relationalDB.NewPlaceHolder(len(stubAllColumns)),
		&stubModelBudi,
	)
	var expectedArgs []any = nil

	assert.Equal(t, expectedArgs, actualArgs)
}

func TestResolveInsertArgs_ExplicitAllColumns(t *testing.T) {
	argsResolver := relationalDB.NewArgsResolver(stubMapScanDestFunc, stubAllColumns)

	actualArgs := argsResolver.ResolveInsertArgs(
		"INSERT INTO stub_model(*) VALUES "+relationalDB.NewPlaceHolder(len(stubAllColumns)),
		&stubModelBudi,
	)
	expectedArgs := []any{
		&stubModelBudi.id, &stubModelBudi.name, &stubModelBudi.age, &stubModelBudi.foreigner, &stubModelBudi.stocksOwned,
	}

	assert.Equal(t, expectedArgs, actualArgs)
}

func TestResolveInsertArgs_SomeColumns(t *testing.T) {
	argsResolver := relationalDB.NewArgsResolver(stubMapScanDestFunc, stubAllColumns)

	actualArgs := argsResolver.ResolveInsertArgs(
		"INSERT INTO stub_model(id,name,age) VALUES "+relationalDB.NewPlaceHolder(len(stubAllColumns)),
		&stubModelBudi,
	)
	expectedArgs := []any{&stubModelBudi.id, &stubModelBudi.name, &stubModelBudi.age}

	assert.Equal(t, expectedArgs, actualArgs)
}

func TestResolveInsertArgs_OneColumn(t *testing.T) {
	argsResolver := relationalDB.NewArgsResolver(stubMapScanDestFunc, stubAllColumns)

	actualArgs := argsResolver.ResolveInsertArgs(
		"INSERT INTO stub_model(name) VALUES "+relationalDB.NewPlaceHolder(len(stubAllColumns)),
		&stubModelBudi,
	)
	expectedArgs := []any{&stubModelBudi.name}

	assert.Equal(t, expectedArgs, actualArgs)
}

func TestResolveInsertArgs_ColumnNotExists(t *testing.T) {
	argsResolver := relationalDB.NewArgsResolver(stubMapScanDestFunc, stubAllColumns)

	actualArgs := argsResolver.ResolveInsertArgs(
		"INSERT INTO stub_model(phone_no) VALUES "+relationalDB.NewPlaceHolder(len(stubAllColumns)),
		&stubModelBudi,
	)
	expectedArgs := []any{}

	assert.Equal(t, expectedArgs, actualArgs)
}

func TestResolveInsertArgs_ColumnNotExistsMixedExistingOnes(t *testing.T) {
	argsResolver := relationalDB.NewArgsResolver(stubMapScanDestFunc, stubAllColumns)

	actualArgs := argsResolver.ResolveInsertArgs(
		"INSERT INTO stub_model(id,name,phone_no,age) VALUES "+relationalDB.NewPlaceHolder(len(stubAllColumns)),
		&stubModelBudi,
	)
	expectedArgs := []any{&stubModelBudi.id, &stubModelBudi.name, &stubModelBudi.age}

	assert.Equal(t, expectedArgs, actualArgs)
}

func TestResolveInsertArgs_WrongQuery(t *testing.T) {
	argsResolver := relationalDB.NewArgsResolver(stubMapScanDestFunc, stubAllColumns)

	actualArgs := argsResolver.ResolveInsertArgs("SELECT s.* FROM "+stubTableName, &stubModelBudi)

	assert.Empty(t, actualArgs)
}

func TestResolveUpdateArgs(t *testing.T) {
	query := "UPDATE " + stubTableName + " SET name=?, age=?, foreigner=?, stocksOwned=? WHERE id=?"

	model := stubModel{id: 7, name: "annie", age: 25, foreigner: true, stocksOwned: 1.20}
	expectedArgs := []any{&model.name, &model.age, &model.foreigner, &model.stocksOwned, &model.id}

	argsResolver := relationalDB.NewArgsResolver(stubMapScanDestFunc, stubAllColumns)
	actualArgs := argsResolver.ResolveUpdateArgs(query, &model, &model.id)

	assert.Equal(t, expectedArgs, actualArgs)
}

func TestResolveUpdateArgs_WithoutWhereClause(t *testing.T) {
	query := "UPDATE " + stubTableName + " SET name=?, age=?, foreigner=?, stocksOwned=?"

	model := stubModel{id: 7, name: "annie", age: 25, foreigner: true, stocksOwned: 1.20}

	argsResolver := relationalDB.NewArgsResolver(stubMapScanDestFunc, stubAllColumns)
	actualArgs := argsResolver.ResolveUpdateArgs(query, &model)

	assert.Nil(t, actualArgs)
}

func TestResolveUpdateArgs_SingleColumnUpdate(t *testing.T) {
	query := "UPDATE " + stubTableName + " SET name=? WHERE id=?"

	model := stubModel{id: 7, name: "annie", age: 25, foreigner: true, stocksOwned: 1.20}
	expectedArgs := []any{&model.name, &model.id}

	argsResolver := relationalDB.NewArgsResolver(stubMapScanDestFunc, stubAllColumns)
	actualArgs := argsResolver.ResolveUpdateArgs(query, &model, &model.id)

	assert.Equal(t, expectedArgs, actualArgs)
}

func TestResolveUpdateArgs_WrongQuery(t *testing.T) {
	query := "SELECT s.* FROM " + stubTableName + " s"

	model := stubModel{id: 7, name: "annie", age: 25, foreigner: true, stocksOwned: 1.20}

	argsResolver := relationalDB.NewArgsResolver(stubMapScanDestFunc, stubAllColumns)
	actualArgs := argsResolver.ResolveUpdateArgs(query, &model, &model.id)

	assert.Empty(t, actualArgs)
}
