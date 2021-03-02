package seeders

import (
	"gorm.io/gorm"
)

type Table struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewTable(db *gorm.DB) *Table {
	return &Table{db: db}
}

func (s *Table) Count(tableName string) int64 {
	var count int64
	s.db.Table(tableName).Count(&count)

	return count
}

func (s *Table) Prepare() error {
	s.tx = s.db.Begin()
	return s.tx.Error
}

func (s *Table) Commit() error {
	return s.tx.Commit().Error
}

func (s *Table) Rollback() error {
	return s.tx.Rollback().Error
}

func (s *Table) Create(models ...interface{}) error {
	for _, model := range models {
		setIdIfNotExists(model)

		if err := s.tx.Create(model).Error; err != nil {
			return err
		}
	}
	return nil
}

func selectStringsFromDB(query *gorm.DB) ([]string, error) {
	result := []string{}

	rows, err := query.Rows()
	if err != nil {
		return result, err
	}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return result, err
		}
		result = append(result, id)
	}
	if err := rows.Err(); err != nil {
		return result, err
	}
	return result, nil
}
