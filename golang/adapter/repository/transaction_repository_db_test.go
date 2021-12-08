package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xXHachimanXx/Imersao-fc-5/adapter/repository/fixture"
	"github.com/xXHachimanXx/Imersao-fc-5/domain/entity"
)

func TestTransactionRepositoryDbInsert(t *testing.T) {
	migrationDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationDir)
	defer fixture.Down(db, migrationDir)

	repository := NewTransactionRepositoryDb(db)
	err := repository.Insert("1", "1", 12.1, entity.APPROVED, "")
	assert.Nil(t, err)
}
