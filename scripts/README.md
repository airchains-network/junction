# Scripts

These scripts are copied from the [Cosmos-SDK](https://github.com/cosmos/cosmos-sdk/tree/v0.42.1/scripts) respository 
with minor modifications. All credits and big thanks go to the original authors.

Please note that a custom [fork](github.com/regen-network/protobuf) by the Regen network team is used.
See [`go.mod`](../go.mod) for version.

## protocgen.sh

If used this script and got this error 

```shell
Generating gogo proto code
./scripts/protocgen.sh: line 16: buf: command not found
```

You need to install buf.

```shell
VERSION="1.50.0"  # Replace with the latest version if different
BIN="/usr/local/bin"
sudo curl -sSL "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" -o "${BIN}/buf"
sudo chmod +x "${BIN}/buf"
```
