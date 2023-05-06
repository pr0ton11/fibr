###
### Build Container
###

FROM golang:1.20-bullseye as build
COPY . /build
WORKDIR /build
RUN go build

###
### Runtime Container
###

FROM debian:bullseye-slim

# Configuration
VOLUME [ "/etc/fibr.yaml" ]

# Models
VOLUME [ "/models" ]

COPY --from=build /build/fibr /usr/bin/fibr
