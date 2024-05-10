# Junction

Welcome to the new version of Junction, a cutting-edge blockchain platform designed for secure and efficient data management. Building on our robust foundation, this version introduces significant enhancements and compatibility with the latest Ignite CLI, aiming to further streamline project structure, execution layers, and validator management.

## Key Features

- Custom execution layers, batch processing, and a democratic validator system for enhanced data management.
- Efficient management of verification keys and dynamic track pilots selection.
- Enhanced query functionalities and robust batch handling.

## Installation & Initialization

Build from source:
```bash
git clone https://github.com/airchains-network/junction.git
cd ~/junction/cmd
env CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o ./build/junctiond main.go
# Find your binary in ~/junction/cmd/build/
```

```bash
wget https://github.com/airchains-network/junction/releases/download/v0.1.0/junctiond
chmod +x junctiond
sudo mv junctiond /usr/local/bin
```
```bash
junctiond init <moniker> --chain-id junction --default-denom amf
```

### Start Junction

```bash
junctiond start
```

## Additional Information

For detailed instructions, configurations, and updates, refer to the [docs]("https://docs.airchains.io"). Stay connected with our community for the latest developments and support.
