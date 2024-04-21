# Simple API Server using Golang

This is a simple http server with an api endpoint that will return data from a sqlite database. The server will be run as a service using systemd.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Uninstallation](#uninstallation)

## Features

- Provides a basic HTTP API endpoint to get memes of Pepe the frog.
- Gets data from a sqlite database.
- Reads environment variables from a .env file.
- Can be installed as a systemd service on Ubuntu/Debian.

## Prerequisites

- Go (>= 1.14) installed.
- systemd
- Open port at the firewall

## Installation

```bash
git clone https://github.com/ismaventuras/simple-api-server-using-golang
cd simple-api-server-using-golgang
```

Then create a `.env` file and fill the variables from `example.env`. Remember to allow the port on the firewall. 

After this, run
```bash
sudo make install
```

## Usage

Once installed, you can access the API server using the configured port you've set. Visit http://localhost:{PORT} in your web browser or use a tool like curl to interact with the API endpoints.

## Uninstallation

```bash
cd simple-api-server-using-golang
sudo make uninstall
```