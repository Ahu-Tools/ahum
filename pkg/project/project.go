package project

type Edge uint
type Database uint

const (
	CONNECT Edge = iota
	GIN
)

const (
	POSTGRES Database = iota
)

type Project struct {
	Name        string
	PackageName string
	RootPath    string

	Edges []Edge
	Dbs   []Database
}

func NewProject(name, rootPath string) Project {
	return Project{
		Name:     name,
		RootPath: rootPath,
	}
}
