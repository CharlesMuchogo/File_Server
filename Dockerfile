FROM golang:latest
RUN mkdir /build
WORKDIR /build
RUN git clone https://github.com/CharlesMuchogo/File_Server.git
WORKDIR /build/File_Server
RUN go build -o main
EXPOSE 8080
ENTRYPOINT [ "/build/File_Server/main" ]docker images