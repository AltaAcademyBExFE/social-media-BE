FROM golang:1.19-alpine3.16

##buat folder APP
RUN mkdir /alta-sosmed

##set direktori utama
WORKDIR /alta-sosmed

##copy seluruh file ke completedep
ADD . .

##buat executeable
RUN go build -o main .

##jalankan executeable
CMD ["./main"]