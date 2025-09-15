package {{.Service.PackageName}}

//Use this file to define your models and their validators

{{$name := .Service.Name}}

type {{$name}} struct{
    ID *uint
}

func New{{$name}}() *{{$name}} {
    return &{{$name}}{
    }
}