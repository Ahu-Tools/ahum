{{$svc := call .Lowerer .ServiceName}}

package {{$svc}}

import (
	"net/http"

	// @ahum: imports
)

func RegisterService(mux *http.ServeMux) {
	// @ahum: versions
}
