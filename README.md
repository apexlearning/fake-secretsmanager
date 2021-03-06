fake-secretsmanager
===================

`fake-secretsmanager` is a stand-in for the full [AWS Secrets Manager](https://aws.amazon.com/secrets-manager/) for testing, local development, kitchen runs, and other such things where using the real deal is neither needed nor desirable.

It should go without saying that this absolutely should not be used for production (or probably even staging) - it does not use SSL to encrypt traffic, and its "secrets" storage is in a plain JSON file.

Installation
------------

To build `fake-secretsmanager` from source, assuming you have Go installed and your Go development environment set up properly, run:

```
$ go get -u github.com/apexlearning/fake-secretsmanager
```

`fake-secretsmanager` has been built using golang 1.10. It may build with earlier or later versions of the compiler, but this has not been tested.

A `Dockerfile` is also provided, if running `fake-secretsmanager` inside docker is more convenient. Build it with `docker build -t your-handy-name/fake-secretsmanager .`, and run it with:

```
$ docker run -p 7887:7887 -v /path/to/your/secrets.json:/opt/fake-secretsmanager/data/secrets.json --name fakesm -d your-handy-name/fake-secretsmanager
```

Usage
-----

```
Usage:
  fake-secretsmanager [OPTIONS]

Application Options:
  -v, --version       Print version info.
  -a, --addr=         IP address to listen on. Default: ':7887'. [$FAKESM_ADDR]
  -f, --secrets-json= Path to JSON file containing the secrets in a hash. The
                      JSON hash key names are the secret names. If the secret
                      is itself JSON, it needs to be escaped and stuffed in
                      there as a normal string. [$FAKESM_SECRETS_JSON]

Help Options:
  -h, --help          Show this help message
```

The `secrets.json` file provides an example of how to format the JSON file that stores the secrets. Each secret must be a string. If it is not a string, it must be properly quoted and escaped as per the example escaped JSON in that file.

To use `fake-secretsmanager`, supply a custom endpoint to the AWS cli or in your code's secretsmanager client constructor.

CLI:

```
$ aws secretsmanager get-secret-value --endpoint-url http://localhost:7887 --secret-id foo/json/escaped/sssssh

$ aws secretsmanager list-secrets --endpoint http://localhost:7887
```

Ruby:

```
secretsmanager = Aws::SecretsManager::Client.new(
  region: my_region,
  endpoint: 'http://localhost:7887',
  # .....
)
```

Supported Functionality
-----------------------

Currently, `fake-secretsmanager` supports the following [AWS Secrets Manager API](https://docs.aws.amazon.com/secretsmanager/latest/apireference/Welcome.html) functionality:

* [GetSecretValue](https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html)
* [ListSecrets](https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecrets.html)

Other functionality is not present yet, but could be added if the need arises (or someone contributes it).

TODO
----

See the TODO file.

BUGS
----

See the BUGS file.

Author
------

Jeremy Bingham (<jeremy.bingham@apexlearning.com>)

Copyright
---------

Copyright 2018, Apex Learning, Inc.

AWS Secrets Manager is copyright Amazon Web Services, Inc. (or maybe its affiliates, or even its parent).

License
-------

fake-secretsmanager is licensed under the terms of the Apache 2.0 License. See the LICENSE file for details.
