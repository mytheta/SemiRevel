FROM golang:latest

LABEL maintainer = "Yutsuki Miyashita <j148015n@st.u-gakugei.ac.jp>"
LABEL description = "ゼミ資料管理のAPI for Hazelab"

RUN curl https://glide.sh/get | sh

ENV GOPATH /go
RUN mkdir -p ${GOPATH}/src

ENV SEMIREVELDIR ${GOPATH}/src/SemiRevel
RUN mkdir ${SEMIREVELDIR}


COPY . ${SEMIREVELDIR}
WORKDIR ${SEMIREVELDIR}
RUN glide install
RUN go build

CMD ${SEMIREVELDIR}/SemiGo
