FROM golang:1.11.1
COPY . /go/src/weather
WORKDIR /go/src/weather

#RUN go get -u github.com/golang/dep/...
#RUN dep ensure

# RUN go run main.go
RUN go build -o /app/weather .
CMD ["/app/weather"]
EXPOSE 8089
