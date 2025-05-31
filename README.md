## qn-decode

网易云音乐格式转化

A command tool for transfering  `qmcflac`|`qmc0`|`qmc3`|`ncm` to `mp3` or `flac`.

The repo is used for learning, if there is any infringement, please contact the author to delete


## Installing

### GO
Using `qn-decode` is easy. First, use go get to install the latest version of the library. This command will install the `qn-decode` generator executable along with the library and its dependencies:
```
go install github.com/jovianlee/qn-decode@latest
```

## Usage
```
A command tool for transfering  'qmcflac'
        |'qmc0'|'qmc3'|'ncm' to 'mp3' or 'flac'.

Usage:
  qn-decode [command]

Available Commands:
  decode      decode music file
  help        Help about any command
  version     Print the version number of qn-decode

Flags:
      --config string   config file (default is $HOME/.qn-decode.yaml)
  -h, --help            help for qn-decode
  -t, --toggle          Help message for toggle

Use "qn-decode [command] --help" for more information about a command.
```

### Reference
 - https://github.com/MBearo/qmcdump
 - https://github.com/yoki123/ncmdump
 - https://github.com/luanxuechao/qn-decode

### Example
```
$  decode -d ~/Downloads
```
```
$  decode -f ~/Downloads/xxxx.qmc3
```
