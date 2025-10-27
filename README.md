# Scripter
A simple CLI tool for those who want to easily setup their microservices!
Scripter allows you to define and save project templates and use them later.
With *script.json* in your working directory you can define scripts which are uses your templates.
Work only in Linux!

## How it works
You can add new template by using `make` and providing a template name.
Then in directory where you want to use your template create `scripts.json` (example in repo)
and use `run` with provided script name from `scripts.json`

## Examples
1. Create new template - `scripter make authentication`
    - Copies current directory to a `.config/scripter/templates/authentication`
2. Use script from `scripts.json` by `scripter run auth`
    - Copies templates to current directory

## Features 
- Lightweight (no dependencies at all)
- Fast
- Saves your templates in `.config/scripter` directory

## Installation
- Clone this repo:
```
git clone https://github.com/LuminousMyrrh/scripter.git
cd scripter
```
- Bulid with make:
```
make
```
- Run:
```
./dist/main
```

## Roadmap
| Version   | Features     |
| --------- |:------------:|
|    0.1    | basic impl   |
|    0.2    | 'run' command|
|    0.3    | 'make' command|
|    0.4    | 'rm' command|
