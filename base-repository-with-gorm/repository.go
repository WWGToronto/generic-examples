package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type RepoType interface {
	TableName() string
}

type BaseRepo[T RepoType, K any] struct {
	db *gorm.DB
}

func NewRepo[T RepoType, K any]() *BaseRepo[T, K] {
	connection, err := gorm.Open(
		sqlite.Open("file::memory:?cache=shared"),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}

	// Test code to make the example work w/o a database & in-memory sqlite.
	connection.Exec(
		`
			create temp table if not exists profiles (
				id int primary key,
				first_name varchar,
				middle_name varchar,
				last_name varchar
			);
		`,
	)

	return &BaseRepo[T, K]{
		db: connection,
	}
}

func (r BaseRepo[T, K]) CommitTx() error {
	result := r.db.Commit()
	return result.Error
}

func (r BaseRepo[T, K]) BeginTx() (BaseRepo[T, K], error) {
	return BaseRepo[T, K]{
		db: r.db.Begin(),
	}, nil
}

func (r BaseRepo[T, K]) RollbackTx() error {
	result := r.db.Rollback()
	return result.Error
}

func (r BaseRepo[T, K]) Create(data T) (*T, error) {
	result := r.db.Create(&data)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create: %w", result.Error)
	}
	return &data, nil
}

func (r BaseRepo[T, K]) GetByID(id K) (*T, error) {
	data := *new(T)
	result := r.db.
		Where("id = ?", fmt.Sprintf("%s", id)).
		First(&data)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get by ID: %w", result.Error)
	}
	return &data, nil
}

func (r BaseRepo[T, K]) Update(id K, data T) (*T, error) {
	result := r.db.
		Model(*new(T)).
		Where("id = ?", fmt.Sprintf("%s", id)).
		Updates(data)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to update: %w", result.Error)
	}
	return r.GetByID(id)
}
