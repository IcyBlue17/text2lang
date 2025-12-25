FROM golang:alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build .
EXPOSE 8080
CMD ["./text2lang"]
