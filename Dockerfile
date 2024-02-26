# ======== Notes ========
# fix submodule
# cargo build
# move bin to dockerfile
# add bash calls to 

# ======== Cairo VM ========
FROM rust:1.67 as build-stage
RUN apt-get update && \
    apt-get install -y make
WORKDIR /app
COPY ./cairo-vm/ ./cairo-vm
WORKDIR /app/cairo-vm/cairo1-run
RUN make deps
RUN make test

# ======== Install Dependencies ========
FROM golang:1.21
RUN apt-get update && \
    curl --proto '=https' --tlsv1.2 -sSf https://docs.swmansion.com/scarb/install.sh | sh; exit 0
COPY --from=build-stage /app/cairo-vm/target/debug/cairo1-run /app/cairo-vm/target/debug/cairo1-run

# ======== Go Server ========
WORKDIR /app
COPY ./cairo/ ./cairo
COPY ./server/ ./server

COPY ./cairo-vm/ ./cairo-vm
WORKDIR /app/cairo-vm/cairo1-run
RUN make deps
# RUN make test
# COPY ../target/debug/cairo1-run /cairo-vm/cairo1-run

# ======== Lambda Works ========
# COPY ./lambdaworks.bin/ ./lambdaworks.bin

# ======== Misc Commands ========
ENV PATH="/root/.local/bin:${PATH}"
EXPOSE 8080

ENTRYPOINT [ "/bin/bash", "-l", "-c" ]
# CMD ["/server"]
# ENTRYPOINT [ "/bin/bash", "-l", "-c" ]
