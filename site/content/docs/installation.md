---
linkTitle: Installation
next: Usage
title: "Installation"
weight: 1
---
First, make sure you have [Go](https://golang.org/dl/) installed on your system.

### Install via `go install`

```bash
go install github.com/pierow2k/polyhymnia@latest
```

### Clone and Build

Alternatively, you can clone the repo and build the application manually:

```bash
git clone https://github.com/pierow2k/polyhymnia.git
cd polyhymnia
sudo make install
```

This will create an executable called `polyhymnia` and install the `man` page to `/usr/local/share/man/man1`.
