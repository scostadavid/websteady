package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/scostadavid/websteady/internal/app/monitorable/dto"
	"github.com/scostadavid/websteady/internal/app/monitorable/entity"
)

type Reader interface {
	GetByID(id int) (*entity.Monitorable, error)
	GetAll() ([]*entity.Monitorable, error)
}

type Writer interface {
	Create(dto *dto.AddMonitorable) (*entity.Monitorable, error)
	Update(monitorable *entity.Monitorable) error
	DeleteById(id int) error
}

type ReaderWriter interface {
	Reader
	Writer
}

type MonitorableRepository struct {
	db *sql.DB
}

func NewMonitorableRepository(db *sql.DB) (*MonitorableRepository, error) {
	return &MonitorableRepository{db: db}, nil
}

func (repo *MonitorableRepository) Close() {
	repo.db.Close()
}

func (r *MonitorableRepository) GetByID(id int) (*entity.Monitorable, error) {
	return &entity.Monitorable{ID: id, Name: "Foo", URL: "foo.com/url"}, nil
}

func (r *MonitorableRepository) GetAll() ([]*entity.Monitorable, error) {
	rows, err := r.db.Query(`SELECT id, name, url FROM monitorables`)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var monitorables []*entity.Monitorable

	for rows.Next() {
		var monitorable entity.Monitorable
		err := rows.Scan(&monitorable.ID, &monitorable.Name, &monitorable.URL)
		if err != nil {
			log.Fatal(err)
		}
		monitorables = append(monitorables, &monitorable)
	}

	return monitorables, nil
}

func (r *MonitorableRepository) Create(dto *dto.AddMonitorable) (*entity.Monitorable, error) {

	query := `INSERT INTO monitorables (name, url) VALUES (?, ?)`
	result, err := r.db.Exec(query, dto.Name, dto.URL)

	fmt.Println(result)
	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	id := int(lastInsertID)

	monitorable := &entity.Monitorable{
		ID:   id,
		Name: dto.Name,
		URL:  dto.URL,
	}

	return monitorable, nil
}

func (r *MonitorableRepository) Update(monitorable *entity.Monitorable) error {
	return nil
}

func (r *MonitorableRepository) DeleteById(id int) error {
	return nil
}
