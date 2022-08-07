FROM golang:1.18-bullseye AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o bulkupdates main.go 

FROM alpine:latest AS runtime
WORKDIR /app
COPY --from=build /src/bulkupdates /app/
EXPOSE 8086
CMD ["./bulkupdates"]