package repository

import (
	"github.com/M1ralai/DFS/src/internal/module/client/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ClientCommRepository struct {
	db *sqlx.DB
}

func NewClientCommRepository(db *sqlx.DB) IClientCommRepository {
	return ClientCommRepository{
		db: db,
	}
}

func (r ClientCommRepository) PostFile(f model.File) error {
	if _, err := r.db.NamedExec(`INSERT INTO files (id, user_id, filename, size)
	VALUES(:id, :user_id, :filename, :size);
	`, f); err != nil {
		return err
	}
	return nil
}

func (r ClientCommRepository) GetFile(id uuid.UUID) (model.File, error) {
	v := new(model.File)
	if err := r.db.Get(v, `SELECT * FROM files WHERE id = $1 ;`, id); err != nil {
		return model.File{}, err
	}
	return *v, nil
}

func (r ClientCommRepository) GetAllFileUser(id uuid.UUID) ([]model.File, error) {
	v := make([]model.File, 0)
	if err := r.db.Select(&v, `SELECT * FROM files WHERE user_id = $1;`, id); err != nil {
		return nil, err
	}
	return v, nil
}

func (r ClientCommRepository) DeleteFile(id uuid.UUID) error {
	if _, err := r.db.Exec(`DELETE FROM files WHERE id = $1 ;`, id); err != nil {
		return err
	}
	return nil
}

func (r ClientCommRepository) PostChunk(c model.Chunk) error {
	if _, err := r.db.NamedExec(`INSERT INTO chunks (id, file_id, nodes)
	VALUES (:id, :file_id, :nodes);`, c); err != nil {
		return err
	}
	return nil
}

func (r ClientCommRepository) GetChunk(id uuid.UUID) (model.Chunk, error) {
	v := new(model.Chunk)
	if err := r.db.Get(v, `SELECT * FROM chunks WHERE id = $1;`, id); err != nil {
		return model.Chunk{}, err
	}
	return *v, nil
}

func (r ClientCommRepository) GetChunksByFileID(id uuid.UUID) ([]model.Chunk, error) {
	v := make([]model.Chunk, 0)
	if err := r.db.Select(&v, `SELECT * FROM chunks WHERE file_id = $1;`, id); err != nil {
		return nil, err
	}
	return v, nil
}

func (r ClientCommRepository) DeleteChunksByFileID(id uuid.UUID) error {
	if _, err := r.db.Exec(`DELETE FROM chunks WHERE file_id = $1;`, id); err != nil {
		return err
	}
	return nil
}

func (r ClientCommRepository) DeleteChunk(id uuid.UUID) error {
	if _, err := r.db.Exec(`DELETE FROM chunks WHERE id = $1`, id); err != nil {
		return err
	}
	return nil
}
