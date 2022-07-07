FROM ubuntu:latest
ENV LANG en_US.utf8
RUN apt update
RUN apt install -y curl sudo net-tools vim telnet redis-server redis-sentinel iputils-ping

# CMD ["tail -F anything"]