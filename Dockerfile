FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux go build -o /weather-next cmd/api.go
RUN chmod +x /weather-next

CMD ["/weather-next"]
