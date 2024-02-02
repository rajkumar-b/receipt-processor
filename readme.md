# Receipt Processor Challenge

This repository hosts the code for the take-home exercise assigned when applying to the 'Backend Engineer' position at Fetch (refer: https://fetch.com/). 

This application provides the basic functionality for processing a receipt, as given in the challenge description (refer: https://github.com/fetch-rewards/receipt-processor-challenge). 

### Project Stack:

**Technology:** Go \
**Framework:** Gin Gonic


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