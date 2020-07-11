# Hierarquias de herança paralelas

Quando tenta criar uma subclasse para uma classe, você se vê precisando criar uma subclasse para outra classe.

## Razões

A adição de classes novas complicam as classes já existentes.

## Soluções

### Solução 1: Mesclar as hierarquias de classe (se possível)

É possível remover as duplicações em duas etapas. Primeiro você referencia a instância de uma hierarquia em outra. Depois, remova a hierarquia da classe com referência.

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

## Recompensas

- Reduz a duplicação do código
- Pode melhorar a organização do código

## Quando ignorar

- às vezes ter hierarquias paralelas é só uma forma de evitar uma bagunça ainda maior na arquitetura do programa. Se tudo isso tornar o código ainda pior, reverta tudo e se acostume com esse código.
