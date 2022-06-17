FROM golang:1.16

ADD app /app
EXPOSE 8080
CMD ["/app"]