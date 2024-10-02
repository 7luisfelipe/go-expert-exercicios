FROM golang:1.23 as build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gocli .

FROM scratch
WORKDIR /app
COPY --from=build /app/gocli .
EXPOSE 8081
ENTRYPOINT [ "./gocli" ]