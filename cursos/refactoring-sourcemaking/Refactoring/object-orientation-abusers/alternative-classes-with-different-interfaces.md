# Classes alternativas com interfaces diferentes

Duas classes fazem a mesma coisa, mas têm métodos com nomes diferentes.

## Razões

A pessoa não sabia que a funcionalidade que estava criando já existia.

## Soluções

Tente colocar a interface das classes em um denominador em comum.

### Solução 1: Renomear os métodos

Torne-os idênticos em todas as classes.

#### Renomeie o Método

O nome de um método não explica o que o método faz.

```go
func (c *Customer) snm() {
    // ...
}
```

Renomeie o método.

```go
func (c *Customer) secondName() {
    // ...
}
```

### Solução 2: Torne as assinaturas e implementações dos métodos idênticos

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

#### Adicione Parâmetros

Um método não tem dados o suficiente para executar certas tarefas.

```go
package customer

func contact() {
    // ...
}
```

Crie um novo parâmetro para passar os dados necessários.

```go
package customer

func contact(date time.Time) {
    // ...
}
```

#### Parametrize os Métodos

Vários métodos fazem coisas parecidas e a única diferença são seus valores, números ou operações internas.

```go
package employee

func fivePercentRaise() {
    // ...
}

func tenPercentRaise() {
    // ...
}
```

Combine os métodos usando um parâmetro que vai receber o valor especial necessário.

```go
package employee

func raise(percentage float64) {
    // ...
}
```

### Solução 3: Extraia apenas as partes idênticas

Se apenas parte da funcionalidade está duplicada, tente extrair uma superclasse.

#### Extraia Superclasse

Você tem duas classes com campos e métodos em comum.

```go
type Department interface {
    totalAnnualCost() float64
    name() string
    headCount() int
}

type Employee interface {
    annualCost() float64
    name() string
    ID() int
}
```

Crie uma interface para ambas e trate os dados por composição.

```go
type Party interface {
    annualCost() float64
    name() string
}

type Employee struct {
    AnnualCost float64
    ID int
}

type Department struct {
    AnnualCost float64
    HeadCount int
}

func (e *Employee) annualCost() float64 {
    // ...
}

func (e *Employee) name() string {
    // ...
}

func (e *Employee) getID() int {
    // ...
}

func (d *Department) annualCost() float64 {
    // ...
}

func (d *Department) name() string {
    // ...
}

func (d *Department) headCount() int {
    // ...
}
```

### Solução 4: Delete os clones

Após decidir qual solução usar e implementá-la, você vai conseguir apagar uma das classes.

## Recompensas

- Se livrar de código duplicado
- O código se torna mais legível e compreensível já que você não precisa descobrir o motivo por trás da criação da segunda classe

## Quando ignorar

Às vezes mesclar classes é impossível ou difícil demais, então não faz sentido. Por exemplo, quando as classes estão em pacotes diferentes em que cada uma tem sua própria versão da classe.
