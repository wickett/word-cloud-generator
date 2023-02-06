FROM alpine:3.16
COPY ./artifacts/linux/word-cloud-generator /opt
EXPOSE 8888
ENTRYPOINT /opt/word-cloud-generator
