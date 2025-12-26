FROM golang:alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build .
EXPOSE 9178
CMD ["./text2lang"]
