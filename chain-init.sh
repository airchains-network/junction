ignite chain build

# after adding the binary file to the profile env path

junctiond init testing --chain-id testing --default-denom uamf

junctiond keys add charlie --keyring-backend test
junctiond keys add diana --keyring-backend test
junctiond keys add ethan --keyring-backend test
junctiond keys add frank --keyring-backend test
junctiond keys add grace --keyring-backend test

junctiond keys list --keyring-backend test

junctiond genesis add-genesis-account charlie 2000000000000000uamf,78058977200000uatom,\
822030800000uusdc,129888250600000uosmo,138474773600000uevmos, 205946733800000utia \
    --keyring-backend test

junctiond genesis add-genesis-account diana 2000000000000000uamf,78058977200000uatom,\
822030800000uusdc,129888250600000uosmo,138474773600000uevmos, 205946733800000utia \
    --keyring-backend test

junctiond genesis add-genesis-account ethan 2000000000000000uamf,78058977200000uatom,\
822030800000uusdc,129888250600000uosmo,138474773600000uevmos, 205946733800000utia \
    --keyring-backend test

junctiond genesis add-genesis-account frank 2000000000000000uamf,78058977200000uatom,\
822030800000uusdc,129888250600000uosmo,138474773600000uevmos, 205946733800000utia \
    --keyring-backend test

junctiond genesis add-genesis-account grace 2000000000000000uamf,78058977200000uatom,\
822030800000uusdc,129888250600000uosmo,138474773600000uevmos, 205946733800000utia \
    --keyring-backend test


junctiond genesis gentx charlie 100000000000000uamf \
    --keyring-backend test \
    --chain-id booba

junctiond genesis collect-gentxs

junctiond start

# and we are good to go 😄👍🔥🔥