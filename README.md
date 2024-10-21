# filen
Open-source linter to control file size on golang.

## Why filen?
Vertical size of a file should be typically 200 lines, with upper limit of 500,
but you can configure it for your project.
```
defaultМaxLinesNum - 500
defaultМinLinesNum - 5
```

- Very small files are useless
- Very big files difficult to understand it (probably you have architecture problems)


## Usage


```
git clone git@github.com:DanilXO/filen.git

cd dist

./filen <path_for_check>
```

Available parameters:
* `-maxLinesNum int` - the maximum number of lines in a file. `500` by default
* `-minLinesNum int` - the minimum number of lines in a file. `5` by default