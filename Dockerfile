# Start from the latest golang base image
FROM golang:1.17 as builder

ENV HOME /app
ENV CGO_ENABLED 0
ENV GOOS linux
ARG PAT_USERNAME
ARG GH_REPO_ACCESS_TOKEN

RUN echo patuser:  ${#PAT_USERNAME}
RUN echo ghToken:  ${#GH_REPO_ACCESS_TOKEN}

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN git config --global url."https://"${PAT_USERNAME}":"${GH_REPO_ACCESS_TOKEN}"@github.com/".insteadOf "https://github.com/"

# setting private repo
RUN go env -w GOPRIVATE="github.com/Maersk-Global/*"

# Test
# RUN go get github.com/Maersk-Global/service-composition-gsdk@v0.2.0

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .
# =================================================== #

FROM alpine:latest

LABEL maintainer="Purple Sea <purplesea.team@maersk.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/config ./config

# Expose port 8080 to the outside world
EXPOSE 8000
EXPOSE 9000

# Command to run the executable
CMD ["./main"]