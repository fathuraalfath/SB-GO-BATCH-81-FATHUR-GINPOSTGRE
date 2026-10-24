# Build stage
FROM golang:alpine AS builder

WORKDIR /app

ENV GOTOOLCHAIN=auto

# Install git dan ca-certificates untuk dependensi
RUN apk add --no-cache git ca-certificates

# Copy go.mod dan go.sum terlebih dahulu untuk memanfaatkan cache
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary Go secara statis untuk Linux
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server .

# Final runtime stage (sangat ringan dan aman)
FROM alpine:latest

WORKDIR /app

# Install sertifikat CA agar koneksi SSL ke Supabase berjalan lancar
RUN apk add --no-cache ca-certificates tzdata

# Copy binary dari stage builder
COPY --from=builder /app/server .

# Expose port default
EXPOSE 8080

# Jalankan server
CMD ["./server"]
