FROM golang:alpine AS build

ARG mentor_env

# copy directory to build container(temp)
RUN mkdir -p /usr/local/go/src/mentor-server/classroom
WORKDIR /usr/local/go/src/mentor-server/classroom
ADD . /usr/local/go/src/mentor-server/classroom
ENV mentor_env=$mentor_env

# build go app
RUN go mod download
RUN go build ./main.go

# use alpine as actual base image
FROM alpine
ARG mentor_env
WORKDIR /app
# copy app from build container
COPY --from=build /usr/local/go/src/mentor-server/classroom /app
ENV mentor_env=$mentor_env

EXPOSE 2304
CMD ["./main"]