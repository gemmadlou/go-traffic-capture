# Go Traffic Capture

## Usage

Setup local ES server

```bash
docker compose up -d
```

Test http request

```bash
http http://localhost:9990/eskimo data=123  'Cookie: Me'
```

## Resources

- Http server - https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04
- Maps in golang. https://gobyexample.com/maps
- Elasticsearch setup. https://geshan.com.np/blog/2023/06/elasticsearch-docker/
```bash
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