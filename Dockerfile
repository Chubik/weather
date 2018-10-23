FROM golang:1.11.1-alpine3.8
COPY . /go/src/weather
WORKDIR /go/src/weather

RUN go build -o /app/weather .
CMD ["/app/weather"]
EXPOSE 8089
