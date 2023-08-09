FROM golang:alpine AS build

ARG mentor_env

# copy directory to build container(temp)
RUN mkdir -p /usr/local/go/src/mentor-server/chatroom
WORKDIR /usr/local/go/src/mentor-server/chatroom
ADD . /usr/local/go/src/mentor-server/chatroom
ENV mentor_env=$mentor_env

# build go app
RUN go mod download
RUN go build ./main.go

# use alpine as actual base image
FROM alpine
ARG mentor_env
WORKDIR /app
# copy app from build container
COPY --from=build /usr/local/go/src/mentor-server/chatroom /app
ENV mentor_env=$mentor_env
RUN apk update && apk add tzdata
RUN export GOOGLE_APPLICATION_CREDENTIALS="/app/key/mentor-app-gcp-service-key.json"

EXPOSE 2303
CMD ["./main"]