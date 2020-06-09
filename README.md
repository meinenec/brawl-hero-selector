# brawl-hero-selector

brawl-hero-selector is a Discord chat bot written in Go.

## Usage

- `!brawl` returns a list of random heroes based on a specified number and hero pool
- `!brawl n pool` returns a list of n heroes from the specified hero pool, where 0<n<26

Supported Hero Pools
- all
- brawl

## Building
 
```
docker build -t brawl-hero-selector .
```

## Running

```
export BOT_TOKEN=<discord bot token>
docker pull meinenec/brawl-hero-selector:latest
docker run -e BOT_TOKEN meinenec/brawl-hero-selector:latest
```
