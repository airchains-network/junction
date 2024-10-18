# Junction
Welcome to the new version of Junction, a cutting-edge settlement layer that coordinates transaction flow and ensures proper settlement. This update introduces the Track-gate module, which integrates new external sequencers, including the Espresso sequencer for enhanced transaction sequencing. It also retains full compatibility with the latest Ignite CLI, optimizing project structure, execution layers, and validator management.

## Key Features

- Custom execution layers, batch processing, and a democratic validator system for enhanced data management.
- Efficient management of verification keys and dynamic track pilots selection.
- Enhanced query functionalities and robust batch handling.
- Integration of the `trackgate` module for seamless integration with external sequencer support.
- Incorporation of Espresso sequencer for improved transaction ordering and significantly enhancing overall system functionality.


## Installation & Initialization

```bash
wget https://github.com/airchains-network/junction/releases/download/v0.2.0/junctiond
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

Thank you for your continued support of the Airchains ecosystem. We look forward to a successful and smooth upgrade, bringing enhanced scalability and efficiency to our network through the integration of the Trackgate Module and incorporating the Espresso sequencer.
