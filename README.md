# gospa

[![GoTemplate](https://img.shields.io/badge/go/template-black?logo=go)](https://github.com/SchwarzIT/go-template)

Example app that combines go's [`html/template`](https://pkg.go.dev/html/template) for handling static content on server-side  and [`vitejs`](https://vitejs.dev/) for client-side spa. Created with [`go-template`](https://github.com/schwarzit/go-template).

## Setup
```bash
# setup vite view
make install-view
```

The project uses `make` to make your life easier. If you're not familiar with Makefiles you can take a look at [this quickstart guide](https://makefiletutorial.com).

Whenever you need help regarding the available actions, just use the following command.

```bash
make help
```


===== todo from here ====

## Test & lint

Run linting

```bash
make lint
```

Run tests

```bash
make test
```
