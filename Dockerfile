FROM golang:1.17-alpine as build
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY go.mod .
COPY *.go .

RUN GOOS=linux GOARCH=amd64 go build -o /httpserver

# second phase
FROM alpine
COPY --from=build /httpserver /
RUN chmod a+x /httpserver
EXPOSE 80
# ADD httpserver /httpserver
CMD [ "/httpserver" ]