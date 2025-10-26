package ownership

type FileOwnership struct {
	file  string
	user  string
	uid   string
	group string
	gid   string
}

func (fo *FileOwnership) User() string  { return fo.user }
func (fo *FileOwnership) UID() string   { return fo.uid }
func (fo *FileOwnership) Group() string { return fo.group }
func (fo *FileOwnership) GID() string   { return fo.gid }

func New(filepath string) *FileOwnership {
	fo := &FileOwnership{
		file:  filepath,
		user:  "?",
		uid:   "?",
		group: "?",
		gid:   "?",
	}
	fo.Update()
	return fo
}
