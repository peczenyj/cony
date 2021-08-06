# Cony

High-level AMQP 0.9.1 client library. It's wrapper around low-level [rabbitmq/amqp091-go](https://github.com/rabbitmq/amqp091-go/) library.

# Goals

Provide a way to work with AMQP declaratively

# Requirments

The library uses [atomic.Value](http://golang.org/pkg/sync/atomic/#Value), so Go 1.4+ is needed.

# Documentation
[![codecov](https://codecov.io/gh/peczenyj/cony/branch/master/graph/badge.svg?token=BW91XYOLO5)](https://codecov.io/gh/peczenyj/cony)
[![GoDoc](https://godoc.org/github.com/peczenyj/cony?status.svg)](https://godoc.org/github.com/peczenyj/cony)
[![Build Status](https://travis-ci.com/peczenyj/cony.svg?branch=master)](https://travis-ci.com/peczenyj/cony)

# Thread-safety

Cony is thread-safe as long as [rabbitmq/amqp091-go](https://github.com/rabbitmq/amqp091-go) is thread-safe. It's recommended to open AMQP channel per thread, so in case of `cony` it should be `Consumer` `Producer` per goroutine.

# License

BSD 2 clause - see LICENSE for more details.

## Fork

This is a fork of github.com/assembla/cony
