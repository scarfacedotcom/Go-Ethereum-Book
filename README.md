# Ethereum Development with Go

> A little guide book on [Ethereum](https://www.ethereum.org/) Development with [Go](https://golang.org/) (golang)

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/miguelmota/merkletreejs/master/LICENSE)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](#contributing)


## Languages

* [English](en/)

## Help & Support

- Join the [#ethereum](https://gophers.slack.com/messages/C9HP1S9V2/) channel on the [gophers slack](https://invite.slack.golangbridge.org/) for Go (golang) help

- The [Ethereum StackExchange](https://ethereum.stackexchange.com/) is a great place to ask general Ethereum question and Go specific questions

## Development

Install dependencies:

```bash
make install
```

Run gitbook server:

```bash
make serve
```


Visit [http://localhost:4000](http://localhost:4000)

## Contributing

Pull requests are welcome!

If making general content fixes:

- please double check for typos and cite any relevant sources in the comments.

If updating code examples:

- make sure to update both the code in the markdown files as well as the code in the [code](code/) folder.

If wanting to add a new translation, follow these instructions:

1. Set up [development environment](#development)

2. Add language to `LANGS.md`

3. Copy the the `en` directory and rename it with the 2 letter language code of the language you're translating to (e.g. `zh`)

4. Translate content

5. Set `"root"` to `"./"` in `book.json` if not already set

## Thanks

Thanks to [@gzuhlwang](https://github.com/gzuhlwang) for the Chinese translation.

And thanks to all the [contributors](https://github.com/miguelmota/ethereum-development-with-go-book/graphs/contributors) who have contributed to this guide book.

## License

Released under the [CC0-1.0](./LICENSE) license.

© [Miguel Mota](https://github.com/miguelmota)