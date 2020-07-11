# Tiro de espingarda

Fazer algumas modificações requer que você faça alterações pequenas em classes demais.

## Razões

Uma única responsabilidade foi separada entre várias classes.

## Soluções

### Solução 1: Consolidar a responsabilidade em uma única classe

#### Mova o Método

Um método é usado mais em uma classe do que na própria.

```go
package class1

func method() {
    // ...
}
```

Crie um novo método para a classe que o usa mais e mova o método para lá. Transforme o método original em referência para o novo método na outra classe ou remova-o por completo.

```go
package class2

func Method() {
    // ...
}
```

```go
package class1

import "class2"

// somewhere in the client code
class2.Method()
```

#### Mova o Campo

Um campo é usado mais em uma classe do que na própria.

```go
package class1

var Field type
```

Crie um novo campo para a classe que o usa mais e mova o método para lá. Transforme o campo original em referência para o novo campo na outra classe ou remova-o por completo.

```go
package class2

var Field type
```

```go
package class1

import "class2"

// somewhere in the client code
class2.Field
```

### Solução 2: Remover classes redundantes

Se mover o código par a mesma classe deixa as classes originais quase vazias, elimite as classes agora redundantes.

#### Alinhe as Classes

Uma classe não faz quase nada e não é responsável por nada e não há responsabilidade a mais planejadas para ela.

```go
package person

type Person struct {
    Name string
}

func telephoneNumber() {
    // ...
}
```

```go
package telephoneNumber

type TelephoneNumber struct {
    OfficeAreaCode int
    OfficeNumber int
}

func telephoneNumber() {
    // ...
}
```

Mova todas as funcionalidades dessa classe para outra.

```go
package person

type Person struct {
    Name string
    OfficeAreaCode int
    OfficeNumber int
}

func telephoneNumber() {
    // ...
}
```

## Recompensas

- Melhor organização
- Menos duplicação de código
- Manutenção mais fácil
