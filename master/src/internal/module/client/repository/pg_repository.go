package repository

import (
	"github.com/M1ralai/DFS/src/internal/module/client/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ClientRepository struct {
	db *sqlx.DB
}

func NewClientRepository(db *sqlx.DB) IClientRepository {
	return &ClientRepository{
		db: db,
	}
}

func (r *ClientRepository) PostFile(f model.File) error {
	if _, err := r.db.NamedExec(`INSERT INTO files (id, user_id, filename, size)
	VALUES(:id, :user_id, :filename, :size);
	`, f); err != nil {
		return err
	}
	return nil
}

func (r *ClientRepository) GetFile(id uuid.UUID) (model.File, error) {
	v := new(model.File)
	if err := r.db.Get(v, `SELECT * FROM files WHERE id = $1 ;`, id); err != nil {
		return model.File{}, err
	}
	return *v, nil
}

func (r *ClientRepository) GetAllFileUser(id uuid.UUID) ([]model.File, error) {
	v := make([]model.File, 0)
	if err := r.db.Select(&v, `SELECT * FROM files WHERE user_id = $1;`, id); err != nil {
		return nil, err
	}
	return v, nil
}

func (r *ClientRepository) DeleteFile(id uuid.UUID) error {
	if _, err := r.db.Exec(`DELETE FROM files WHERE id = $1 ;`, id); err != nil {
		return err
	}
	return nil
}

func (r *ClientRepository) PostChunk(c model.Chunk) error {
	if _, err := r.db.NamedExec(`INSERT INTO chunks (id, file_id, nodes)
	VALUES (:id, :file_id, :nodes);`, c); err != nil {
		return err
	}
	return nil
}

func (r *ClientRepository) GetChunk(id uuid.UUID) (model.Chunk, error) {
	v := new(model.Chunk)
	if err := r.db.Get(v, `SELECT * FROM chunks WHERE id = $1;`, id); err != nil {
		return model.Chunk{}, err
	}
	return *v, nil
}

func (r *ClientRepository) GetChunksByFileID(id uuid.UUID) ([]model.Chunk, error) {
	v := make([]model.Chunk, 0)
	if err := r.db.Select(&v, `SELECT * FROM chunks WHERE file_id = $1;`, id); err != nil {
		return nil, err
	}
	return v, nil
}

func (r *ClientRepository) DeleteChunksByFileID(id uuid.UUID) error {
	if _, err := r.db.Exec(`DELETE FROM chunks WHERE file_id = $1;`, id); err != nil {
		return err
	}
	return nil
}

func (r *ClientRepository) DeleteChunk(id uuid.UUID) error {
	if _, err := r.db.Exec(`DELETE FROM chunks WHERE id = $1`, id); err != nil {
		return err
	}
	return nil
}

func (r *ClientRepository) IncrementReplicaCount(chunkID uuid.UUID) (int, error) {
	var count int
	if err := r.db.Get(&count, `UPDATE chunks SET replica_count = replica_count + 1 WHERE id = $1 RETURNING replica_count`, chunkID); err != nil {
		return 0, err
	}
	return count, nil
}
