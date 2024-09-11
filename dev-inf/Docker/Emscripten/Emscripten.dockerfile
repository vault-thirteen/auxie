# Stable Debian O.S. has an extremely outdated version of the 'cmake' tool.
FROM debian:trixie-slim

RUN apt-get -y update
RUN apt-get -y upgrade

RUN apt-get -y install git
RUN git version

RUN apt-get -y install python3
RUN python3 --version

RUN apt-get -y install cmake
RUN cmake --version

RUN apt-get -y install xz-utils
RUN xz --version

RUN apt-get -y install lsb-release

WORKDIR /home
RUN mkdir emscripten

WORKDIR /home/emscripten
RUN git clone https://github.com/emscripten-core/emsdk.git

WORKDIR /home/emscripten/emsdk
RUN git pull
RUN ./emsdk install latest
RUN ./emsdk activate latest
RUN ls

# Set PATH for Emscripten SDK.
# Uncomment this if you want to use the SDK inside the Dockerfile.
# The default shell for the RUN instructions is ["/bin/sh", "-c"].
# We need to temporarily override this to make the "source" command available.
#WORKDIR /home/emscripten/emsdk
#SHELL ["/bin/bash", "-c"] 
#RUN source ./emsdk_env.sh
#SHELL ["/bin/sh", "-c"]

WORKDIR /home
RUN mkdir data # <- This folder can be shared with a host O.S.

WORKDIR /home/data
