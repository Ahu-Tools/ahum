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

type ProjectInfo struct {
	Name        string
	PackageName string
	GoVersion   string
	RootPath    string
}

type Project struct {
	Info   ProjectInfo
	Infras []Infra
}

func NewProjectInfo(name, packageName, goVersion, rootPath string) *ProjectInfo {
	return &ProjectInfo{
		name,
		packageName,
		goVersion,
		rootPath,
	}
}

func NewProject(info ProjectInfo, infras []Infra) Project {
	return Project{
		Info:   info,
		Infras: infras,
	}
}
