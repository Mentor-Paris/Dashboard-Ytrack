# Dashboard-Ytrack

## Dependencies

To run this project you need to install and get libraires :

[Golang 1.18](https://go.dev/dl/)
[Gin](https://pkg.go.dev/github.com/gin-gonic/gin#readme-quick-start)

## Run  
First steps, create your module
```
go mod init *YOUR_NAME_MODULE*
```
```
go mod tidy
```
To install Gin, You can use the below Go command
```
go get -u github.com/gin-gonic/gin
```
To run this project
```
go run .
```

## How pages work localhost

The [/userPYC page](http://localhost:8080/userPYC) is the JSON of all users of the PYC campus.

The [/usernational page](http://localhost:8080/usernational) is the JSON of all users of Ytrack.

The [/user/:id page](http://localhost:8080/user/567) is the JSON about one user of the PYC campus.

The [leaderboardPyc page](http://localhost:8080/leaderboard) is the leaderboard of all users of the PYC campus.

The [leaderboardnational page](http://localhost:8080/leaderboardnational) is the leaderboard of all users of Ytrack.

The [stats page](http://localhost:8080/graphique) contains many graphics of stats of all users of Ytrack.

The [progress page](http://localhost:8080/progress) contains the all date of start and end of projects.

The [students page](http://localhost:8080/students) is the JSON about the logs of students.

The [students/:id page](http://localhost:8080/students/33) is the JSON about the logs of one students.

The [go-bot page](http://localhost:8080/go-bot) is the JSON of all informations about users of the PYC campus.
