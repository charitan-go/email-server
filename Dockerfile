# FROM golang:1.23-alpine

FROM golang:1.23-bookworm

# RUN apt-get install 

COPY . /app 

WORKDIR /app 

RUN go mod tidy 

# EXPOSE 50051

ENTRYPOINT ["go", "run", "./cmd"]
