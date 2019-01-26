# pkg-lister

[![CircleCI](https://circleci.com/gh/rsmitty/pkg-lister.svg?style=svg)](https://circleci.com/gh/rsmitty/pkg-lister)

A toy package index that responds to network requests about packages and their dependencies. The goal of this project was to demonstrate an understand of implementing a project in Go, while trying to focus on making the project as "prod ready" as possible.

Quick Info:
- Github: [https://github.com/rsmitty/pkg-lister](https://github.com/rsmitty/pkg-lister)
- Docker Image: `rsmitty/pkg-lister:latest`
- CircleCI Runs: [https://circleci.com/gh/rsmitty/pkg-lister](https://circleci.com/gh/rsmitty/pkg-lister)
- Public Deployment of pkg-lister: `pkg-lister.rsmitty.cloud:8080`

Basic Command Usage:
```
$ ./pkg-lister --help
Usage of ./pkg-lister:
  -port string
    	The port to listen on (default "8080")
```


An online version of this code is available at `pkg-lister.rsmitty.cloud:8080` and can be interacted with via telnet like:
```
$ telnet pkg-lister.rsmitty.cloud 8080
Trying 165.227.254.188...
Connected to pkg-lister.rsmitty.cloud.
Escape character is '^]'.
INDEX|A|
OK
QUERY|B|
FAIL
INDEX|B|A
OK
QUERY|B|
OK
```

## Build

**Note:** The darwin and linux binaries are already included in the `bin/` directory and on the GitHub release page. There is also a pre-built docker image at `rsmitty/pkg-lister:latest`

Using Make:

There is a makefile in the `build/` directory of this repository. Using that file, binaries and a docker image can be built. 
- Issuing `make linux` or `make darwin` will generate the go binary for the corresponding OS and place it in the `bin/` directory.
- Issuing `make docker` will run a docker build/push to create a docker image called `rsmitty/pkg-lister`.
- Issuing `make all` is a combo of the above, essentially running `make linux` followed by `make docker`. 

Dockerfile:

There is a Dockerfile in the `build/` directory that can be used to build docker image. This image is based off of `ubuntu:latest`.
- You must generate binaries first using the make instructions above.
- From the `build/` directory, issue `docker build -t $DOCKER_IMAGE_NAME:$DOCKER_IMAGE_VERSION -f ./Dockerfile ../` to create a local image with the name of your choosing.

## Test

The [go testing](https://golang.org/pkg/testing/) package is used to carry out tests against the pkg-lister code base.

- Issue `go test -v` to run the tests.

The DO package tree tests can also be run against this program. You will need to either pull and run a docker image or run directly from the binaries generated with make above.

For docker:

- Start the container with `docker run -ti -p 8080:8080 rsmitty/pkg-lister:latest`
- Run the DO test suite with `./do-package-tree_linux`, or your OS of choice

For binaries:

- Generate binaries using make instructions above
- Run the binary from the `bin/` directory with a simple `./pkg-lister-linux` or `./pkg-lister-darwin`
- Run DO test suite

## Deploy

Deploying this code can be done a couple of different ways, both based on docker. The second option below would be the recommended route.

First, you can simply do a docker run on some cloud server if desired. This would be very similar to running the docker test instructions above, but you may with to enable auto restart and run in the background: `docker run -d -p 8080:8080 --restart always rsmitty/pkg-lister:latest`

Second, there are Kubernetes manifests included in this repo in the `/deploy` directory. These can be created in a Kubernetes cluster with `kubectl create -f .` from the directory. Note that the pkg-lister service expects a LoadBalancer to be available for getting an IP. This should work fine on a DO managed Kubernetes cluster. In fact, this image is currently running on a cluster of mine out in DO and provides the interaction shown at the top of the README.

## Design Considerations

This section is just highlights some of my thought processes throughout building this project and some explanations for some of the routes taken.

- For the package index, I stuck to just using a hashmap as opposed to bringing a datastore into the mix. Although I really feel like this would be more "prod ready" with a clustered DB on the backend and a scaleable server layer to accept requests, it seemed to contradict the instructions around not bringing external dependencies into the project. But listed below are some of the options I considered. I think Etcd would probably be my choice if I were to pick and implement.
  - `database/sql` would work, but requires external drivers for every sql-type db.
  - Etcd also seemed like it would be a really good fit for persistence here, but again, external dependency.
  - There is also an option for just writing to a flat, local file for persistence. But that didn't feel "cloud-native" or "prod-ready".

- I tried to incorporate CI into the process in a way that I would do so in a real project. CircleCI monitors the git repo and will run go tests and build a new `rsmitty/pkg-lister:latest` image upon commit to master. 

- There is a mutex in play as part of the `crud.go` file. I found out quickly while spawning goroutines for handling input that concurrent writes and indexes of the hashmap caused a panic. Not surprising, but just wanted to note this, especially since it looks like the lock is missing from the QUERY input case. My research found that concurrent reads were fine, so it was not necessary there.

- With regards to testing, I tried to use some best practices as I found them in my research. An example of this is using "Table Driven Tests", basically a struct containing all test data for the given test. That can be seen throughout the various test cases. Additionally, the test for the handler in `main_test.go` is particularly interesting. While I've used the Go http server and http testing before, using connectors from the net package in this way was a new one for me. One of the recommendations I found from Mitchell Hashimoto (HashiCorp) regarding testing networking was: "If you're testing networking, make a real network connection. Don’t mock net.Conn, no point". I took this to heart and implemented a mock server/client with `net.Dial()`, which I think worked out really well. Just wanted to highlight one of the more interesting tests I got to work on.