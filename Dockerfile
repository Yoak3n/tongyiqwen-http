FROM golang:alpine as builder
LABEL authors="Yoake"
WORKDIR /app/tongyiqwen
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy -v
RUN go build -o main

FROM scratch as runtime
LABEL authors="Yoake"
WORKDIR /app/tongyiqwen
COPY --from=builder /app/tongyiqwen/main main
COPY --from=builder /app/tongyiqwen/config.example.yaml config.yaml
CMD ["./main"]