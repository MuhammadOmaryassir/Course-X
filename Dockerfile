FROM golang:latest
RUN mkdir -p /app
ADD . /app
WORKDIR /app
RUN go install
RUN go build ./server.go
CMD ["./server"]