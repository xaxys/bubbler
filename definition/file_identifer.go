package definition

type FileIdentifer struct {
	Name string // file name written in import statement, can be relative or absolute
	Path string // absolute path
}

func (f FileIdentifer) String() string {
	return f.Path
}
