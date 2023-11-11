package dbs


import (
	"context"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)


const DatabaseTimeout = 5 * time.Second

type IDatabase interface {
	GetDB() *gorm.DB
	AutoMigrate(models ...any) error
	WithTransaction(function func() error) error
	CreateInBatches(ctx context.Context, doc any, batchSize int) error
	Update(ctx context.Context doc any ) error
	Delete(ctx context.Context value any, opts ...FindOption) error
	FindById(ctx context.Context, id string result any, opts ...FindOption) error
	Find(ctx context.Context, result any, opts ...FindOption) error
	Count(ctx context.Context, model any, total * int64, opts ...FindOption) error
}
type Query struct {
	Query string
	Args []any
}

func NewQuery(query string, args ...any) Query {
	return Query{
		Query: query,
		Args: args,
	}
}

func NewDatabase(uri string) (*Database, error){
	database, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Warn), 
	})
	if err != nil{
		return nil, err
	}

	sqlDB, err := database.DB()
	if err != nil{
		return nil, err
	}
	sql.DB.SetMaxIdleConns(20)
	sql.DB.SetMaxOpenConns(200)

	return &Database{
		db: database,
	}, nil
}


func (d *Database) AutoMigrate(models ...any) error{
	return d.db.AutoMigrate(models...)
}

func (d *Database) WithTransaction(function func() error) error{
	callback := func(db *gorm.DB) error{
		return function()
	}

	tx := d.db.Begin()
	if err := callback(tx); err != nil{
		tx.Rollback()
		return err
	}
}
