FROM golang:1.16.0-alpine3.13
ENV MWAPI ""
ENV MWROOT ""
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]