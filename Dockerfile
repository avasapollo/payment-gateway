FROM alpine:latest

ADD payment-gateway /usr/local/bin/payment-gateway

CMD ["/usr/local/bin/payment-gateway"]