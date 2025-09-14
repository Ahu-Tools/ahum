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
	RootPath    string
}

type Project struct {
	Info         ProjectInfo
	InfrasConfig []InfraConfig
	InfrasJson   []JSONInfra
}

func NewProjectInfo(name, packageName, rootPath string) *ProjectInfo {
	return &ProjectInfo{
		name,
		packageName,
		rootPath,
	}
}

func NewProject(info ProjectInfo, infrasConfig []InfraConfig, infrasJson []JSONInfra) Project {
	return Project{
		Info:         info,
		InfrasConfig: infrasConfig,
		InfrasJson:   infrasJson,
	}
}
