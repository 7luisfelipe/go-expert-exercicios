FROM golang:1.23 as build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goapi ./main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/goapi .
EXPOSE 8081
ENTRYPOINT [ "./goapi" ]