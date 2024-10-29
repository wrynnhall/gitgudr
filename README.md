# Gitgudr

Gitgudr is a collection of essential Git tools designed to simplify and automate Git workflow tasks. Gitgudr combines multiple helper commands to streamline the often-repetitive or easily overlooked aspects of working with Git repositories. Built with Go, Gitgudr aims to be lightweight, intuitive, and highly customizable.

## Features

Gitgudr provides a suite of commands to help with everyday Git tasks:

### WIPr (In Progress)
> **Purpose**: Automatically stage, commit, and push uncommitted changes.

WIPr is perfect for those moments when you wrap up work but forget to push your changes. This tool automatically stages any unstaged changes, commits them with a timestamped or user-provided message, and pushes to the current branch. WIPr ensures you don’t lose any recent work due to oversight.

Usage:
```bash
gitgudr wipr -m "your optional commit message"
```

### Cleanr (Planned)
> **Purpose**: Identify and clean up unused Git repositories in your file system.

Old, inactive repositories can clutter your file system, taking up space and making navigation more cumbersome. Cleanr scans your directories for Git repositories and identifies those that haven’t been accessed or updated within a customizable time frame, giving you the option to archive or delete them.

Usage:
```bash
gitgudr cleanr -p /path/to/scan -a 30d
```
Options:
- `-p, --path` - Specify the path to scan for old repositories.
- `-a, --age` - Define the age threshold for inactivity (e.g., `30d` for 30 days).

### Chatr (Planned)
> **Purpose**: ChatGPT-powered assistant for Git workflows.

Chatr is a ChatGPT-integrated client that aids in common Git operations, like suggesting commit messages based on `git diff` outputs or explaining complex merge conflicts. Just type a command, and Chatr provides helpful suggestions to streamline your Git workflow with AI assistance.

Usage:
```bash
gitgudr chatr commit-suggest
```

## Installation

To install Gitgudr, make sure you have [Go](https://golang.org/) installed, then run:

```bash
go install github.com/wrynnhall/gitgudr
```

## Usage

To see a list of available commands and options:

```bash
gitgudr --help
```

Each command is designed to be simple, efficient, and highly useful for Git users at any level.

## Roadmap

- **Additional Commands**: We're continuously developing new commands to enhance Gitgudr's functionality.
- **Improved AI Integration**: Expanding Chatr’s capabilities to include more context-aware suggestions.
- **Custom Configuration**: Allowing users to define preferences for WIPr and Cleanr for more control.

