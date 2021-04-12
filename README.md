# payment-gateway

<p align="center">
  <img width="470" height="300" src="https://user-images.githubusercontent.com/22316360/114379738-10ceb200-9b81-11eb-9272-7e6bf9449176.jpg">
</p>

## Introduction
The application uses MongoDB to store the transactions, it uses grpc-gateway to expose the REST on the port `8080` and grpc `50051`.

## Endpoints contract
To test the grpc calls it is possible to use [BloomRPC](https://github.com/uw-labs/bloomrpc) and get the grpc client contract from the proto file of the application [here](https://github.com/avasapollo/payment-gateway/blob/main/web/proto/v1/payment-gateway.proto).
The grpc-gateway generate the swagger to documentation for the REST endpoint, it is possible to see the contract [here](https://github.com/avasapollo/payment-gateway/blob/main/web/proto/v1/payment-gateway.swagger.json)

## Testing
The application contains unit tests, integration tests and e2e tests.
The command to run the unit test is `go test ./... -v`
To run the integration tests and e2e test the command is `go test ./... -v -tags=integration` (MongoDB container is required to run these tests).

## Build and Run the application
It is possible to use the `make all` to run the unit tests and build the application and the `docker-compose up` to build the docker containers for the payment-gateway application and MongoDB.
The `docker-compose` links the container in automatic, the `docker-compose` file is [here](https://github.com/avasapollo/payment-gateway/blob/main/docker-compose.yml).
````
make all
docker-compose up
````
The `docker-compose` expose the ports `8080` (REST) and `50051` (GRPC).

## Test the application
### health
The health endpoint show the status of the application
```curl
curl -X GET \
  http://localhost:8080/health \
  -H 'cache-control: no-cache' \
  -H 'postman-token: e0a69579-575a-3b95-285d-9e7cc58f7c46'
```
response 200 OK
```json
{
    "status": "UP"
}
```
### /v1/authorize
authorize the transaction
```curl
curl -X POST \
  http://localhost:8080/v1/authorize \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: 3366ca89-6aef-2ac2-889f-b675d9a218c2' \
  -d '{
	"card": {
		"name": "Andrea Vasapollo",
		"card_number": "4242424242424242",
		"expire_month": "12",
		"expire_year": "2022",
		"cvv": "123"
	},
	"amount": {
		"value": 100,
		"currency": "EUR"
	}
}'
```
response 200 OK
```json
{
    "result": "ok",
    "authorization_id": "b0b555e2-beb8-4f1a-9485-74bd499c845d",
    "amount": {
        "value": 100,
        "currency": "EUR"
    }
}
```
### /v1/void
void the transaction
```curl
curl -X POST \
  http://localhost:8080/v1/void \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: fffdce69-fd2b-54d0-d21e-616ae78d856e' \
  -d '{
	 "authorization_id": "b0b555e2-beb8-4f1a-9485-74bd499c845d"
}'
```
response 200 OK
```json
{
    "result": "ok",
    "amount": {
        "value": 100,
        "currency": "EUR"
    }
}
```
### /v1/capture
capture the transaction
```curl
curl -X POST \
  http://localhost:8080/v1/capture \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: 2d7b93d8-89f5-8a9b-8d3d-85d9f4225841' \
  -d '{
	 "authorization_id": "4f3b8ee1-7a8b-4dc9-bdae-e63e6f0ad9f5",
	 "amount": 10
}'
```
response 200 OK
```json
{
    "result": "ok",
    "amount": {
        "value": 90,
        "currency": "EUR"
    }
}
```
### /v1/refund
refund the transaction
```curl
curl -X POST \
  http://localhost:8080/v1/refund \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: 4002b371-9664-743a-4527-f2ffd9115242' \
  -d '{
	 "authorization_id": "4f3b8ee1-7a8b-4dc9-bdae-e63e6f0ad9f5",
	 "amount": 1
}'
```
response 200 OK
```json
{
    "result": "ok",
    "amount": {
        "value": 9,
        "currency": "EUR"
    }
}
```

## Consideration
I assumed the amount in the response of the capture endpoint is the difference from the amount of the transaction and what the client has already captured.
I assumed the amount in the response of the refund endpoint is the difference from the amount of the captured and what the client has already refunded.

## Points to improve
- More e2e tests
- On the payer layer I would like to handle better the specific edge cases.

