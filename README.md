# Scripter

![Platform: Linux](https://img.shields.io/badge/platform-linux-lightgrey)

A simple Linux CLI for fast microservice setup with user-defined project templates. Scripter saves development time by letting you reuse scaffolding and automate environment initialization—perfect for microservice developers looking to standardize workflows.

## Table of Contents
- [How it works](#how-it-works)
- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)
- [Roadmap](#roadmap)
- [License](#license)

## How it works

Define and save templates from any directory with `scripter make <templateName>`.  
In your project folder, create a `scripts.json` file (see `examples/scripts.json`) and run predefined scripts using `scripter run <scriptName>`.

## Installation

```
git clone https://github.com/LuminousMyrrh/scripter.git
cd scripter
make
./dist/scripter
```
*Linux only.* 

## Usage

- Use `-h/--help flag` for list of commands and their description

- Create a template:
```
scripter make authentication
```
Copies the current folder to `~/.config/scripter/templates/authentication`.

- Run from scripts.json:
```
scripter run auth
```
Copies the template files into the current directory.

## Features
- Lightweight, fast CLI
- No dependencies
- Templates saved to `~/.config/scripter`
