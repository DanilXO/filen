# filen

Open-source linter to control file size on Go.

## Why filen?

Vertical size of a file should be typically 200 lines, with upper limit of 500,
but you can configure it for your project.

- default `maxLines`: `500`.
  Very big files are difficult to understand it (probably you have architecture problems)
- default `minLines`: `5`.
  Very small files are useless.

## Usage

```bash
go install github.com/DanilXO/filen/cmd/filen@latest

./filen <path_for_check>
```

Available parameters:

* `-maxLines int` - the maximum number of lines in a file. `500` by default
* `-minLines int` - the minimum number of lines in a file. `5` by default
* `-ignoreComments bool` - ignore comment lines or not. `false` by default
