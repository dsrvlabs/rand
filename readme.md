# RAND
**RAND** is a blockchain built using Cosmos SDK and Tendermint
RAND provide a random number for various dAPPs that need random number like gacha game, lottery and so on.
The random number generated by CCV validator set. we use CCV Validator set as a seed for the rand.

## Get started

Install [go](https://go.dev/dl/)

## Build and install to go bin path

```
make install
```

## Initialize config

Come up with a moniker for your node, then run:

```
randd init $MONIKER
```
 
 
 
## Launch with genesis file or run as standalone chain

To launch as a consumer chain, download and save shared genesis file to `~/.rand/config/genesis.json`. Additionally add peering information (`persistent_peers` or `seeds`) to `~/.rand/config/config.toml`

To instead launch as a standalone, single node chain, run:

```
randd add-consumer-section
```

## Launch node

```
randd start
```
