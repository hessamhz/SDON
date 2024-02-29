############################
# STEP 1 build executable binary
############################
FROM golang:1.21-alpine AS builder

# Install git.
# Git is required for fetching the dependencies.

RUN apk update && apk add --no-cache git ca-certificates bash

# Create appuser.
ENV USER=napcore_user
ENV UID=10001 


# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR $GOPATH/napcore

COPY . .
# Fetch dependencies.
# Using go mod.

RUN go mod download
RUN go mod verify

# Build the binary.

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -ldflags '-extldflags "-static"'

COPY ./.env $GOPATH/napcore/
COPY ./update_cookies.sh $GOPATH/napcore/
RUN chmod +x $GOPATH/napcore/update_cookies.sh

COPY ./.env /go/bin/.env
COPY ./update_cookies.sh /go/bin/update_cookies.sh
RUN chmod +x /go/bin/update_cookies.sh

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o /go/bin/napcore .

RUN which ls
RUN which sh
RUN ls /go/bin
RUN ls -l /go/bin

############################
# STEP 2 build a small image
############################
FROM alpine

# Install ca-certificates and bash if needed
RUN apk update && apk add --no-cache ca-certificates bash

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable.
COPY --from=builder /go/bin/napcore /napcore

# Copy the .env file and script if they need to be in the image
COPY --from=builder /go/bin/.env /.env
COPY --from=builder /go/bin/update_cookies.sh /update_cookies.sh
RUN chmod +x /go/bin/update_cookies.sh

COPY --from=builder /go/bin/.env /go/bin/.env
COPY --from=builder /go/bin/update_cookies.sh /go/bin/update_cookies.sh
RUN chmod +x /go/bin/update_cookies.sh

ENV PATH="/go/bin:${PATH}"

# Use an unprivileged user.
USER napcore_user:napcore_user

# Run the binary.
ENTRYPOINT ["/napcore"]
