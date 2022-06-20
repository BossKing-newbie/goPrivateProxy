FROM golang:1.16
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
ADD app /app
EXPOSE 8138
RUN cat hosts >> /etc/hosts
CMD ["/app"]