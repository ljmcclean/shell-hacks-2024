FROM golang:1.23.1 as builder

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o serve .

FROM golang:1.23.1

WORKDIR /server

COPY --from=builder /src/serve .

EXPOSE 8080

ENTRYPOINT ["./serve"]
