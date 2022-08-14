FROM ubuntu:22.04
SHELL ["/bin/bash", "-c"]
WORKDIR /space

ENV installer https://github.com/caddyserver/caddy/releases/download/v2.5.2/caddy_2.5.2_linux_amd64.deb

RUN apt update && apt full-upgrade -y && apt install -y curl
RUN curl -L $installer -o installer.deb 
RUN apt install -y ./installer.deb && rm installer.deb

COPY ./dist ./dist/
COPY ./Caddyfile .

CMD ["caddy"]
