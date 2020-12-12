# Go Training

[Go Course](http://go-course.org) lab assignments repository.

Repo originally forked from [anz-bank/go-course](https://github.com/joshcarp/afterwork-go)
## Instructions
- In each directory there is a problem to solve. 
- Most questions will ask for specific function signatures, but don't be afraid to change anything to become more testable.
- One PR should contain __ONLY__ one solution.
- branch names should follow the convention `<username>/<problem>`
- Required status checks on linting actions and testing actions, with test coverage above 90% of code lines required.
- If you think there could be an improvement in instructions, or if you think there are good resources for a specific problem, feel free to open a PR.

## Prerequisites

-   Install `go 1.15` according to [official installation instruction](https://golang.org/doc/install)
-   Install `golangci-lint` according to [instructions](https://github.com/golangci/golangci-lint#local-installation)

## Build, execute, test, lint

Run the hello-world sample with

    go run 00_hello_world/main.go

Alternative build and execute a binary with

    go build -o hello_world ./00_hello_world
    ./hello_world

Test it with

    go test ./...

Lint it with

    golangci-lint run

Review coverage with

    go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## Pre-PR checklist

-   Ensure your source code changes
    -   Build
    -   Test
    -   Lint
    -   Have above 80% test coverage
-   Ensure [good commit messages](https://chris.beams.io/posts/git-commit/)
    -   Separate subject from body with a blank line
    -   Limit the subject line to 60 characters
    -   Use the imperative mood in the subject line
    -   Do not end the subject line with a period
    -   Wrap the body at 80 characters
    -   Use the body to explain what and why vs. how
    -   Use `git rebase -i COMMIT_HASH` to rework your commits if necessary
-   Review the "Files changed" section of your PR
-   "Think of the reviewer: your code needs to be reviewable and that should be a prime concern when writing your code, commits and PRs." (@camh-anz)
-   Review a colleague's PR and add a link to the review in your own PR description

