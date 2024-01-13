# Code-Tex-Converter

## Overview

Code-Tex-Converter is a command-line tool that converts programming source code in a specified directory into LaTeX listing format. This tool simplifies and streamlines the process of incorporating source code into LaTeX documents during documentation tasks.

## Installation

To install Code-Tex-Converter, follow these steps (requires a Go language development environment):

```bash
go get github.com/yudai2929/code-tex-converter
```

## Usage

Use the tool with the following command syntax:

```bash
code-tex-converter --target <directory-path> [options]
```

### Options

- `--target`, `-t`: Specifies the target directory to convert to LaTeX listing format. This option is required.

- `--lstinput`, `-i`: If specified, it uses lstinputlisting to incorporate source code into LaTeX. Default is false.

- `--base`, `-b`: Specifies the base directory from which relative paths are calculated. If not specified, the current working directory is used as the base.

## Example

To convert source code in a specific directory to LaTeX listing format:

```bash
code-tex-converter --target /path/to/source/code
```

To use lstinputlisting for incorporating source code:

```bash
code-tex-converter --target /path/to/source/code --lstinput
```

To convert source code in a specific directory to LaTeX listing format, with a specific base directory:

```bash
code-tex-converter --target /path/to/source/code --base /path/to/base/directory
```
