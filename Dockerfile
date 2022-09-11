FROM scratch AS cli
ARG TARGETOS="linux"
ARG TARGETARCH="amd64"
COPY ./dist/ghapp_${TARGETOS}_${TARGETARCH} /ghapp
ENTRYPOINT ["/ghapp"]
CMD ["help"]
