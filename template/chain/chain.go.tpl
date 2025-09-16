package {{.Service.PackageName}}

{{$data := print .Service.PackageName "data"}}
{{$service := print .Service.Name "svc"}}

import (
	{{$data}} "{{.PackageName}}/data/{{.Service.PackageName}}"
	{{$service}} "{{.PackageName}}/service/{{.Service.PackageName}}"
)

type Chain struct {
	svc *{{$service}}.Service
}

func NewChain() *Chain {
    //TODO: Initialise an implementation for repo
	var repo {{$data}}.Repo
	panic("Unimplemented Repo")

	svc := {{$service}}.NewService(repo)
	return &Chain{
		svc,
	}
}
