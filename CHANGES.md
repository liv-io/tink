# CHANGES

This file lists changes made to the container. It follows semantic versioning
guidelines. The content is sorted in reverse chronological order and formatted
to allow easy grepping by scripts.

The headers are:
- bugs
- changes
- enhancements
- features

## 0.1.10 (2026-11-02)

### Changes

- Consolidate container build into the repository
- Add Containerfile, cmd.sh, and Go toolchain GPG verification keys
- Rename server.go to main.go following standard Go conventions

### Enhancements

- Update Echo from v5.4.3 to v5.4.4

## 0.1.9 (2026-10-25)

### Enhancements

- Update Echo from v5.4.0 to v5.4.3
- Update Go from 1.27.1 to 1.27.3

## 0.1.8 (2026-10-04)

### Enhancements

- Update Echo from v5.3.0 to v5.4.0
- Update Go from 1.26.5 to 1.27.1

## 0.1.7 (2026-07-29)

### Changes

- Update Gate API currency pair

## 0.1.6 (2026-07-21)

### Enhancements

- Update Echo from v5.3.0 to v5.3.1

## 0.1.5 (2026-07-13)

### Enhancements

- Update Echo from v5.2.1 to v5.3.0
- Update Go from 1.26.4 to 1.26.5

## 0.1.4 (2026-06-18)

### Enhancements

- Update Echo from v5.1.1 to v5.2.1
- Update Go from 1.26.2 to 1.26.4

## 0.1.3 (2026-06-08)

### Changes

- Remove exchange Huobi

## 0.1.2 (2026-05-02)

### Enhancements

- Update Echo from v5.0.4 to v5.1.1
- Update Go from 1.25.7 to 1.26.2

## 0.1.1 (2026-02-15)

### Enhancements

- Update Echo from v5.0.3 to v5.0.4
- Update Go from 1.25.4 to 1.25.7

## 0.1.0 (2026-02-12)

### Enhancements

- Improve 'getPriceData' function
- Improve error handling and error messages
- Upgrade echo from v4 to v5
- Update echo from v5.0.0 to v5.0.1
- Update echo from v5.0.1 to v5.0.2
- Update echo from v5.0.2 to v5.0.3

### Features

- Add fetch functions for Bitfinex, Bitget, Bitrue, BTSE, Coinbase, Crypto, and Gate
- Add helper function 'validatePrice' for data validation
