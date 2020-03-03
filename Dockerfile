FROM golang

RUN go get github.com/lib/pq
RUN go get github.com/Nhanderu/brdoc
RUN go get github.com/tkanos/gonfig