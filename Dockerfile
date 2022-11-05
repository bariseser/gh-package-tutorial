FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download && go build -o server ./main.go

FROM golang:1.19-alpine as runner
COPY --from=builder /app/server /app/server
EXPOSE 9999
CMD ["/app/server"]
