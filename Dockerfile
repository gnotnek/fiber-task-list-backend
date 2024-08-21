FROM golang:1.22.0-alpine3.14 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /fiber-task-list-backend

FROM alpine:3.14

COPY --from=builder /fiber-task-list-backend /fiber-task-list-backend

EXPOSE 3000

CMD ["/fiber-task-list-backend"]