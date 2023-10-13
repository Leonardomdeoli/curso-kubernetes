FROM golang:1.15
copy . .
RUN go build -o server .
CMD ["./server"]