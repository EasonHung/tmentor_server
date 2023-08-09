FROM golang:alpine AS build

ARG mentor_env

# copy directory to build container(temp)
RUN mkdir -p /usr/local/go/src/mentor-server/case_system
WORKDIR /usr/local/go/src/mentor-server/case_system
ADD . /usr/local/go/src/mentor-server/case_system
ENV mentor_env=$mentor_env

# build go app
RUN go mod download
RUN go build ./main.go

# use alpine as actual base image
FROM alpine
ARG mentor_env
WORKDIR /app
# copy app from build container
COPY --from=build /usr/local/go/src/mentor-server/case_system /app
ENV mentor_env=$mentor_env
ENV GOOGLE_APPLICATION_CREDENTIALS="/app/key/mentor-app-gcp-service-key.json"

EXPOSE 2305
CMD ["./main"]