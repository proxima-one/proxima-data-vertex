FROM golang:alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .
EXPOSE 4000
RUN chmod
CMD ["sudo ./main"]
