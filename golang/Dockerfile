FROM golang:1.9 as builder
WORKDIR /go/src/hello/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello .

FROM scratch
COPY --from=builder /go/src/hello/hello .
EXPOSE 80
CMD ["./hello"] 
