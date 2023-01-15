# Feserve

## Overview

Feserve is a lightweight application created to make it easier for Frontend Developers to deploy their applications, without having to use Nginx, Node.js or the like which take up a lot of storage space.

## Installation

First download the application that suits your OS and architecture [here](https://github.com/ryanbekhen/feserve/releases). Then export to the path you want.

## Setup

Create a file in the root of the project with the name app.yaml with the following contents:

```yaml
version: 1
port: 8000
publicDir: build
```

With the configuration above the application will run on port 8000 with `build` as the directory that will be presented to the public.

## Usage

Run with the command below:

```shell
./feserve
```
