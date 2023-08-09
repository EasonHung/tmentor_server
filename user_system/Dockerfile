FROM golang:alpine AS build

ARG mentor_env

# copy directory to build container(temp)
RUN mkdir -p /usr/local/go/src/mentor-server/user_system
WORKDIR /usr/local/go/src/mentor-server/user_system
ADD . /usr/local/go/src/mentor-server/user_system
ENV mentor_env=$mentor_env

# build go app
RUN go mod download
RUN go build ./main.go

# use alpine as actual base image
FROM alpine
ARG mentor_env
WORKDIR /app
# copy app from build container
COPY --from=build /usr/local/go/src/mentor-server/user_system /app
ENV mentor_env=$mentor_env
RUN apk update && apk add tzdata
ENV GOOGLE_APPLICATION_CREDENTIALS="/app/key/mentor-app-gcp-service-key.json"

EXPOSE 2306
CMD ["./main"]