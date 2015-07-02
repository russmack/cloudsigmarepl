# cloudsigma repl

A REPL for the CloudSigma IaaS REST API.

[![Build Status](https://travis-ci.org/russmack/cloudsigmarepl.svg?branch=master)](https://travis-ci.org/russmack/cloudsigmarepl)

## Usage
```
Set your login in the config.json.
go build
./cloudsigmarepl
```

## Example Session
```
Welcome to the CloudSigma IaaS REPL.
------------------------------------

Type "help" for help.

> set config location
Enter service location:
> zrh
Response:
 zrh

> cloud status
Response:
 {"free_tier": {"dssd": 53687091200, "mem": 1073741824}, "free_tier_monthly": {"tx": 5497558138880}, "guest": true, "signup": true, "sso": ["github", "twitter", "google", "facebook", "linkedin"], "trial": true}

> create server
Name:
> Repl Test Server
CPU:
> 1000
Memory:
> 536870912
VNC password:
> mypassword
Response:
 {"objects": [{"auto_start": false, "context": true, "cpu": 1000, "cpu_model": null, "cpu_type": "amd", "cpus_instead_of_cores": false, "drives": [], "enable_numa": false, "grantees": [], "hv_relaxed": false, "hv_tsc": false, "hypervisor": "kvm", "jobs": [], "mem": 536870912, "meta": {}, "name": "Repl Test Server", "nics": [], "owner": {"resource_uri": "/api/2.0/user/xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/", "uuid": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"}, "permissions": [], "pubkeys": [], "requirements": [], "resource_uri": "/api/2.0/servers/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/", "runtime": null, "smp": 1, "status": "stopped", "tags": [], "uuid": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", "vnc_password": "mypassword"}]}

> bye
Bye!
```

## Features

- [X] Cloud status
- [X] Locations
- [X] Capabilities
- [X] Profile
- [X] Balance
- [X] Subscriptions
- [X] Transactions
- [X] Pricing
- [X] Discounts
- [X] Current usage
- [X] Burst usage
- [X] Daily usage
- [X] Licenses
- [X] Notification preferences [list, edit]
- [X] Notification contacts [list, create, edit, delete]
- [X] Servers [list, create, delete, start, stop, shutdown]
- [X] Drives [list, list detailed, get single, create, delete]
- [X] Snapshots [list, list detailed, get single, delete]
- [X] Vlans [list, detailed list, get single]
- [X] IP addresses [list, detailed list, get single]
- [X] Access Control Lists [list, get single]
- [X] Tags [list, get single]
- [X] Logs [list]
- [X] Firewall Policies [list]

## License
BSD 3-Clause: [LICENSE.txt](LICENSE.txt)

[<img alt="LICENSE" src="http://img.shields.io/pypi/l/Django.svg?style=flat-square"/>](LICENSE.txt)
