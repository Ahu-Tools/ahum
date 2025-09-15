package {{.Service.PackageName}}

//Use this file to define your models and their validators

{{$name := .Service.Name}}
{{$repo := print .Service.Name "Repo"}}

type {{$repo}} interface{
    Create({{$name}}) ({{$name}}, error)
    Find({{$name}}) ({{$name}}, error)
    Update({{$name}}) ({{$name}}, error)
    Delete(id uint) error
}