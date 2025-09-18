{{$svc := call .Lowerer .ServiceName}}
{{$vname := call .Lowerer .VersionName}}

package {{$vname}}

{{$conn := print $svc $vname "connect"}}

import (
	"net/http"

	"{{.PackageName}}/edge/connect/gen/{{$svc}}/{{$vname}}/{{$conn}}"
	// @ahum: imports
)

func RegisterVersion(mux *http.ServeMux) {
	edge := NewEdge()
	path, handler := {{$conn}}.NewServiceHandler(edge)
	mux.Handle(path, handler)
	// @ahum: edges
}
