FROM golang:1.26.4 AS builder

WORKDIR /app 

COPY go.mod go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o todo-backend .

FROM alpine:3.22 

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/todo-backend .

USER appuser

CMD ["./todo-backend"]
