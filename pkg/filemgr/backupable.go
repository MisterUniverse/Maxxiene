package filemgr

// Backupable defines the interface for backup operations
type Backupable interface {
	Backup() error
}
