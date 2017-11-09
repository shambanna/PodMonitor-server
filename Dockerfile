#FROM golang:latest
#RUN mkdir /app
#ADD . /app/
#WORKDIR /app
#RUN export GIT_TERMINAL_PROMPT=1
#RUN go get
#RUN go build -o main .
#CMD ["/app/main"]
#FROM golang:onbuild
#EXPOSE 8080
# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
#ADD . /go/src/github.com/golang/example/outyet
ADD . /go/src/github.com/PodMonitor-server
RUN go get -v -u github.com/gorilla/mux
RUN go get gopkg.in/yaml.v2
# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
#RUN go install github.com/golang/example/outyet
RUN go install github.com/PodMonitor-server
# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/PodMonitor-server

# Document that the service listens on port 8080.
EXPOSE 8080