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
RUN make deps
RUN make test

# ======== Lambda Works ========
WORKDIR /app
COPY ./lambdaworks/ ./lambdaworks

WORKDIR /app/lambdaworks/provers/cairo
RUN cargo install --features=cli,instruments,parallel --path .
# COPY /h.cargo/bin/platinum-prover /root/.cargo/bin/platinum-prover


# ======== Install Dependencies ========
# FROM golang:1.21
# RUN apt-get update 
# # && \
#     # curl --proto '=https' --tlsv1.2 -sSf https://docs.swmansion.com/scarb/install.sh | sh; exit 0
# COPY --from=build-stage /app/cairo-vm/cairo1-run/target/debug/cairo1-run /app/cairo-vm/cairo1-run/target/debug/cairo1-run
# COPY --from=build-stage /usr/local/cargo/bin/platinum-prover /usr/local/cargo/bin/platinum-prover

# ======== Go Server ========
WORKDIR /app
COPY ./cairo/ ./cairo
COPY ./server/ ./server

# COPY ./cairo-vm/ ./cairo-vm
# WORKDIR /app/cairo-vm/cairo1-run
# RUN make deps
# RUN make test
# COPY ../target/debug/cairo1-run /cairo-vm/cairo1-run


# ======== Misc Commands ========
ENV PATH="/root/.local/bin:${PATH}"
EXPOSE 8080

# ENTRYPOINT [ "/bin/bash", "-l", "-c" ]
# CMD ["/server"]
# ENTRYPOINT [ "/bin/bash", "-l", "-c" ]
