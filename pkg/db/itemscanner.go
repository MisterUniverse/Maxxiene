package db

import "database/sql"

// ItemScanner interface for scanning from an sql.Rows object
type ItemScanner interface {
	ScanRow(*sql.Rows) error
	NewInstance() ItemScanner
}

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Task      string `json:"task"`
}

// ScanRow implements the ItemScanner interface for todos
func (td *Todo) ScanRow(rows *sql.Rows) error {
	return rows.Scan(&td.ID, &td.Completed, &td.Task)
}

func (t *Todo) NewInstance() ItemScanner {
	return &Todo{}
}

// Note struct
type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content []byte `json:"content"`
}

func (n *Note) ScanRow(rows *sql.Rows) error {
	return rows.Scan(&n.ID, &n.Title, &n.Content)
}

func (n *Note) NewInstance() ItemScanner {
	return &Note{}
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

func (f *File) NewInstance() ItemScanner {
	return &File{}
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

func (p *Picture) NewInstance() ItemScanner {
	return &Picture{}
}

type HexDump struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Data        string `json:"data"`
}

func (h *HexDump) ScanRow(rows *sql.Rows) error {
	return rows.Scan(&h.ID, &h.Description, &h.Data)
}

func (h *HexDump) NewInstance() ItemScanner {
	return &HexDump{}
}

type MemoryDump struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Data        []byte `json:"data"`
}

func (m *MemoryDump) ScanRow(rows *sql.Rows) error {
	return rows.Scan(&m.ID, &m.Description, &m.Data)
}

func (m *MemoryDump) NewInstance() ItemScanner {
	return &MemoryDump{}
}
