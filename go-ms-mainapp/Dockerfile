# Gunakan base image golang di alpine
FROM golang:1.19-alpine

# Install dependencies yang dibutuhkan (git)
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy semua file
COPY . .

# Download dependencies dan build aplikasi
RUN go mod tidy
RUN go build -o go-ms-mainapp

# Expose port untuk aplikasi (sesuaikan dengan port aplikasi Go Fiber)
EXPOSE 5000

# Jalankan aplikasi
CMD ["./go-ms-mainapp"]
