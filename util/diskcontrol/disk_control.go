package diskcontrol

type DiskControler interface {
	Load() error
	Save() error
}
