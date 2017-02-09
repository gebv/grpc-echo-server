# grpc-echo-server

``` proto3
syntax = "proto3";

package echo;

service Server {
  rpc Echo (EchoDTO) returns (EchoDTO) {}
}

message EchoDTO {
  string str = 1;
  bytes raw = 2;
}
```