FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hertz-api main.go
RUN chmod +x hertz-api

FROM scratch

WORKDIR /app

COPY --from=builder /app/hertz-api /app/
COPY --from=builder /app/config.toml /app/

EXPOSE 8888

CMD ["/app/hertz-api"]