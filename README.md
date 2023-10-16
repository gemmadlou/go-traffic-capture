# Go Traffic Capture

## Usage

Setup local ES server

```bash
docker compose up -d
```

Create new index

```bash
curl -X PUT http://localhost:9200/traffic
```

> Returns `{"acknowledged":true,"shards_acknowledged":true,"index":"traffic"}`

## Resources

- Http server - https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04
- Maps in golang. https://gobyexample.com/maps
- Elasticsearch setup. https://geshan.com.np/blog/2023/06/elasticsearch-docker/
```bash
# Create example traffic
curl -X POST -H 'Content-Type: application/json' -d '{ "request_uri": "/hello-world", "base_url": "https://example.com", "headers": [] }' http://localhost:9200/traffic/_doc

# Get traffic
curl -X GET "localhost:9200/traffic/_search"
```