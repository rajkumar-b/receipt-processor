# Receipt Processor Challenge

This repository hosts the code for the take-home exercise assigned when applying to the 'Backend Engineer' position at Fetch (refer: https://fetch.com/). 

This application provides the basic functionality for processing a receipt, as given in the challenge description (refer: https://github.com/fetch-rewards/receipt-processor-challenge). 

### Project Stack:

**Technology:** Go \
**Framework:** Gin Gonic \
**REST Documentation Middleware** Swagger UI

## Project Execution (via Docker)

To simply test the project deliverables, follow below steps:
1. Download docker desktop (from https://www.docker.com/products/docker-desktop/) and run it minimized
2. Clone the git repository (Use cmd: `git clone https://github.com/rajkumar-b/receipt-processor.git`)
3. Switch to the cloned repository's root and build the docker image via cmd: `docker-compose -f docker/docker-compose.yml build`
4. Spawn up the container from the built image using cmd: `docker-compose -f docker/docker-compose.yml up`
5. Test the APIs
6. Once done, bring down the container either using exit commands like 'CTRL+C' in same terminal or using the command `docker-compose -f docker/docker-compose.yml down` via different terminal in same root.


## Testing the APIs

Easy way to go, is to just use the Swagger UI to test the endpoints. Link: http://localhost:8080/swagger/index.html

If the choice of preference is Postman, or other similar tools, follow instruction below:
1. Install Postman (from https://www.postman.com/downloads/ ; or similar tool)
2. Make sure docker is up and running (till step 5 of Project Execution)
3. Validate a simple ping request by choosing `GET` method and using url ```http://localhost:8080/ping``` .
It should return a valid JSON reponse with a message as 'pong' on hitting SEND.
4. To send and save a receipt to the system, choose `POST` and use the url ```http://localhost:8080/receipts/process``` .
Go to `Body` section and choose `raw` button, followed by `JSON` in the dropdown. 

Try any of the below examples to send data to the server. 


#### Example 1

```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
```

#### Example 2

```json
{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}
```

Note down the received Receipt ID to test out the next end point. Mix match different values in above examples to try different other test cases.

5. To Get the points for a receipt, simply plug the `id` from above step into the url ```http://localhost:8080/receipts/{id}/points``` (remove curly braces) and use a `GET` request like earlier to get the JSON reply. 

Here is the breakdown of points for above two examples:

#### Example 1

```text
Total Points: 28
Breakdown:
     6 points - retailer name has 6 characters
    10 points - 4 items (2 pairs @ 5 points each)
     3 Points - "Emils Cheese Pizza" is 18 characters (a multiple of 3)
                item price of 12.25 * 0.2 = 2.45, rounded up is 3 points
     3 Points - "Klarbrunn 12-PK 12 FL OZ" is 24 characters (a multiple of 3)
                item price of 12.00 * 0.2 = 2.4, rounded up is 3 points
     6 points - purchase day is odd
  + ---------
  = 28 points
```

#### Example 2

```text
Total Points: 109
Breakdown:
    50 points - total is a round dollar amount
    25 points - total is a multiple of 0.25
    14 points - retailer name (M&M Corner Market) has 14 alphanumeric characters
                note: '&' is not alphanumeric
    10 points - 2:33pm is between 2:00pm and 4:00pm
    10 points - 4 items (2 pairs @ 5 points each)
  + ---------
  = 109 points
```

<br><br><br>

If the choice of preference is commandline/terminal via curl, follow instructions below:

1. For ping test, use ```curl -X 'GET' \
  'http://localhost:8080/ping' \
  -H 'accept: application/json'```

2. For sending and storing a receipt, use ```curl -X 'POST' \
  'http://localhost:8080/receipts/process' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '<payload>'```

Example
```
curl -X 'POST' \
  'http://localhost:8080/receipts/process' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}'
```


3. For sending and storing a receipt, use ```curl -X 'GET' \
  'http://localhost:8080/receipts/<receipt-id>/points' \
  -H 'accept: application/json'```

Example
```
curl -X 'GET' \
  'http://localhost:8080/receipts/fe7e63ba-0841-4c94-b60a-8bf794f0400c/points' \
  -H 'accept: application/json'
```

<br><br><br>
Side Note:

The api.yml file gives the regex for Retailers under a receipt as `"^[\\w\\s\\-]+$"` which is essentially alphabets, numbers, underscores, hyphens and spaces. However, the example uses `"&"` both in the yml file, as well as the calculatePoints example 2 in readme. So, including `&` into the validator regex for now. However, I strongly suggest it is taken care of somewhere, either the regex updated, or the retailer name. 
<br><br><br><br>

_______________________________________ 
_______________________________________ 
`For Developers` 
_______________________________________ 
_______________________________________

## Project Creation
Before creating the project, ensure that the following are installed:

1. Go language: https://go.dev/doc/install
2. GitHub CLI: https://cli.github.com/

For initial project creation upto the initial commit, follow below steps:

1. Initialize a local git repository: `git init receipt-processor`
2. Move to the created directory: `cd receipt-processor`
3. Create a corresponding GitHub repo: `gh repo create --public receipt-processor`
4. Link the local git repo to the remote GitHub repo: `git remote add origin https://github.com/<git-profile>/receipt-processor.git`
5. Initialize a Go project: `go mod init <site-name>/receipt-processor`
6. Add all changes to local repo: `git add .`
7. Commit the changes to local repo: `git commit -m "initial commit"`
8. Push the changes to the remote branch (master): `git push --set-upstream origin master`

Now the basic backbone of the Go project is ready.

To setup go with Gin Gonic Framework, use: `go get -u github.com/gin-gonic/gin` (refer: https://gin-gonic.com/docs/quickstart/)

This project uses default gitignore taken from github for Go (refer: https://github.com/github/gitignore/blob/main/Go.gitignore) during initialization. Any follow-up changes can be tracked via respective branch.

## Project Execution (in Development / Debug mode)

To run the project, simply use `go run .` or `go run main.go` and follow the link: http://localhost:8080/ping to execute a basic server ping (GET request via REST API).

To build and run the project, use `go build .` followed by `.\receipt-processor.exe` to initiate the server on localhost. Follow the same link as above to hit the API endpoints.

## Project Tests

To run the test on handler, simply use `go test rajkumar.app/receipt-processor/handler`.

To run the test with coverage report, use `go test --cover rajkumar.app/receipt-processor/handler`.

To run the coverage report for all tests, use `go test ./... -coverprofile cover.out`, followed by `go tool cover -func cover.out`.

To generate a html cover report, run the former command from above line, followed by `go tool cover -html cover.out -o coverage.html`. The coverage details can be seen via the output html file `coverage.html`.

## Project Extras
To add Swagger UI for REST APIs, follow the instructions from https://lemoncode21.medium.com/how-to-add-swagger-in-golang-gin-6932e8076ec0
