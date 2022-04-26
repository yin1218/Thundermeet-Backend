FROM golang:latest
RUN mkdir /usr/local/go/src/thundermeet-backend
WORKDIR /usr/local/go/src/thundermeet-backend
ADD . /usr/local/go/src/thundermeet-backend

RUN go mod download
RUN go build ./main.go
CMD ["./main"]

