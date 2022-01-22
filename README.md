# logbook-api

logbook-api is a REST API to make divelogs in ATMOS Platform to public.

![](https://github.com/umatare5/logbook-api/blob/images/promo.gif)

## Syntax

```bash
$ logbook
NAME:
   logbook - Serve your dive-logs

USAGE:
   logbook COMMAND [options...]

VERSION:
   0.1.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --web.listen-address value, -L value   Set IP address (default: "0.0.0.0")
   --web.listen-port value, -P value      Set port number (default: 8080)
   --web.debug, -D                        Set debug mode (default: false)
   --config.atmos.server value, -s value  Set FQDN for ATMOS Platform (default: "http://localhost:8080/") [$LOGBOOK_ATMOS_API]
   --config.atmos.token value, -t value   Set access-tokens to access to ATMOS Platform [$LOGBOOK_ATMOS_TOKEN]
   --help, -h                             show help (default: false)
   --version, -v                          print the version (default: false)
```

- **✓ Supported variables**

  Following variables are supported.

  | Environment Variable | Description                              |
  | :------------------- | :--------------------------------------- |
  | LOGBOOK_ATMOS_API    | FQDN for ATMOS Platform.                 |
  | LOGBOOK_ATMOS_TOKENS | Access-tokens to use ATMOS Platform API. |

## Install

You can download the binary archive on [releases page](https://github.com/umatare5/logbook-api/releases).

After unarchived, start logbook-api easily with the following command.

```bash
./logbook
```

## Endpoints

### Admin API

|     | Method | Path              | Parameters | Description            |
| :-- | :----- | :---------------- | :--------- | :--------------------- |
| ✅  | GET    | /api/admin/health | -          | Get status of this API |

### Divelog API

|     | Method | Path                        | Parameters                     | Description              |
| :-- | :----- | :-------------------------- | :----------------------------- | :----------------------- |
| ✅  | GET    | /api/v1/divelogs            | limit=number,<br>cursor=string | Get dive-logs            |
| ✅  | GET    | /api/v1/divelog/{divelogId} | divelogId=String               | Get detail of a dive-log |
