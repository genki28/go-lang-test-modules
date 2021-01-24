# 2020/10/14最新versionを取得
FROM golang:1.15.2-alpine
# アップデートとgitのインストール！！
RUN apk update && apk add git
# appディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app
# ENV GO111MODULE=on

COPY go.mod ./

RUN go mod download

COPY . .

RUN go get -u github.com/oxequa/realize
CMD ["realize", "start", "--run"]