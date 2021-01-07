# go-influxdb

```
$ docker run --name influxdb -d -p 8086:8086 -v $PWD:/var/lib/influxdb influxdb
$ docker exec -u 0 -it influxdb sh
# influx
Connected to http://localhost:8086 version 1.8.3
InfluxDB shell version: 1.8.3
>
> create database go_influx
```

- go-chi chi router: This is a lightweight, idiomatic and composable router for building Go HTTP services.
- InfluxDB Go Client
- UI Library tachyons.io is a library for creating fast loading, highly readable, and 100% responsive interfaces with as little css as possible.