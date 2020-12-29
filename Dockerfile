FROM golang:1.15
WORKDIR /go/src/github.com/rwxproject/go-web-hostname
RUN go get -d -v github.com/gorilla/mux
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
WORKDIR /app/
COPY --from=0 /go/src/github.com/rwxproject/go-web-hostname/app .
CMD ["./app"]  