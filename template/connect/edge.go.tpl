{{$svc := call .Lowerer .ServiceName}}
{{$vname := call .Lowerer .VersionName}}
{{$genpkg :=print $svc $vname}}
package {{$vname}}

import (
	"context"

	"connectrpc.com/connect"
	{{$genpkg}} "{{.PackageName}}/edge/connect/gen/{{$svc}}/{{$vname}}"

	// @ahum: imports
)

type Edge struct {
}

func NewEdge() *Edge {
	return &Edge{
	}
}

func (e *Edge) Health(c context.Context, req *connect.Request[{{$genpkg}}.HealthRequest]) (*connect.Response[{{$genpkg}}.HealthResponse], error) {
	res := connect.NewResponse(&{{$genpkg}}.HealthResponse{
		Message: "UP",
	})

	return res, nil
}

// @ahum: methods
