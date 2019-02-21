FROM golang:alpine as builder
RUN apk update && apk add git && go get gopkg.in/natefinch/lumberjack.v2

WORKDIR /weather
ENV CGO_ENABLED=0
COPY . .
RUN go build -o weather .

FROM  alpine as runner

WORKDIR /weather
COPY --from=builder weather/ ./

CMD ["/weather/weather"]
EXPOSE 3001
EXPOSE 3000 
