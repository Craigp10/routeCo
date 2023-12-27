### RouteCo test project

https://grpc.io/docs/languages/go/basics/

This project is a simple route mapping application, the client may provide coordinates and recieve stats on features within the coordinates (a box created from the two coordinates) or stats on a specific feature if only 1 coordinate is provided. The flow follows:

1. Client creates connection
2. Client makes call with coordinates
3. Server responds with a route and streams features of the route back
4. At the end, the server sends stats about the route

## Why use gRPC in this application?

gRPC is useful because it allows the client to make calls to the server without worrying about a url + endpoints to request. Functions can be called as if they are local. In this case, since the two servers are local to the application, we can also take advantage of the gRPC protobuffers. Where the object structures can be defined once and encoded to Go and utilized by both the client and server to make function calls.

# RouteCo

Defines a routeCo service that contains 4 functions and the necessary structures.

```
  service RouteGuide {...}

  rpc GetFeature(Point) returns (Feature) {}
  rpc ListFeatures(Rectangle) returns (stream Feature) {}
  rpc RecordRoute(stream Point) returns (RouteSummary) {}
  rpc RouteChat(stream RouteNote) returns (stream RouteNote) {}

  message Point
  message Rectangle
  message Feature
  message RouteNote
  message RouteSummary
```

Compiled by calling the below command in the root directory of the project.
`protoc --go_out=. --go_opt=paths=source_relative \
 --go-grpc_out=. --go-grpc_opt=paths=source_relative \
 routeCo/routeCo.proto`

# Server/

`main` starts the server, over a tcp connection. If TLS is enabled then encryptions are checked. Served on port 50051

The server handles routes and features on a coordinate plane. Coordinates can be provided as a single coordinate or a stream/route. The public functions of the server are as follows:

<b>GetFeature</b> given a coordinate, will return the 'feature' at that coordinate. If no feature found then will return no feature

<b>ListFeatures</b> Given two points, a rectangle, will list all of the features as a stream within the rectangle on the coordinate plane

<b>RecordRoute</b> Given a stream of points, returns stats on the provided route.

<b>RouteChat</b> Given a stream of coordinates, returns with all of the messages at each of the provided coordinates.

# Client/

The client utilizies the route co service through hitting the various rpc functions defined. The stream is utilizied by `runRouteChat` passing a stream of messages for a specified feature. `ListFeatures` utilizies a stream to send all features within the provided route box.
