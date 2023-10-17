# Go Traffic Capture

A simple Go server to capture traffic and send it to Elasticsearch. The main idea is that
when mirroring traffic, it'll allow you to capture the body and headers for monitoring
or replaying events against a test environment.

## Toolset

- [DB Snapshot Replication](https://github.com/gemmadlou/go-db-snapshot-replication)
- [Traffic Capture](https://github.com/gemmadlou/go-traffic-capture)
- [Traffic Replay](https://github.com/gemmadlou/go-traffic-replay)

## Usage

Set up environment variables.

```bash
cp .env.example .env
```

Setup local ES server.

```bash
docker compose up -d
```

Start Go server.

```bash
go run main.go

# Development server
nodemon --exec go run main.go --signal SIGTERM
```

Send test request to Go server

```bash
# Using httpie
http http://localhost:9990/eskimo data=123 pet=cat 'Cookie: Monster'  
```

## Authentication

Only Basic auth is configurable.

```bash
AUTH=basic
AUTH_USERNAME=username
AUTH_PASSWORD=password
```

## Resources

- Http server - https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04
- Maps in golang. https://gobyexample.com/maps
- Elasticsearch setup. https://geshan.com.np/blog/2023/06/elasticsearch-docker/
```bash
# Create index
curl -X PUT http://localhost:9200/traffic

# Create example traffic
curl -X POST -H 'Content-Type: application/json' -d '{ "request_uri": "/hello-world", "base_url": "https://example.com", "headers": [], "body": {} }' http://localhost:9200/traffic/_doc

# Get traffic
curl -X GET "localhost:9200/traffic/_search"
```
- Create json data. https://www.digitalocean.com/community/tutorials/how-to-use-json-in-go
- https://christiangiacomi.com/posts/simple-put-patch-request-go/
- https://opster.com/guides/elasticsearch/how-tos/elasticsearch-delete-index/
```bash
curl -X DELETE "http://localhost:9200/traffic"
```
- https://www.alexedwards.net/blog/basic-authentication-in-go
- Timestamp. https://yourbasic.org/golang/current-time/
- Elastcsearh data format. https://stackoverflow.com/questions/38790030/elasticsearch-date-format
- 