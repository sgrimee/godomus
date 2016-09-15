# ldclient - Command line client for Lifedomus home automation server

## Features
- get list of sites, users, rooms or devices in a room
- get available scenarios

## Configuration
Parameters like server url, user, password can be given on the command line or saved in a configuration file. By default this is $HOME/.ldclient.yaml
See [ldclient_template.yaml](ldclient_template.yaml) for an example you can adapt and rename.

## Installation

## Usage

```
ldclient get sites
ldclient get users
ldclient get scenarios
ldclient get rooms
ldclient get devices --room 5
```

You can also change the output format between yaml, json and text (default). For json it is interesting to pipe into the great [jq](https://stedolan.github.io/jq/) utility 

```
ldclient get devices --room 5 -o json | jq
```

## Todo
- launch scenario
- turn device on/off
- listen to events from the server

## Disclaimer

This is community work not provided by nor affiliated with Lifedomus or Delta Dore in any way. Please do not contact them with issues regarding this software. See LICENSE file.
