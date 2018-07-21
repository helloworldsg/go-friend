# :two_men_holding_hands: go-friend [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![Go Report Card][rpt-img]][rpt]

### How to run
1. Build: `make build`
2. Run web server: `./go-friend`
3. Run unit test: `make test`

### Alternative way to run
1. Use go get: `go get github.com/hmoniaga/go-friend`
2. Run: `go-friend`

## Assumptions
### General
1. System assume it expects all inputs with case sensitive
2. System will perform basic validation before it executes command
3. Build will generate `go-friend` binary which listen to 8080 port by default.
4. No security is to be implemented
5. A sample, custom basic in-memory database is implemented

## Functionality

### Add Friend
1. create a friend connection between two email addresses only
2. both email addresses provided are already registered in the system
3. automatically subscribes each other's updates

### List Friends
1. retrieve the friends list for an email address

### List Mutual Friends
1. retrieve the common friends list between two email addresses
2. both email addresses provided are already registered in the system

### Follow, i.e. receive updates
1. subscribes updates a user
2. if target is inside user's block list, it will be rejected 

### Block
1. both email addresses provided are already registered in the system
2. block updates from an email address
3. if they are not connected as friends, then he/she will no longer be able to add as friend
4. if connected as friends, then he/she will no longer receive notifications from the blocked

### Notify
1. returns list of eligible subscribers that will receive updates

## Design
### Concurrency
1. Concurrency is required as this is a REST API serving multiple users
2. System is accessible only via REST API

### Data Design
1. Data is designed to be easy to horizontally scale, i.e. partitioned by user.

### Software Design
1. System is developed using Go (Golang)
2. It uses standard libraries to achieve the solution and to simplify dependencies

### Testing
1. Unit test and coverage is provided and uses standard libraries

[doc-img]: https://godoc.org/github.com/hmoniaga/go-friend?status.svg
[doc]: https://godoc.org/github.com/hmoniaga/go-friend
[ci-img]: https://travis-ci.org/hmoniaga/go-friend.svg?branch=master
[ci]: https://travis-ci.org/hmoniaga/go-friend
[cov-img]: https://codecov.io/gh/hmoniaga/go-friend/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/hmoniaga/go-friend
[rpt-img]: https://goreportcard.com/badge/hmoniaga/go-friend
[rpt]: https://goreportcard.com/report/hmoniaga/go-friend