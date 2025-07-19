FROM golang:1.24

WORKDIR /app

COPY salesman/go.mod ./

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/air-verse/air@latest

RUN go mod download && go mod verify

COPY salesman/ .

CMD ["air", "-c", ".air.toml"]
