package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Rodrigueslcs/challenge-backend/entity"
)

type VideoMySQL struct {
	db *sql.DB
}

func NewVideoMySQL(db *sql.DB) *VideoMySQL {
	return &VideoMySQL{
		db: db,
	}
}

func (r *VideoMySQL) List() ([]*entity.Video, error) {
	smtp, err := r.db.Prepare(`select id, title, description, url from video `)
	if err != nil {
		return nil, err
	}
	var videos []*entity.Video
	rows, err := smtp.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var v entity.Video
		err = rows.Scan(&v.ID, &v.Title, &v.Description, &v.URL)
		if err != nil {
			return nil, err
		}
		videos = append(videos, &v)
	}
	return videos, nil
}

func (r *VideoMySQL) Get(id int) (*entity.Video, error) {
	smtp, err := r.db.Prepare(`select id, title, description, url from video where id = ?`)
	if err != nil {
		return nil, err
	}
	var video entity.Video
	rows, err := smtp.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&video.ID, &video.Title, &video.Description, &video.URL)
	}
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	if video.ID < 1 {
		return nil, errors.New("video nao encontrado")
	}
	return &video, nil
}

func (r *VideoMySQL) Create(e *entity.Video) (int, error) {
	stmt, err := r.db.Prepare(`
		insert into video ( title, description, url) 
		values(?,?,?)`)
	if err != nil {
		return 0, err
	}
	resp, err := stmt.Exec(
		e.Title,
		e.Description,
		e.URL,
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

func (r *VideoMySQL) Update(e *entity.Video) error {
	stmt, err := r.db.Prepare("update video set title = ?, description = ?, url = ? where id = ?")
	if err != nil {
		return err
	}
	resp, err := stmt.Exec(
		e.Title,
		e.Description,
		e.URL,
		e.ID,
	)
	if err != nil {
		return err
	}

	_, err = resp.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (r *VideoMySQL) Delete(id int) error {
	_, err := r.db.Exec("delete from video where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
