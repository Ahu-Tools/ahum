package {{.Service.PackageName}}

{{$chain := print .Service.Name "Chain"}}
{{$service := print .Service.Name "Service"}}
{{$repo := print .Service.Name "Repo"}}

import (
	data "{{.PackageName}}/data/{{.Service.PackageName}}"
	service "{{.PackageName}}/service/{{.Service.PackageName}}"
)

type {{$chain}} struct {
	svc *service.{{$service}}
}

func New{{$chain}}() *{{$chain}} {
    //TODO: Initialise an implementation for repo
	var repo data.{{$repo}}
	panic("Unimplemented {{$repo}}")

	svc := service.New{{$service}}(repo)
	return &{{$chain}}{
		svc,
	}
}
