# Passgen

## About

Passgen is a command line app written in [Go](https://golang.org/) using the [Cobra](https://github.com/spf13/cobra) library. This project showcases a code structure that makes it easy to testing all the business logic of the CLI app.

The app itself is a simple password generator inspired by an [XKCD comic](https://xkcd.com/936/):

![Password Strength XKCD comic](https://imgs.xkcd.com/comics/password_strength.png)

### Motivation

I'm a big fan of the [Cobra](https://github.com/spf13/cobra) library. However, I find it hard to test my code that depends on Cobra directly and Cobra Generator scaffolding tool produces a bit ugly code. In this demo project I tried to overcome those issues by organizing my code with a little bit of forethought.

### Built with

* [Cobra](https://github.com/spf13/cobra)
* [Viper](https://github.com/spf13/viper)
* [English words from bip-0039](https://github.com/bitcoin/bips/blob/master/bip-0039/bip-0039-wordlists.md)

## Getting started

I suggest you to clone the project locally and explore the code yourself. My idea was to keep the domain logic as simple as possible, while still demonstrating how to achieve good test coverage.

### Prerequisites

You'll need to have [Go installed locally](https://golang.org/dl/) in order to test or build passgen.

### Testing

To run tests execute:

```sh
$ go test ./...
```

in the project's root directory.

### Installation

To build the app execute:

```sh
$  go build ./cmd/passgen/
```

in the project's root directory. This will generate the executable `passgen` file that you can run.

## Usage

Passgen app is not really intended for widespread usage, but you can still play around with it. Underlying Cobra library provides CLI related features, like a sub-command structure, help flag, file config etc.

The actual sub-command for generating passwords is `generate`. You can run it with:

```sh
$ ./passgen generate
richswordmeadowmanage
```

Or, for more features, like this:

```sh
./passgen generate -csm 24
Agree Melt Light Honey Window!
```

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

## Authors

- Josip Antoliš - [Antolius](https://github.com/Antolius)
