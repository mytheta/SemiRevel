FROM golang:latest

LABEL maintainer = "Yutsuki Miyashita <j148015n@st.u-gakugei.ac.jp>"
LABEL description = "ゼミ資料管理のAPI for Hazelab"

ENV GOPATH /go
RUN mkdir -p ${GOPATH}/src

ENV SEMIREVELDIR ${GOPATH}/src/SemiRevel
RUN mkdir ${SEMIREVELDIR}


# COPY . ${SEMIREVELDIR}
WORKDIR ${SEMIREVELDIR}
# Go dep!
RUN go get -u github.com/golang/dep/...
RUN dep ensure

# Revel Run
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel
CMD ${SEMIREVELDIR}/SemiRevel
