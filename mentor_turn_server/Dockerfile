FROM golang:alpine AS build


# copy directory to build container(temp)
RUN mkdir -p /usr/local/go/src/turnserver
WORKDIR /usr/local/go/src/turnserver
ADD . /usr/local/go/src/turnserver

# build go app
RUN go mod download
RUN go build ./main.go

# use alpine as actual base image
FROM alpine
WORKDIR /app
# copy app from build container
COPY --from=build /usr/local/go/src/turnserver /app

EXPOSE 9000
CMD ["./main"]