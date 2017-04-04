FROM debian:jessie-slim

RUN apt-get update && apt-get install -y ca-certificates

# Copy the local package files to the container's workspace.
RUN mkdir -p /web_crawler
COPY ./deploy/example1 /web_crawler/example1
COPY ./deploy/example2 /web_crawler/example2
COPY ./deploy/example3 /web_crawler/example3

WORKDIR /web_crawler

# Run the outyet command by default when the container starts.
#ENTRYPOINT /web_crawler/example1
