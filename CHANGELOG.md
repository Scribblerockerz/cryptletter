# Changelog

All notable changes to this project will be documented in this file.

## [3.1.6] - 2025-04-18

### Dependencies

- Bump golang.org/x/net from 0.37.0 to 0.38.0 ([#53](https://github.com/Scribblerockerz/cryptletter/issues/53))
- Bump github.com/go-co-op/gocron from 1.5.0 to 1.37.0 ([#54](https://github.com/Scribblerockerz/cryptletter/issues/54))
- Bump github.com/minio/minio-go/v7 from 7.0.10 to 7.0.90 ([#56](https://github.com/Scribblerockerz/cryptletter/issues/56))
- Bump github.com/gorilla/mux from 1.8.0 to 1.8.1 ([#57](https://github.com/Scribblerockerz/cryptletter/issues/57))
- Bump github.com/spf13/viper from 1.7.0 to 1.20.1 ([#55](https://github.com/Scribblerockerz/cryptletter/issues/55))

### Documentation

- Replace npm with pnpm

### Features

- Add freebsd build target to build script

### Fixes

- Typo in dependabot config

### Miscellaneous Tasks

- Golang and some web security updates ([#50](https://github.com/Scribblerockerz/cryptletter/issues/50))
- Re-add pnpm-lock.yaml
- Add stable node version via .nvmrc
- Remove package-lock.json in favor of pnpm-lock.yaml

## [3.1.5] - 2024-04-17

### Fixes

- Trim null bytes from receiving payload

## [3.1.4] - 2024-04-16

### Fixes

- Race condition while decoding body

## [3.1.3] - 2022-05-26

### Dependencies

- Bump axios from 0.21.1 to 0.21.2 in /web ([#33](https://github.com/Scribblerockerz/cryptletter/issues/33))
- Bump follow-redirects from 1.14.3 to 1.14.8 in /web ([#36](https://github.com/Scribblerockerz/cryptletter/issues/36))
- Bump url-parse from 1.5.3 to 1.5.10 in /web ([#38](https://github.com/Scribblerockerz/cryptletter/issues/38))
- Bump minimist from 1.2.5 to 1.2.6 in /web ([#39](https://github.com/Scribblerockerz/cryptletter/issues/39))
- Bump async from 2.6.3 to 2.6.4 in /web ([#40](https://github.com/Scribblerockerz/cryptletter/issues/40))
- Bump nanoid from 3.1.23 to 3.3.4 in /web ([#41](https://github.com/Scribblerockerz/cryptletter/issues/41))

### Features

- Add git-cliff for changelog generation

### Miscellaneous Tasks

- Replace yarn with pnpm for node dependencies
- Split release and build scripts for ease of use
- Add vue runtime dependencies

## [3.1.2] - 2021-09-09

### Dependencies

- Bump path-parse from 1.0.6 to 1.0.7 in /web ([#28](https://github.com/Scribblerockerz/cryptletter/issues/28))
- Bump postcss from 7.0.35 to 7.0.36 in /web ([#31](https://github.com/Scribblerockerz/cryptletter/issues/31))
- Bump url-parse from 1.5.1 to 1.5.3 in /web ([#30](https://github.com/Scribblerockerz/cryptletter/issues/30))

### Other

- Fix quickstart settings in readme.md
- Fix docker cmd declaration

## [3.1.1] - 2021-06-15

### Dependencies

- Bump dns-packet from 1.3.1 to 1.3.4 in /web ([#23](https://github.com/Scribblerockerz/cryptletter/issues/23))
- Bump ws from 6.2.1 to 6.2.2 in /web ([#24](https://github.com/Scribblerockerz/cryptletter/issues/24))

### Other

- Update README.md ([#25](https://github.com/Scribblerockerz/cryptletter/issues/25))
- Fix missing serve command in docker image

## [3.1.0] - 2021-05-28

### Other

- Update .gitignore
- Add password protection for ltetter creation
- Add toaster to display error messages
- Cleanup
- Add minio as development storage
- Update page title
- Add basic file encryption handling
- Add attachment meta data handling
- Add decryption for attachments
- Cleanup and reafactor
- Add translations and restyle attachments
- Add in progress placeholder for attachments
- Add ttl tracking to local asset handler
- Add scheduler to trigger cleanups
- Disable attachment support by default
- Update dependencies
- Cleanup debug messages
- Move eslint config to standalone file
- Fix formatting issues with eslint
- Update roadmap
- Add s3 support
- Refactor default config options
- Update readme and cleanup

## [3.0.1] - 2021-05-04

### Other

- Update build script
- Update npm dependencies

## [3.0.0] - 2021-05-04

### Other

- Update build script
- Add license
- Update node dependencies
- Migrate project structure to a more golang way
- Remove templating server
- Add vue3 setup
- Migrate some components over to vue
- Migrate new and show pages to vue
- Add CORS handling for dev environment
- Migrate functionality to vue app
- Add initialization stylings
- Add animated encryption
- Update roadmap
- Fix automatic env mapping
- Fix missing translations for raw state buttons
- Add base url fallback for frontend app
- Add support for build with embeded web code
- Add example for extendable locale configuration
- Add css and js injection for customizations
- Update readme for v3
- Update changelog

## [2.0.1] - 2019-09-01

### Other

- Update node dependencies
- Update roadmap

## [2.0.0] - 2019-05-06

### Other

- Initial commit
- Add configuration support with toml files
- Add todos
- Add dynamic partil registration
- Migrated old theme and templates to the new codebase
- Implemented first featureset
- Moved things around
- Updated build script
- Add proper logging
- Update roadmap todo list
- Add frontend refactoring
- Fix TTL calculation
- Implement readable remaining time
- Update remaining time format
- Add environment configuration and prepare docker build
- Add new todos to the list
- Enhance theming capabilities
- Add noscript message
- Add interchangable text partials
- Add changelog

<!-- generated by git-cliff -->
