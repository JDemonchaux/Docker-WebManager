FROM scratch

COPY * /

EXPOSE 443
ENTRYPOINT ["/Docker-WebManager"]

#docker run -d -p 443:5050 --privileged -v /var/run/docker.sock:/var/run/docker.sock dwm:latest