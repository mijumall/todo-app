FROM ubuntu:22.04
SHELL ["/bin/bash", "-c"]

ENV cockroach_installer cockroach-v22.1.4.linux-amd64.tgz
ENV workspace /space

WORKDIR $workspace

RUN apt update && apt full-upgrade -y && apt install -y less vim curl 

COPY ./$cockroach_installer .
RUN tar -xz -f $cockroach_installer && rm $cockroach_installer
RUN mv cockroach* cockroachdb

ENV PATH $PATH:$workspace/cockroachdb

WORKDIR $workspace/db/

CMD ["cockroach", "--help"]

