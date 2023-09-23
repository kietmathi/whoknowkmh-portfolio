package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database interface {
	Create(interface{}) error
	Find(interface{}, ...interface{}) error
	Updates(updates interface{}) error
	First(interface{}, ...interface{}) error
	Save(interface{}) error
	Delete(interface{}, ...interface{}) error
	Pluck(column string, value interface{}) error
	AutoMigrate(dst ...interface{}) error
	Where(condition interface{}, args ...interface{}) Database
	Order(value interface{}, reorder ...bool) Database
	Model(value interface{}) Database
	Select(query interface{}, args ...interface{}) Database
}

type Client interface {
	Database() Database
	Close() error
}

type sqliteClient struct {
	db *gorm.DB
}

func NewClient(dbPath string) (Client, error) {
	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &sqliteClient{db: database}, nil
}

func (sc *sqliteClient) Database() Database {
	return &sqliteDatabase{db: sc.db}
}

func (sc *sqliteClient) Close() error {
	sqlDB, err := sc.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

type sqliteDatabase struct {
	db *gorm.DB
}

func (sd *sqliteDatabase) Create(value interface{}) error {
	return sd.db.Create(value).Error
}

func (sd *sqliteDatabase) Updates(updates interface{}) error {
	return sd.db.Updates(updates).Error
}

func (sd *sqliteDatabase) Find(out interface{}, where ...interface{}) error {
	return sd.db.Find(out, where...).Error
}

func (sd *sqliteDatabase) First(out interface{}, where ...interface{}) error {
	return sd.db.First(out, where...).Error
}

func (sd *sqliteDatabase) Save(value interface{}) error {
	return sd.db.Save(value).Error
}

func (sd *sqliteDatabase) Delete(value interface{}, where ...interface{}) error {
	return sd.db.Delete(value, where...).Error
}

func (sd *sqliteDatabase) Pluck(column string, value interface{}) error {
	return sd.db.Model(value).Pluck(column, value).Error
}

func (sd *sqliteDatabase) AutoMigrate(dst ...interface{}) error {
	return sd.db.AutoMigrate(dst...)
}

func (sd *sqliteDatabase) Where(condition interface{}, args ...interface{}) Database {
	return &sqliteDatabase{db: sd.db.Where(condition, args...)}
}

func (sd *sqliteDatabase) Order(value interface{}, reorder ...bool) Database {
	return &sqliteDatabase{db: sd.db.Order(value)}
}

func (sd *sqliteDatabase) Model(value interface{}) Database {
	return &sqliteDatabase{db: sd.db.Model(value)}
}

func (sd *sqliteDatabase) Select(query interface{}, args ...interface{}) Database {
	return &sqliteDatabase{db: sd.db.Select(query, args...)}
}
