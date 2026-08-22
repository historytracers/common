[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](CODE_OF_CONDUCT.md)
![CodeRabbit Pull Request Reviews](https://img.shields.io/coderabbit/prs/github/historytracers/common?utm_source=oss&utm_medium=github&utm_campaign=historytracers%2Fcommon&labelColor=171717&color=FF570A&link=https%3A%2F%2Fcoderabbit.ai&label=CodeRabbit+Reviews)

# Common

Shared common code for the [History Tracers](https://github.com/historytracers) organization.

This repository contains libraries, utilities, types, and configuration shared across multiple History Tracers projects, including:

- [historytracers/historytracers](https://github.com/historytracers/historytracers) — the main History Tracers application
- Other projects under the [historytracers](https://github.com/historytracers) organization

## Purpose

Keeping common code in a single place ensures consistency, reduces duplication, and makes it easier to maintain shared logic across the entire ecosystem.

## Usage

Each consuming project references this repository as a dependency. See the respective project's documentation for installation and usage instructions.

## Development

```bash
# Install dependencies
npm install

# Build
npm run build

# Run tests
npm test
```

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md) if present, otherwise refer to the contributing guide in the main repository.

All contributions are expected to follow the [Code of Conduct](./CODE_OF_CONDUCT.md).

## License

See [LICENSE](./LICENSE) if present, otherwise refer to the license in the main repository.
