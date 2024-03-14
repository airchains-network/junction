ignite chain build

# after adding the binary file to the profile env path

junctiond init testing --chain-id testing --default-denom amf

junctiond keys add charlie --keyring-backend test
junctiond keys add diana --keyring-backend test
junctiond keys add ethan --keyring-backend test
junctiond keys add frank --keyring-backend test
junctiond keys add grace --keyring-backend test

junctiond keys list --keyring-backend test

junctiond genesis add-genesis-account charlie 10000000000000000amf,10000000000000000uatom,\
10000000000000000usdc,10000000000000000uosmo,10000000000000000uevmos \
    --keyring-backend test

junctiond genesis add-genesis-account diana 10000000000000000amf,10000000000000000uatom,\
10000000000000000usdc,10000000000000000uosmo,10000000000000000uevmos \
    --keyring-backend test

junctiond genesis add-genesis-account ethan 10000000000000000amf,10000000000000000uatom,\
10000000000000000usdc,10000000000000000uosmo,10000000000000000uevmos \
    --keyring-backend test

junctiond genesis add-genesis-account frank 10000000000000000amf,10000000000000000uatom,\
10000000000000000usdc,10000000000000000uosmo,10000000000000000uevmos \
    --keyring-backend test

junctiond genesis add-genesis-account grace 10000000000000000amf,10000000000000000uatom,\
10000000000000000usdc,10000000000000000uosmo,10000000000000000uevmos \
    --keyring-backend test


junctiond genesis gentx charlie 1000000000000000amf \
    --keyring-backend test \
    --chain-id booba

junctiond genesis collect-gentxs

junctiond start

# and we are good to go 😄👍🔥🔥