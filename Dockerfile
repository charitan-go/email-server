# FROM golang:1.23-alpine

FROM golang:1.23-bookworm

RUN apt-get update -y && apt-get install -y iputils-ping

COPY . /app 

WORKDIR /app 

RUN go mod tidy 

# EXPOSE 50051

ENTRYPOINT ["go", "run", "./cmd"]
