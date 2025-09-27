# Gunakan base image untuk Golang
FROM golang:1.24-alpine

# Set working directory di dalam container
WORKDIR /app

# Copy file environment
COPY .env .

# Copy semua file source ke dalam container
COPY . .

# Unduh dependencies
RUN go mod tidy

# Build aplikasi dan beri nama binary sesuai yang akan dijalankan
RUN go build -o golang-eis

# Jalankan aplikasi
CMD ["./golang-eis"]
