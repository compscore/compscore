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

## Configuration

Configuring Compscore is super easy! There are to two files you will need to edit [config.yml](./config.yml), which is used for configuring Compscore itself, and [.env](./.env), which configured how compscore is deployed.

### config.yml

`config.yml` is used to configure Compscore itself, there are a couple sections you will need to edit:

#### name

This section does not do anything, it just so you can name your configuration.

#### users

Use the following format to define new user:

```yaml
users:
    - username: username_1
      password: password_1
    - username: username_1
      password: password_1
    ...
```

#### teams

Use the following format for team creation:

```yaml
teams:
  # amount of teams to create
  amount: 15

  # Name of users to create for compeitors
  # This example use `Team XX` format (`Team 01` - `Team 15`)
  nameFormat: Team { .Team }

  # Default password of all competition teams
  password: changeme123!
```

#### scoring

Use the following format:

```yaml
scoring:
  # length of scoring rounds in seconds
  interval: 30
```

#### engine

This section defines engine configuration, more than likely you will never have to edit this.

```yaml
engine:
  # file location of unix socket for interacting with compscore
  socket: /tmp/compscore.sock

  # grpc timeout in seconds for server running over unix socket
  timeout: 5
```

#### checks

This section is for defining all checks to be ran in Compscore, it is a list of configurations as defined by these check's repositories.

Check out the check here: [# Scorechecks](#scorechecks)

### .env

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
