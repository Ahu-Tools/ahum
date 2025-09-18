{{$meth := .MethodName}}
message {{$meth}}Request {
}

message {{$meth}}Response {
    string message = 1;
}