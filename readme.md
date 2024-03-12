# Online Payment Platform and Bank Simulator

This project consists of an online payment platform API and a bank simulator to test the API.

## Prerequisites

- Go 1.16 or higher
- Docker (for mongoDB)

## Setup

1. Clone the repository to your local machine.

2. Navigate to the project directory.

3. If you're using Docker, you can start a MongoDB instance using the following command:

```bash
docker build -t my-mongo .
docker run -d -p 27017:27017 --name mongodb my-mongo
```

4. Install the project dependencies with the following command:

```bash
go mod download
```

## Running

1. Start the online payment platform API server with the following command:

```bash
go run main.go
```

The API server will start listening on port 3000.

2. Open another terminal and start the bank simulator with the following command:

```bash
go run bank/bank.go
```

The bank simulator will start listening on port 4000.

## Usage

The online payment platform API has the following endpoints:

- `POST /login`: Generate the JWT needed for the other calls 
- `POST /process_payment`: Processes a payment. Accepts a JSON with payment details in the request body.
- `GET /get_payment_details/:transaction_id`: Retrieves the details of a previous payment.
- `POST /process_refund`: Processes a refund. Accepts a JSON with refund details in the request body.

The bank simulator has the following endpoint:

- `POST /process_payment`: Simulates the processing of a payment. Returns a JSON with a simulated response.
- `POST /refund_payment`: Simulates the refund of a payment. Returns a JSON with a simulated response.

## Testing

You can test the online payment platform API by making requests to the above endpoints. For example, you can use `curl` or Postman to make the requests.

```bash
curl -X POST -H "Content-Type: application/json" -d '{"username":"john","password":"doe"}' http://localhost:3000/login
```
Use a valid token
```bash
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNzEwNTM0MjA4LCJuYW1lIjoiSm9obiBEb2UifQ.4WyIxrq1eztMtGk8KgKHunGMrvgNTNur-56C7NdqJLc" -d '{"merchant_id":"66d16a2a-2959-45c0-894d-11b639e7e4d1","amount":10.5,"card_number":"4111111111111111","card_expiry":"042026","cvv":"333"}' http://localhost:3000/process_payment
```
Use a valid token and a valid transaction_id
```bash
curl -X GET -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNzEwNTM0MjA4LCJuYW1lIjoiSm9obiBEb2UifQ.4WyIxrq1eztMtGk8KgKHunGMrvgNTNur-56C7NdqJLc" http://localhost:3000/get_payment_details/c61a266a-0b2c-4a4c-88b4-19e95d5ed61d
```

Use a valid token and a valid transaction_id
```bash
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNzEwNTM0MjA4LCJuYW1lIjoiSm9obiBEb2UifQ.4WyIxrq1eztMtGk8KgKHunGMrvgNTNur-56C7NdqJLc" -d '{"transaction_id":"c61a266a-0b2c-4a4c-88b4-19e95d5ed61d"}' http://localhost:3000/refund_payment
```
