package db

import "database/sql"

// ItemScanner interface for scanning from an sql.Rows object
type ItemScanner interface {
	ScanRow(*sql.Rows) error
}

// Note struct
type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content []byte `json:"content"`
}

// ScanRow implements the ItemScanner interface for Note
func (n *Note) ScanRow(rows *sql.Rows) error {
	return rows.Scan(&n.ID, &n.Content)
}

// For a File
type File struct {
	ID       int    `json:"id"`
	Filename string `json:"filename"`
	Data     []byte `json:"data"`
}

func (f *File) ScanRow(rows *sql.Rows) error {
	return rows.Scan(&f.ID, &f.Filename, &f.Data)
}

// For a Picture
type Picture struct {
	ID       int    `json:"id"`
	Filename string `json:"filename"`
	Data     []byte `json:"data"`
}

func (p *Picture) ScanRow(rows *sql.Rows) error {
	return rows.Scan(&p.ID, &p.Filename, &p.Data)
}

type HexDump struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Data        string `json:"data"`
}

func (h *HexDump) ScanRow(rows *sql.Rows) error {
	return rows.Scan(&h.ID, &h.Description, &h.Data)
}

type MemoryDump struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Data        []byte `json:"data"`
}

func (m *MemoryDump) ScanRow(rows *sql.Rows) error {
	return rows.Scan(&m.ID, &m.Description, &m.Data)
}
