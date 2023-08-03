# Florida Men

## Description

This is a small project, that prints the florida man story fitting to the current date. It's main focus is to take first steps in Go.

Currently there is a simple fetcher that fetches the story of the current day from https://floridamanbirthday.org and writes it to the console.
More fetchers and writers can be added by implementing the `StoryWriter` and `StoryFetcher` Interfaces.

### Build 

To build an executible run
``` bash
bin/build
```

### Development

To run the project while developing use
```bash
go run floridamen.go
```

Excute tests:
```bash
bin/tests
```

Excute linter:
```bash
bin/lint
```