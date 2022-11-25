# LOVE
**love** is a blockchain built using Cosmos SDK and Tendermint
We forked strange consumer chain.
Thanks for strangelove's contributions and support to the Game of Chains.
Strange + Love = The Best Consumer Chains of Game of Chains ❤️

## Get started

Install [go](https://go.dev/dl/)

## Build and install to go bin path

```
make install
```

## Initialize config

Come up with a moniker for your node, then run:

```
loved init $MONIKER
```
 
 
 
## Launch with genesis file or run as standalone chain

To launch as a consumer chain, download and save shared genesis file to `~/.love/config/genesis.json`. Additionally add peering information (`persistent_peers` or `seeds`) to `~/.love/config/config.toml`

To instead launch as a standalone, single node chain, run:

```
loved add-consumer-section
```

## Launch node

```
loved start
```
