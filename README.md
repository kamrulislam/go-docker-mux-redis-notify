# GO Mux Redis Sendgrid Email Docker App

It is a REST Go server which uses Redis to save data and also displays the use of a third party email sending using SendGrid. It also provides a docker-compose to build and run all together. 

## Code Structure

The code is best run inside docker containers. To make the process simple, the project includes a `docker-compose.yml` file. It contains two containers `api` and `redis`. Before running the application, the api container needs to be built. A `Dockerfile` is provided to make the build process easy.

All application code sits inside `commands` folder other than docker setup and Go Module files. The project is built using  `Go Module` which comes with go version `go1.13`. All necessary build will be handled by the docker setup itself. 

## Setup

To build the application please follow the following steps:

- Clone the repository and go inside the cloned repository
```bash
git clone https://github.com/kamrulislam/go-docker-mux-redis-notify.git
cd go-docker-mux-redis-notify
```
- Build the API docker image first using `docker-compose build api` command
- Run the application using `docker-compose up` command
- Wait until you see something like the following from API(`api_1`) and Redis(`cache`)
```bash
cache    | 1:M 08 Dec 2019 01:13:07.716 * Ready to accept connections
api_1    | Server listening!
```
- The `api` runs on `8080` port and also maps the same port on host machine.


## Checking Process

To create a customer either you can use any API Testing Client (eg., Postman, Talend) or you can use `curl` command in the terminal. Let's see how the functionalities of the application work.

### Create a Customer

- Endpoint: `http://localhost:8080/customer`
- Request: POST
- Body: 
```json
{
  "Name" : "Kamrul",
  "Email" : "kamruli@gmail.com",
  "Phone" : "0400000000"
}
```
- cURL:
```bash
curl 'http://localhost:8080/customer' -H 'Connection: keep-alive'  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36' -H 'Content-Type: application/json' -H 'Accept: */*' -H 'Sec-Fetch-Site: cross-site' -H 'Sec-Fetch-Mode: cors' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-AU,en-GB;q=0.9,en-US;q=0.8,en;q=0.7'  --data-binary $'{\n  "Name" : "Kamrul",\n  "Email" : "kamruli@gmail.com",\n  "Phone" : "0400000000"\n}' --compressed
```
- Response:
```
{"Name":"Kamrul","Email":"kamruli@gmail.com","Phone":"0400000000"}
```

### Find a customer by email

- Endpoint: `http://localhost:8080/customer/kamruli@gmail.com`
- Request: GET

- cURL:
```bash
curl 'http://localhost:8080/customer/kamruli@gmail.com' -H 'Connection: keep-alive'  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36' -H 'Content-Type: application/json' -H 'Accept: */*' -H 'Sec-Fetch-Site: cross-site' -H 'Sec-Fetch-Mode: cors' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-AU,en-GB;q=0.9,en-US;q=0.8,en;q=0.7'  --compressed
```
- Response:
```
{"Name":"Kamrul","Email":"kamruli@gmail.com","Phone":"0400000000"}
```

### Send notification to a customer identified by email


- Endpoint: `http://localhost:8080/notify/kamruli@gmail.com`
- Request: GET

- cURL:
```bash
curl 'http://localhost:8080/notify/kamruli@gmail.com' -H 'Connection: keep-alive'  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36' -H 'Content-Type: application/json' -H 'Accept: */*' -H 'Sec-Fetch-Site: cross-site' -H 'Sec-Fetch-Mode: cors' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-AU,en-GB;q=0.9,en-US;q=0.8,en;q=0.7'  --compressed
```
- Response:
```
{"success":true}
```


