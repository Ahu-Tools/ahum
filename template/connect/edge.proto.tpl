syntax = "proto3";

{{$svc := call .Lowerer .ServiceName}}
{{$vname := call .Lowerer .VersionName}}

package {{call .Snaker (print .ServiceName " " .VersionName)}};

option go_package = "{{.PackageName}}/edge/connect/gen/{{$svc}}/{{$vname}};{{print $svc $vname}}";

service Service {
  rpc Health(HealthRequest) returns (HealthResponse) {
  }

  // @ahum: methods
}

message HealthRequest {
}

message HealthResponse {
  string message = 1;
}

// @ahum: messages
