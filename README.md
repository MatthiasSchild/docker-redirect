# docker-redirect

- [Docker Hub](https://hub.docker.com/r/matthiasschild/redirect)

This repository contains the code for my redirect docker image.
It uses the HTTP Code 303 to redirect and appends the path to the redirected URL.

## Environment variables

- REDIRECT: The address, to which the visitor should be redirected
- ADDR: The web address for the server (default is `:8080`)

## Docker run example

    docker run --name redirect -p 8080:8080 -e REDIRECT https://example.org/ matthiasschild/redirect

## Docker compose example

    version: "3"

    services:
      redirect:
        image: matthiasschild/redirect
        ports:
          - 8080:8080
        environment:
          - REDIRECT=https://example.org/

## Docker compose example with traefik

    version: "3"

    services:
      redirect:
        image: matthiasschild/redirect
        environment:
          - REDIRECT=https://example.org/
        labels:
          - "traefik.enable=true"
          - "traefik.frontend.rule=Host:example.org"
          - "traefik.backend=redirect"
          - "traefik.port=8080"