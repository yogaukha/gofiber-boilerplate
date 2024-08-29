FROM golang:1.20-alpine3.18

# for hot reload using AIR
RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

# RUN go build -o /rice-go

# EXPOSE 3000

ENTRYPOINT [ "air" ]
