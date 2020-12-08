ARG GOLANG_VERSION=1.15

# Use base golang image from Docker Hub
FROM golang:$GOLANG_VERSION AS build

WORKDIR /go/src/github.com/micromaniacs/gcloud-grpc-boilerplate

# Install dependencies in go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the application source code
COPY . ./

# Compile the application to /srv and minimize binary size.
# Skaffold passes in debug-oriented compiler flags
ARG SKAFFOLD_GO_GCFLAGS
RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
        upx \
    && echo "Go gcflags: ${SKAFFOLD_GO_GCFLAGS}"
    && go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -mod=readonly -v -o /srv cmd/srv/main.go \
    && upx -9 /srv

# Now create separate deployment image
FROM gcr.io/distroless/base

# Definition of this variable is used by 'skaffold debug' to identify a golang binary.
# Default behavior - a failure prints a stack trace for the current goroutine.
# See https://golang.org/pkg/runtime/
ENV GOTRACEBACK=single

# Copy template & assets
WORKDIR /srv
COPY --from=build /srv ./srv

ENTRYPOINT ["./srv"]
