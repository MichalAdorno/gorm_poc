FROM golang:1.14.1-alpine3.11 AS builder
RUN mkdir /source
COPY ["./","/source/"]
RUN cd /source && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o main .  
FROM alpine:3.11.5 AS production
COPY --from=builder ["/source/main","/"]
RUN chmod a+x /main && \
    apk --no-cache add curl
EXPOSE 8080
ENTRYPOINT ["/main"]