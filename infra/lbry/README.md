# README

## Generate a keypair (privatekey + address)

```bash
$ go run keygen.go
LBRY_PK=cV9WB2d7VTCJkkJy5RrzkQvqtiFJvAXbL8XUte8YRBH9zG2u1LTu
LBRY_ADDRESS=mufonb3XbgFWH6ayP84sz81x6rS9oCxvqR
```

## Build your docker container

```bash
docker build .
```

## Run the container

```bash
# Regtest
docker run -p 18332:18332 lbrycrd:latest "mufonb3XbgFWH6ayP84sz81x6rS9oCxvqR"
```
