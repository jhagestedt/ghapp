FROM golang:alpine AS build

FROM scratch AS cli
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ARG TARGETOS="linux"
ARG TARGETARCH="amd64"
COPY ./dist/ghapp_${TARGETOS}_${TARGETARCH} /ghapp
ENTRYPOINT ["/ghapp"]
CMD ["help"]
