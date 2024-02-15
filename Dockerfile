# image is hella fat, optimize by compiling go code
# cairo needed at runtime so keep scarb
FROM golang:1.21

# Install Scarb and thus Cairo
# exit 0 because downloading scarb can't detect shell
RUN apt-get update && \
    curl --proto '=https' --tlsv1.2 -sSf https://docs.swmansion.com/scarb/install.sh | sh; exit 0

WORKDIR /app
COPY ./cairo/ ./cairo
COPY ./server/ ./server

# Set any necessary environment variables here
# Add /root/.local/bin to PATH (for scarb)
ENV PATH="/root/.local/bin:${PATH}"

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the binary.
# CMD ["/server"]
