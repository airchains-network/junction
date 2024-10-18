# Junction
Welcome to the new version of Junction, a cutting-edge settlement layer that coordinates transaction flow and ensures proper settlement. This update introduces the Track-gate module, which integrates new external sequencers, including the Espresso sequencer for enhanced transaction sequencing. It also retains full compatibility with the latest Ignite CLI, optimizing project structure, execution layers, and validator management.

## Key Features

- Custom execution layers, batch processing, and a democratic validator system for enhanced data management.
- Efficient management of verification keys and dynamic track pilots selection.
- Enhanced query functionalities and robust batch handling.
- Integration of the `trackgate` module for seamless integration with external sequencer support.
- Incorporation of Espresso sequencer for improved transaction ordering and significantly enhancing overall system functionality.

Download the Correct Binary
---------------------------

Before starting, ensure that you download the appropriate binary for your operating system:

1.  Visit the [release page](https://github.com/airchains-network/junction/releases/tag/v0.2.0) and download the correct binary for your OS:

    -   **Linux**: `junctiond-linux-amd64`
    -   **macOS**: `junctiond-darwin-amd64`
    -   **Windows**: `junctiond-windows-amd64`
2.  After downloading the binary, rename it to `junctiond`:

```shell
mv <downloaded-binary> junctiond
```

3. Make it executable:

```shell
chmod +x junctiond
```

4. Move it to your local binary path:

```shell
sudo mv junctiond /usr/local/bin
```

Installation & Initialization
-----------------------------

```shell
junctiond init <moniker> --chain-id junction --default-denom amf
```

### Start Junction

```shell
junctiond start
```

Additional Information
----------------------

For detailed instructions, configurations, and updates, refer to the [docs](https://docs.airchains.io). Stay connected with our community for the latest developments and support.

Thank you for your continued support of the Airchains ecosystem. We look forward to a successful and smooth upgrade, bringing enhanced scalability and efficiency to our network through the integration of the Trackgate Module and incorporating the Espresso sequencer.