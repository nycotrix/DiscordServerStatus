# Server Status
Monitors a list of servers and sends a chat notification when a server goes on or offline.

## Features

- send channel notifications
- track server up/down time
- **TCP** - should work with all servers
- **UDP** - [Source RCON Protocol](https://developer.valvesoftware.com/wiki/Source_RCON_Protocol) is supported
- [Docker](https://hub.docker.com/r/mgerb/server-status)

## Configuration
- Download the latest release [here](https://github.com/mgerb/ServerStatus/releases)
- Add your bot token as well as other configurations to **config.json**
- Execute the OS specific binary!

### Mentioning Roles/Users
- list of user/role ID's must be in the following format (see below for obtaining ID's)
- `<@userid>`
- `<@&roleid>`

### Polling Interval
The polling interval is how often the bot will try to ping the servers.
A good interval is 10 seconds, but this may need some adjustment if
it happens to be spamming notifications.

- time in seconds
- configurable in **config.json**
```

## Usage
To get the current status of your servers simply type `!ServerStatus` in chat.

![Server Status](./readme_files/screenshot1.png)

## Compiling from source
- Make sure Go and Make are installed
- make all
