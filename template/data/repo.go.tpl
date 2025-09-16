package {{.Service.PackageName}}

//Use this file to define your repository interface

{{$name := .Service.Name}}

type Repo interface{
    Create({{$name}}) ({{$name}}, error)
    Find({{$name}}) ({{$name}}, error)
    Update({{$name}}) ({{$name}}, error)
    Delete(id uint) error
}