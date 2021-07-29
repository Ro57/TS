# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest alpine base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="PKT <cjd@cjdns.fr>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY ./bin ./bin

# Init .lnd folder
RUN mkdir /root/.lnd

RUN cd bin \
    && ./gencerts --host="*" --force \
    && mv rpc.key /root/.lnd/tls.key \
    && mv rpc.cert /root/.lnd/tls.cert 

# # Expose lnd ports (server, rpc).
# EXPOSE 9735 10009

# Command to run lightning network pkt node. Use shell form of CMD
# because pld doesn't work on PID 1 
CMD ./bin/pld --replication-server-addr=':8081'  --issuence-server-addr=':5050' --pkt.simnet