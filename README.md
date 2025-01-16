## Detailing a bug in connectrpc.com/grpcreflect

### Run Server

```
$ go run server/main.go
```

### Run Client
```
$ go run client/main.go
Fetching acme.weather.v1.WeatherService (this service imports a file in the same package)
acme/weather/v1/service.proto acme.weather.v1
acme/weather/v1/types.proto acme.weather.v1

Fetching acme.weather.v2.WeatherService (split into multiple packages)
acme/weather/v2/service.proto acme.weather.v2
acme/weather/shared/types.proto acme.weather.shared
```
