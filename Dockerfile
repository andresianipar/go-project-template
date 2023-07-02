# Build the application from source
FROM golang:1.20-alpine as build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod tidy && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /nuke-deployer-system ./cmd

# Deploy the application binary into a lean image
FROM alpine:latest AS build-release-stage

RUN apk -U upgrade && apk add --no-cache bash

WORKDIR /app

COPY --from=build-stage /nuke-deployer-system ./

EXPOSE 52603
CMD [ "./nuke-deployer-system" ]
