# Common Go pkgs

## Access private repo

### Adding GOPRIVATE env variable
```bash

add to ~/.bashrc or ~/.profile (depending on OS)

export GOPRIVATE=github.com/CreateView/gocore
```

### Git over ssh
```bash
~/.gitconfig

[url "ssh://git@github.com/"]
	insteadOf = https://github.com/
```

### Git over HTTPS
```bash
$HOME/.netrc

machine github.com login USERNAME password APIKEY

```
### more details at [https://go.dev/doc/faq#git_https](https://go.dev/doc/faq#git_https)
