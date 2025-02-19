FROM golang:1.24.0-alpine

WORKDIR /app

# Copy and download deps first to cache them
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/app

CMD ["./main"]