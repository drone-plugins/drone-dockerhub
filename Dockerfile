# Docker image for Drone's DockerHub plugin
#
#     CGO_ENABLED=0 go build -a -tags netgo
#     docker build --rm=true -t plugins/drone-dockerhub .

FROM gliderlabs/alpine:3.1
RUN apk-install ca-certificates
ADD drone-slack /bin/
ENTRYPOINT ["/bin/drone-dockerhub"]
