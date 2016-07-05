statusAS-front
==============

Dashboard frontend for statusAS-api.

Has some golang, vuejs, bootstrap and css.

URL is http://server2.mydom.local:9000/dashboard or the IP/hostname you entered in the config.toml file.

Adjust config.toml according to your setup.
```toml
[frontserver]
frontservername = "server2.mydom.local"
frontserverport = 9000
apiservername = "server1.mydom.local"
apiserverport = 8080
```

Then build, run and check the url.
```
$ go build
$ ./statusAS-front
```

Usage:
```
$ ./statusAS-front --help
Usage of statusAS-front:
  -createtemplate
        Create an example config.toml file

$ ./statusAS-front -createtemplate
config.toml already exist in directory.

$ rm config.toml

$ ./statusAS-front -createtemplate
config.toml created.
```

If some service has messages you can mouse over the service name and it will show you the messages.
