# Obsessão por primitivas

Uso de primitivas ao invés de objetos pequenos para tarefas pequenas (moeda, alcances, números de telefone etc).
Uso de constantes para informação (USER_ADMIN_ROLE = 1)
Uso de constantes de string como nomes de campos para uso em arrays.

## Razões

Criar um novo campo de primitiva é muito mais fácil que criar uma nova classe. Primitivas são usadas para "simular" tipos. Logo, ao invés de usar um tipo específico, você tem um conjunto de strings que formam valores para uma entidade.

## Soluções

### Solução 1: Substituir o conjunto de campos por um objeto

#### Substituir Campos Por Objeto

Pode ser possível agrupar alguns campos em sua própria classe. Também é bom mover o comportamento associado a esses campos para essa classe.

```go
package order

var customer string
```

Crie uma nova classe ou pacote, coloque o campo antigo e seu comportamento nela e chame o pacote novo no pacote antigo.

```go
package order

import "customer"

// ...
customer.Name
// ...
```

```go
package customer

var Name string
```

### Solução 2: Campos primitivos em parâmetros de método

Troca parâmetros por objeto.

#### Troque o Parâmetro Por um Objeto ou Estrutura

Seu método contém parâmetros repetidos.

```go
func amountInvoicedIn(start time.Time, end time.Time)
func amountReceivedIn(start time.Time, end time.Time)
func amountOverdueIn(start time.Time, end time.Time)
```

Substitua os parâmetros por objetos ou estruturas.

```go
func amountInvoicedIn(date DateRange)
func amountReceivedIn(date DateRange)
func amountOverdueIn(date DateRange)
```

#### Preserve o Objeto Inteiro

Há diversos valores de um objeto para passar para um método

```go
low := daysTempRange.getLow()
high := daysTempRange.getHigh();
withinPlan := plan.withinRange(low, high)
```

Passe o objeto inteiro ao invés disso.

```go
withinPlan = plan.withinRange(daysTempRange)
```

### Solução 3: Livre-se dos códigos de tipo

Quando dados complicados são armazenados em variáveis.

#### Substitua o Tipo por Classe

Uma classe ou estrutura tem campo que contém código que pode extraído para um tipo.

```go
type Person struct {
    O int
    A int
    B int
    AB int
    bloodgroup int
}
```

Crie um novo pacote e use-o como atributo da classe original.

```go
type Person struct {
    BloodGroup
}

type BloodGroup struct {
    O int
    A int
    B int
    AB int
}
```

#### Substitua o Tipo por Subclasses

Há tipo que afeta diretamente o comportamento do programa (valores que aivam vários códigos com condicionais).

```go
package employee

type Employee struct {
    ENGINEER int
    SALESMAN int
    type int
}
```

Crie subclasses ou pacotes para cada tipo. Depois, extraia o comportamento relevante da classe original para essas subclasses. Substitua o fluxo de controle com polimorfismo.

```go
package employee

import (
    "engineer"
    "salesman"
)

// ...
engineer.Type
Salesman.Type
// ...
```

```go
package engineer
```

```go
package salesman
```

#### Substitua o Tipo por Estado/Padrão Strategy

Há tipo que afeta diretamente o comportamento do programa mas não é possível usar subclasses para se livrar dele..

```go
package employee

type Employee struct {
    ENGINEER int
    SALESMAN int
    type int
}
```

Substitua o código por um objeto de estado. Use interface para uniformizar os tipos.

```go
type Employee interface{
    Type()
}

type Engineer struct {
    Type int
}

func (e *Engineer) Type(){
    return e.Type
}

type Salesman struct {
    Type int
}

func (s *Salesman) Type(){
    return s.Type
}

func Type(e Employee){
    return e.Type()
}
```

#### Substitua o Array por Objeto

Você tem um array que contém vários tipos de dados.

```go
row := []string{"Liverpool", "15"}
```

```go
type Performance struct {
    Name string
    Wins string
}

// ...
row := &Performance{
    Name: "Liverpool",
    Wins: "15,
}
```

## Recompensas

- O código se torna mais flexível
- É mais fácil entender e organizar o código; não é mais necessário adivinhar o motivo daquelas constantes estranhas ou por que estão num array
- Mais facilidade para encontrar códigos duplicados
