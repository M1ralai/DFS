package model

type File struct {
	ID       string `db:"id"`
	UserID   string `db:"user_id"`
	FileName string `db:"filename"`
	Size     int    `db:"size"`
}

type Chunk struct {
	ID     string   `db:"id"`
	FileID string   `db:"file_id"`
	Nodes  []string `db:"nodes"`
}
