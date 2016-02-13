FROM debian:last

COPY * /

EXPOSE 443
ENTRYPOINT ["/Docker-WebManager"]

#docker run -d -p 5050:443 --privileged -v /var/run/docker.sock:/var/run/docker.sock dockerweb