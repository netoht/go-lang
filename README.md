# [go-lang](https://golang.org/)

## Instalando Go Lang

Acesse a página de [download](https://golang.org/dl/) e siga as instruções.

```sh
tar -C $HOME/Applications/ -xzf goX.X.X.X.tar.gz
```

Adicionar ao seu arquivo de bashrc as variáveis:

```sh
# Local de instalação
export GOROOT=$HOME/Applications/go
export PATH=$PATH:$GOROOT/bin

# Diretório onde será seu workspace
export GOPATH=$HOME/Github/go-lang/workspace
export PATH=$PATH:$GOPATH/bin
```

Para verificar se sua instalação foi feita corretamente, faça:

```sh
# Diretório do projeto
mkdir -p $GOPATH/src/github.com/user/hello

# Arquivo de hello world
vim $GOPATH/src/github.com/user/hello/hello.go
# Conteúdo:
#    package main
#
#   import "fmt"
#
#   func main() {
#       fmt.Printf("hello, world\n")
#   }

# Cria o executável hello em $GOPATH/bin
go install github.com/user/hello

# Executa o executável gerado
$GOPATH/bin/hello
hello, world
```

Se você viu a mensagem `hello, world` sua instalação foi concluída com sucesso. 
Siga para o passo [como escrever código Go](https://golang.org/doc/code.html).

## Como escrever código Go

Estrutura do `workspace`:

```
tree $GOPATH
src/   # Contém arquivos fonte
pkg/   # Contém pacote de objetos
bin/   # Contém binários executáveis
```

## Referências

- [Tour pela Go Lang](https://tour.golang.org/)
- [Go tool e estrutura de projetos](https://golang.org/doc/code.html)
- [Convenções de Go](https://golang.org/doc/effective_go.html)
- [Lista de pacotes nativos para Go](https://golang.org/pkg/)
- [Lista de pacotes para Go - 1](https://github.com/golang/go/wiki/Projects)
- [Lista de pacotes para Go - 2](https://godoc.org/)
- [Lista de artigos sobre Go](https://golang.org/doc/#articles)
