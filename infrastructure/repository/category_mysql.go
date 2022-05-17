package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Rodrigueslcs/challenge-backend/entity"
)

type CategoryMySQL struct {
	db *sql.DB
}

func NewCategoryMySQL(db *sql.DB) *CategoryMySQL {
	return &CategoryMySQL{
		db: db,
	}
}

func (r *CategoryMySQL) List() ([]*entity.Category, error) {
	smtp, err := r.db.Prepare(`select id, title, color from category `)
	if err != nil {
		return nil, err
	}
	var categories []*entity.Category
	rows, err := smtp.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c entity.Category
		err = rows.Scan(&c.ID, &c.Title, &c.Color)
		if err != nil {
			println(err.Error())
			return nil, err
		}
		categories = append(categories, &c)
	}
	fmt.Println(categories)
	return categories, nil
}

func (r *CategoryMySQL) Get(id int) (*entity.Category, error) {
	smtp, err := r.db.Prepare(`select id, title, color from category where id = ?`)
	if err != nil {
		return nil, err
	}
	var category entity.Category
	rows, err := smtp.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&category.ID, &category.Title, &category.Color)
	}
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	if category.ID < 1 {
		return nil, errors.New("categoria nao encontrado")
	}
	return &category, nil
}

func (r *CategoryMySQL) Create(e *entity.Category) (int, error) {
	stmt, err := r.db.Prepare(`
		insert into category ( title, color) 
		values(?,? )`)
	if err != nil {
		return 0, err
	}
	resp, err := stmt.Exec(
		e.Title,
		e.Color,
	)
	if err != nil {
		return 0, err
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *CategoryMySQL) Update(e *entity.Category) error {
	stmt, err := r.db.Prepare("update category set title = ?, color = ? where id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		e.Title,
		e.Color,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryMySQL) Delete(id int) error {
	_, err := r.db.Exec("delete from category where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
