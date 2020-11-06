FROM golang:1.15
RUN mkdir app/
ADD . /app
WORKDIR /app
RUN go get github.com/bwmarrin/discordgo
RUN go get github.com/mattn/go-sqlite3
RUN go get github.com/lus/dgc
RUN go build -o main .
CMD ["/app/main"]
