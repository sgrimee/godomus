# godomus
GO client library for the Lifedomus home automation server

## Applications
- [ldclient - CLI utility](https://github.com/sgrimee/godomus/tree/master/ldclient)
- Soon: [hkdomus - HomeKit bridge](https://github.com/sgrimee/hkdomus)

## Features
- Login
- Get sites, users, rooms, scenarios, devices in a room
- Listen to events from the server

## TODO
- Handle login failures
- Thread-safe the websocket server events piece
- Launch scenarios
- Set device state

## Testing

Running tests require a credentials to a working server. The same config file as the cli client can be used with a few additional entries. Please check [this file](test_config_template.yaml) for a template and copy it to $HOME/.ldclient.yaml or to .ldclient.yaml in this folder (for tests only).

## Disclaimer

This is community work not provided by nor affiliated with Lifedomus or Delta Dore in any way. Please do not contact them with issues regarding this software. See LICENSE file.