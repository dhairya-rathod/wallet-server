# Wallet Server Repository

Welcome to the Wallet Server Repository! This project was initiated with the primary goal of learning GoLang. While exploring this repository, you might encounter occasional errors or code snippets that do not adhere strictly to Go coding standards. however, each serves as a learning opportunity.

## Project Overview

This project primarily focuses on developing REST APIs for various functionalities, including account creation and performing CRUD (Create, Read, Update, Delete) operations on expenses. The routing mechanism is implemented using the chi library.

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
