# Common Go pkgs

## Access private repo

### Git over HTTPS
```bash
$HOME/.netrc

machine github.com login USERNAME password APIKEY

```

### Git over ssh
```bash
~/.gitconfig

[url "ssh://git@github.com/"]
	insteadOf = https://github.com/
```

### more details at [https://go.dev/doc/faq#git_https](https://go.dev/doc/faq#git_https)
