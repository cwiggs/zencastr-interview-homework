From golang:1.13-alpine

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o server .

CMD [ "/app/server" ]
