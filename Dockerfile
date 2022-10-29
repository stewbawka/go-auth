FROM golang:1.19

ENV GO111MODULE=on

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/stewbawka/go-auth

COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["go-auth"]
