# Build Container
FROM golang:alpine AS build

WORKDIR /app

COPY ./src/NoxFollow/ ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /follow

# Release Container
FROM scratch 

WORKDIR /

COPY --from=build /follow /follow

EXPOSE 8080

ENTRYPOINT ["/follow"]