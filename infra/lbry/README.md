# README

## Generate a keypair (privatekey + address)

```bash
$ go run keygen.go
LBRY_PK=eagPs6RBxmTQyjni3K7vqPNBwjN4o5R8CEwP4eyHavMJMz29MCen
LBRY_ADDRESS=smtdQvMJRLaWwNaFUjdBtFzUR4evxQJcB9
```

## Build your docker container

```bash
docker build .
```

## Run the container

```bash
# Regtest
docker run -p 18332:18332 lbrycrd:latest "smtdQvMJRLaWwNaFUjdBtFzUR4evxQJcB9"
```
