FROM gliderlabs/alpine:3.2

ADD ./check/check /opt/resource/check
ADD ./in/in /opt/resource/in
ADD ./out/out /opt/resource/out

RUN chmod +x /opt/resource/*

