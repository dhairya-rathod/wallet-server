# Wallet

This is a server repo of a wallet and created with the intention to learn GoLang. You may find some errors or code snippets which is not as per the Go coding standard.

This project includes REST APIs for account creation & CRUD operation of expenses. I used chi library for routing.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

clean up binary from the last build
```bash
make clean
```