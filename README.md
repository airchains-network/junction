# Junction

Welcome to **Junction**, a cutting-edge settlement layer designed to coordinate transaction flow and ensure efficient settlement across decentralized ecosystems. This latest update introduces powerful new modules: **`wasm`**, **`vrf`** (Verifiable Random Function), and **`cipher-ledger`**, significantly enhancing transaction security, confidentiality, and execution efficiency.

## 🚀 Key Enhancements in This Update

### **New Modules Introduced**
- **`wasm` Module**: Brings full support for WebAssembly (WASM), enabling efficient smart contract execution.
- **`vrf` Module**: Introduces Verifiable Random Functions, enhancing randomness-based selection processes.
- **`cipher-ledger` Module**: Manages Fully Homomorphic Encryption (FHE) functionalities, ensuring private and secure on-chain computations.

### **Improved Query & Transaction Processing**
- Optimized verification key management.
- Enhanced query functionalities for seamless blockchain interactions.
- Robust batch handling for large-scale transaction processing.

## 📥 Installation & Setup

### Download the Binary
Ensure you download the appropriate binary for your operating system:

- **Visit the [release page](https://github.com/airchains-network/junction/releases/tag/v0.3.0-rc2)** and select the correct binary:
  - **Linux**: `junctiond-linux-amd64`
  - **macOS**: `junctiond-darwin-amd64`
  - **Windows**: `junctiond-windows-amd64`

- Rename the downloaded file for convenience:

  ```shell
  mv <downloaded-binary> junctiond
  ```

- Make it executable:

  ```shell
  chmod +x junctiond
  ```

- Move it to your system’s binary path:

  ```shell
  sudo mv junctiond /usr/local/bin
  ```

### Initialize the Network

To set up **Junction**, run the following command:

```shell
junctiond init <moniker> --chain-id junction --default-denom amf
```

This initializes a node with your chosen moniker and configures it for the **Junction** network.

### Start the Junction Node

Launch your node with:

```shell
junctiond start
```

Your node is now active and participating in the **Junction** settlement layer!

## 📚 Documentation & Support

For comprehensive setup guides, configuration details, and troubleshooting steps, refer to our official **[documentation](https://docs.airchains.io)**.

Stay engaged with our **community** for real-time discussions, updates, and support.

## 🎯 Vision & Future Roadmap
The integration of `wasm`, `vrf` and `cipher-ledger` modules marks a significant milestone in our journey toward an advanced, privacy-preserving, and high-performance settlement layer. With these upgrades, **Junction** is poised to deliver enhanced security, scalability, and seamless transaction coordination across the Airchains ecosystem.
