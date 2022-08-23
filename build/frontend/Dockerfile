FROM ubuntu:22.04
SHELL ["/bin/bash", "-c"]
WORKDIR /space

COPY ./dist ./dist/
COPY ./file-server .

# Example usage
# ./file-server 0.0.0.0:8001 /space/work
# ./file-server 0.0.0#8001

CMD ["./file-server", "0.0.0.0:8001", "/space/dist/"]
