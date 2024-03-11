FROM golang:alpine as builder
LABEL authors="Yoake"
WORKDIR /app/tongyiqwen
RUN apk --no-cache add ca-certificates  && update-ca-certificates
RUN apk --no-cache add tzdata
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy -v
RUN go build -o main

FROM scratch as runtime
LABEL authors="Yoake"
ENV TZ=UTC+8
WORKDIR /app/tongyiqwen
COPY --from=builder /app/tongyiqwen/main main
COPY --from=builder /app/tongyiqwen/config.example.yaml config.yaml
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
EXPOSE 20104
CMD ["./main"]