FROM alpine:3.16
COPY ./artifacts/linux/word-cloud-generator /opt
EXPOSE 8888
CMD /opt/word-cloud-generator
