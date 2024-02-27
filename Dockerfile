# ======== Notes ========
# fix submodule
# cargo build
# move bin to dockerfile
# add bash calls to 

# ======== Cairo VM ========
FROM rust:1.74 as build-stage
RUN apt-get update && apt-get install -y make

WORKDIR /app
COPY ./cairo-vm/ ./cairo-vm

WORKDIR /app/cairo-vm/cairo1-run
RUN make deps; exit 0
RUN make test

# ======== Lambda Works ========
WORKDIR /app
COPY ./lambdaworks/ ./lambdaworks

WORKDIR /app/lambdaworks/provers/cairo
RUN cargo install --features=cli,instruments,parallel --path .


# ======== Install Dependencies ========
FROM golang:1.21
RUN apt-get update && \
    curl --proto '=https' --tlsv1.2 -sSf https://docs.swmansion.com/scarb/install.sh | sh; exit 0

# not sure where to find this??
COPY --from=build-stage /app/cairo-vm/target/debug/cairo1-run /app/cairo-vm/cairo1-run
COPY --from=build-stage /app/cairo-vm/cairo1-run/corelib /app/cairo-vm/corelib
COPY --from=build-stage /usr/local/cargo/bin/platinum-prover /root/.local/bin

# ======== Go Server ========
WORKDIR /app
COPY ./cairo/ ./cairo
COPY ./server/ ./server

# ======== Misc Commands ========
ENV PATH="/root/.local/bin:${PATH}"
EXPOSE 8080

