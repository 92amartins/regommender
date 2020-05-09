## Regommender
A simple recommender system API.

## Usage

This project relies on [Docker Compose](https://docs.docker.com/compose/) for development and deployment.

The command below will spin up two docker containers. The first holds a Redis instance which serves as the backend for this project. The second holds the rest-api itself.

```
docker-compose up
```

The api is very simple. Currently it only supports two methods: `set_recommendations` and `get_recommendations`.

For setting a recommendation:

```bash
curl -X POST 'localhost:8080/recommendation/' -d '{
    "source": "product1",
    "target": "product2",
    "score": 0.78
}'
```

Then, you can get the recommendation by doing:

```bash
curl 'localhost:8080/recommendations/product1'
# ["product2"]
```

The recommendations are stored in [sorted set objects](https://redis.io/topics/data-types) rather than standard Redis keys, so Redis handle recommendations ordering automatically:

```bash
curl -X POST 'localhost:8080/recommendation/' -d '{
    "source": "product1",
    "target": "product2",
    "score": 0.78
}'

curl -X POST 'localhost:8080/recommendation/' -d '{
    "source": "product1",
    "target": "product3",
    "score": 0.95
}'

curl 'localhost:8080/recommendations/product1'
# ["product3","product2"]
```