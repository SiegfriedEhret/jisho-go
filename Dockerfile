FROM scratch
MAINTAINER Siegfried Ehret <siegfried@ehret.me>
WORKDIR /
ADD jisho kanjidb.sqlite /
ENTRYPOINT ["./jisho"]
