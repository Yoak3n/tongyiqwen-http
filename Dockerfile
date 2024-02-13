FROM golang:alpine
LABEL authors="Yoake"
WORKDIR /app/tongyiqwen
COPY . .
RUN go env -w  GOPROXY=https://goproxy.cn,direct
RUN go mod tidy -v
RUN go build -o main
CMD ["./main"]
