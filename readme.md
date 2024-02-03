# Receipt Processor Challenge

This repository hosts the code for the take-home exercise assigned when applying to the 'Backend Engineer' position at Fetch (refer: https://fetch.com/). 

This application provides the basic functionality for processing a receipt, as given in the challenge description (refer: https://github.com/fetch-rewards/receipt-processor-challenge). 

### Project Stack:

**Technology:** Go \
**Framework:** Gin Gonic

## Project Execution (via Docker)

To simply test the project deliverables, follow below steps:
1. Download docker desktop (from https://www.docker.com/products/docker-desktop/) and run it minimized
2. Clone the git repository (Use cmd: `git clone https://github.com/rajkumar-b/receipt-processor.git`)
3. Switch to the cloned repository's root and build the docker image via cmd: `docker-compose -f docker/docker-compose.yml build`
4. Spawn up the container from the built image using cmd: `docker-compose -f docker/docker-compose.yml up`
5. Test the APIs
6. Once done, bring down the container either using exit commands like 'CTRL+C' in same terminal or using the command `docker-compose -f docker/docker-compose.yml down` via different terminal in same root.
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