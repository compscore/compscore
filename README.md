# compscore

Cyber Competition Scoring Platform

## Scorechecks

- [DNS](https://github.com/compscore/dns)
- [FTP](https://github.com/compscore/ftp)
- [LDAP](https://github.com/compscore/ldap)
- [MySQL](https://github.com/compscore/mysql)
- [Ping](https://github.com/compscore/ping)
- [SMB](https://github.com/compscore/smb)
- [SSH](https://github.com/compscore/ssh)
- [Web](https://github.com/compscore/web)
- [WinRM](https://github.com/compscore/winrm)

**Need something else?**
Build you own with [Check Template](https://github.com/compscore/check-template)

## Installation

### Initial

```sh
git clone https://github.com/compscore/compscore.git
cd compscore
```

### Starting Local/Development Instance

```sh
# Generate neccesary code base on your configuration
make generate

# Install `compscore` binary onto your system to use
make install
```

## Command Line Usage

|   subcommands   | description                                                                             |
| :-------------: | :-------------------------------------------------------------------------------------- |
|   `generate`    | `generate all nessesary code based on configured checks`                                |
|    `server`     | `start compscore server interactively`                                                  |
|    `version`    | `provide all current verstion information of compscore install`                         |
|    `engine`     | `interact with scoring engine`                                                          |
| `engine status` | `get current status of scoring engine`                                                  |
| `engine start`  | `spawn a daemon verions of server; if already exists, will start scoring on the engine` |
| `engine pause`  | `pause scoring engine once round round complete`                                        |
|  `engine kill`  | `kill scoring engine process safely`                                                    |
