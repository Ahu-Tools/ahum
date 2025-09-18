{{$svc := call .Lowerer .ServiceName}}
{{$vname := call .Lowerer .VersionName}}
{{$genpkg :=print $svc $vname}}

func (e *Edge) {{.MethodName}}(c context.Context, req *connect.Request[{{$genpkg}}.{{.MethodName}}Request]) (*connect.Response[{{$genpkg}}.{{.MethodName}}Response], error) {
	res := connect.NewResponse(&{{$genpkg}}.{{.MethodName}}Response{
		Message: "Congratulation, {{.MethodName}} is up!",
	})

	return res, nil
}