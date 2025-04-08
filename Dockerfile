FROM golang:1.24.1

COPY . .
RUN go build -o server .
CMD ["./server"]
