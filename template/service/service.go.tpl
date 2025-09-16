package {{.Service.PackageName}}

{{$data := print .Service.PackageName "data"}}

import (
	{{$data}} "{{.PackageName}}/data/{{.Service.PackageName}}"
)

type Service struct {
	repo {{$data}}.Repo
}

func NewService(repo {{$data}}.Repo) *Service {
	return &Service{
		repo,
	}
}