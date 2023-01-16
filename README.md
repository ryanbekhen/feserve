# Feserve

## Overview

Feserve is a lightweight application created to make it easier for Frontend Developers to deploy their applications, without having to use Nginx, Node.js or the like which take up a lot of storage space.

## Installation

### Binary File

Here I use linux with amd64 architecture as an example. Please adjust to your OS and Architecture [here](https://github.com/ryanbekhen/feserve/releases). Then download, verify the signature, and extract it like the following example.

```shell
wget https://github.com/ryanbekhen/feserve/releases/download/v0.1.0/feserve_0.1.0_linux_amd64.zip
wget https://github.com/ryanbekhen/feserve/releases/download/v0.1.0/checksums.txt
unzip feserve_0.1.0_linux_amd64.zip 
sha256sum --ignore-missing -c checksums.txt
```

After running the above command, move the binary file to `/usr/local/bin` with the following command.

```shell
sudo mv feserve /usr/local/bin
```

### Via `go install`

```shell
go install github.com/ryanbekhen/feserve
```
> **Note**: go version go1.19.5 or later

## Setup

### Directory Structure

```text
root-directory/
|- build/
|- app.yaml
```

### Configuration `app.yaml`

```yaml
version: 1
port: 8000
publicDir: build
```

With the above configuration, feserve will run on port `8000` and `public/` as its public directory.

## Usage

### Local

To run it locally, just run `feserve` in your root directory with the following command.

```shell
feserve
```
Then open a browser at http://localhost:8000.

### Docker

To run it within docker, create a file `Dockerfile` like the following example.

```Dockerfile
# application build
FROM node:16-alpine As build
WORKDIR /app
COPY . .
RUN npm ci 
RUN npm run build
ENV NODE_ENV production

# application serve
FROM ghcr.io/ryanbekhen/feserve:0.1.0
WORKDIR /app
COPY app.yaml .
COPY --from=build /app/build /app/build
EXPOSE 8000
ENTRYPOINT ["feserve"]
```

It can also be done in the following way if we have built it first.

```Dockerfile
FROM ghcr.io/ryanbekhen/feserve:0.1.0
WORKDIR /app
COPY app.yaml .
COPY build ./build
EXPOSE 8000
ENTRYPOINT ["feserve"]
```

To try to run it simply with the following command.

```shell
docker build -t image-name .
docker run --rm -p 8000:8000 image-name
```

Then open a browser at http://localhost:8000.

> **Note**: If you get an `unauthorized` error, it's because the image is only in the Github registry. Please run the command `docker login ghcr.io -u username -p github_personal_access_token` with scope `read:packages` then run it again.

## Security
If you discover a security vulnerability within Feserve, please send an e-mail to ryanbekhen.official@gmail.com.

## License
This program is free software: you can redistribute it and/or modify it under the terms of the Apache license. Feserve and any contributions are copyright © by Achmad Irianto Eka Putra 2023