FROM golang:1.18-bullseye AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o products main.go 

FROM alpine:latest AS runtime
WORKDIR /app
COPY --from=build /src/products /app/
EXPOSE 8085
CMD ["./products"]