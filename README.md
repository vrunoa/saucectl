saucectl
========
![build](https://github.com/saucelabs/saucectl-internal/workflows/saucectl%20pipeline/badge.svg?branch=master)
========

A command line interface for the Sauce Labs Testrunner Toolkit. This repository contains the Go binary that you use to kick off tests. If you look for more documentation on it, please have a look into our [example repo](https://github.com/saucelabs/testrunner-toolkit).

# Development Requirements

- [Go](https://golang.org/) (v1.14 or higher)
- [Homebrew](https://brew.sh/) (v2.2.13 or higher)

# Install

Run the following to install all dependencies:

```sh
$ make install
```

# Build

To build the project, run:

```sh
$ make build
```

# Test

To execute unit tests, run:

```sh
$ make test
```

# Using `saucectl`

## The `new` Command
```sh
saucectl new
```

This command will ask you to choose one of the frameworks: 
- [Puppeteer](https://github.com/puppeteer/puppeteer)
- [Playwright](https://github.com/microsoft/playwright)
- [TestCafe](https://github.com/DevExpress/testcafe) 

After that, a `./sauce/config.yml` file and an example test under
the `tests` directory will be created, where you can start working from.

## The `run` Command
```sh
saucectl run
```
This command will run the test based on the `./.sauce/config.yml` file.

### Flags

#### `config`
```sh
saucectl run --config <path>
```
Using the `--config` will run the tests specified by that config file.

#### `region`
```sh
saucectl run --region <region>
```
Using the `--region` flag will set the Sauce Labs region for the test execution.
The region corresponds to the available regions at saucelabs.com and affects
where your job information and assets are going to be stored.

#### `timeout`
```sh
saucectl run --timeout <seconds>
```
Using the `--timeout` flag will set the test timeout for the test runner framework. 

# Licensing
`saucectl` is licensed under the Apache License, Version 2.0. See [LICENSE](https://github.com/saucelabs/saucectl/blob/master/LICENSE) for the full license text.
