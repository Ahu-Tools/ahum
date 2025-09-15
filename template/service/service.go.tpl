package {{.Service.PackageName}}

{{$service := print .Service.Name "Service"}}
{{$repo := print .Service.Name "Repo"}}

import (
	data "{{.PackageName}}/data/{{.Service.PackageName}}"
)

type {{$service}} struct {
	repo data.{{$repo}}
}

func New{{$service}}(repo data.{{$repo}}) *{{$service}} {
	return &{{$service}}{
		repo,
	}
}