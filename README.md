# LoadBalancerClient

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

A load balancer client for querying ebay's backend

A Client for querying ebay's backend, with failure tolerate and round-robin load balance strategy.


## Table of Contents

- [Background](#background)
- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)

## Background

Suppose there is a web service, for example, the URL is http://www.ebay.com. And behind the DNS, there are multiple servers providing the service, and each server has a different IP.
Please write a client program to query the IPs behind the DNS, and send 100 requests to different IPs  according to client-side loadbalancing strategy in order to achieve better performance and lower  lantency.

## Install

clone the repository.

## Usage

```sh
$ ./bin/load_balancer
```

## Maintainers

[@Mr-win7](https://github.com/Mr-win7).

