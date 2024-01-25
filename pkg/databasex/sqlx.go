package databasex

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Configs map[string]*Config

type Config struct {
	DSN     string        `mapstructure:"dsn" json:"dsn" yaml:"dsn"`
	ShowSql bool          `mapstructure:"show_sql" json:"show_sql" yaml:"show_sql"`
	Timeout time.Duration `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
}

func NewDBs(configs Configs) (map[string]*sql.DB, error) {
	dbs := make(map[string]*sql.DB)
	for key, config := range configs {
		db, err := NewDB(config)
		if err != nil {
			return nil, err
		}
		dbs[key] = db
	}
	return dbs, nil
}

func NewDB(config *Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DSN)
	if err != nil {
		return nil, err
	}
	if config.Timeout <= 0 {
		config.Timeout = 3 * time.Second
	}
	ctx, _ := context.WithTimeout(context.Background(), config.Timeout)
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}

func NewGormDBs(configs Configs) (map[string]*gorm.DB, error) {
	dbs := make(map[string]*gorm.DB)
	for key, config := range configs {
		db, err := NewGormDB(config)
		if err != nil {
			return nil, err
		}
		dbs[key] = db
	}
	return dbs, nil
}

func NewGormDB(config *Config) (*gorm.DB, error) {
	var opts []gorm.Option
	if config.ShowSql {
		opts = append(opts, &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	}
	db, err := gorm.Open(mysql.Open(config.DSN), opts...)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}
	return db, err
}

func NewSqlxDBs(configs Configs) (map[string]*sqlx.DB, error) {
	dbs := make(map[string]*sqlx.DB)
	for key, config := range configs {
		db, err := NewSqlxDB(config)
		if err != nil {
			return nil, err
		}
		dbs[key] = db
	}
	return dbs, nil
}

func NewSqlxDB(config *Config) (*sqlx.DB, error) {
	db, err := NewDB(config)
	if err != nil {
		return nil, err
	}
	return sqlx.NewDb(db, "mysql"), nil
}
