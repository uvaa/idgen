FROM golang:alpine3.12 AS build
ENV GOPROXY=https://goproxy.io,direct
WORKDIR /app
COPY . .
RUN go build

FROM alpine:3.12
ENV GIN_MODE=release
WORKDIR /app
EXPOSE 80
COPY --from=build /app/idgen .
CMD ["./idgen"]