package repository

import (
	"github.com/scostadavid/tiger/internal/app/monitor/entity"
)

// interface
type Reader interface {
	GetByID(id int) (*entity.Monitorable, error)
	GetAll() ([]*entity.Monitorable, error)
}

type Writer interface {
	Create(monitorable *entity.Monitorable) error
	Update(monitorable *entity.Monitorable) error
	DeleteById(id int) error
}

type ReaderWriter interface {
	Reader
	Writer
}

type MonitorableRepository struct { // implements readerwriter
	// db connection
}

// db connection arguments
func NewMonitorableRepository() *MonitorableRepository {
	return &MonitorableRepository{}
}

func (r *MonitorableRepository) GetByID(id int) (*entity.Monitorable, error) {
	return &entity.Monitorable{ID: id, Name: "Foo", URI: "foo.com/url"}, nil
}

func (r *MonitorableRepository) GetAll() ([]*entity.Monitorable, error) {
	return nil, nil
}

func (r *MonitorableRepository) Create(monitorable *entity.Monitorable) error {
	return nil
	// query
	// _, err := r.db.Exec(ctx, query, ...parameters)
	// return err
}

func (r *MonitorableRepository) Update(monitorable *entity.Monitorable) error {
	return nil
}

func (r *MonitorableRepository) DeleteById(id int) error {
	return nil
}
