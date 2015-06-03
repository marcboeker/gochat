# Gochat

This is a simple application to demonstrate how easy it is with Go to build a networking application.

## Installation

We are using [gb](http://getgb.io/) as build tool.

    git clone git@github.com:marcboeker/gochat.git
    cd gochat
    gb build

## Usage

To start the server:

    ./gochat -m server

To connect a client with the username "kermit":

    ./gochat -m client -u kermit

Once the client is connected, enter some text and press enter.
