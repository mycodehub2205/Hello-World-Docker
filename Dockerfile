FROM alpine

WORKDIR /

ADD ./main .

CMD [ "./main" ]