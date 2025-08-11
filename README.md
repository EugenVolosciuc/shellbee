<p align="center">
  <img src="assets/shellbee-logo.png" alt="ShellBee Logo" width="150">
</p>

<h1 align="center">ShellBee</h1>

**shellbee** is a lightweight CLI tool for storing, organizing, and running your favorite terminal commands in a flash. Think of it as your busy little helper in the shell — always ready to buzz into action.

---

## Features

- Save frequently used commands with easy-to-remember aliases
- Quickly run saved commands by their alias
- List all your stored command aliases
- Search for commands by keyword
- Delete aliases you no longer need
- Simple, intuitive command syntax with short aliases

---

## Installation

_(TODO: Add installation instructions here — e.g., building from source, downloading binaries, etc.)_

---

## Usage

### Save a command alias

```
shb save {{name}} {{command}}
shb s {{name}} {{command}}
```

Example:

```bash
shb save greet "echo 'Hello World'"
```

---

### Run a saved command

```
shb run {{name}}
shb r {{name}}
```

Example:

```bash
shb run greet
```

---

### List saved command aliases

```
shb list
shb l
```

---

### Search commands by keyword

```
shb find {{query}}
shb f {{query}}
```

Example:

```bash
shb find echo
```

---

### Delete a command alias

```
shb delete {{name}}
shb d {{name}}
```

---

### Show CLI help

```
shb help
shb h
```

---

## Configuration & Storage

ShellBee stores your commands in a JSON file located in a sensible place depending on your operating system, such as:

- macOS: `~/Library/Application Support/shellbee/shellbee.json`
- Linux: `$XDG_DATA_HOME/shellbee/shellbee.json` or `~/.local/share/shellbee/shellbee.json`
- Windows: `%AppData%\shellbee\shellbee.json`

This file is created automatically on first use.

---

## Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](#) or submit a pull request.

---

## License

[MIT License](https://github.com/EugenVolosciuc/shellbee/blob/main/LICENSE.md)
