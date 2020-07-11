# Declarações switch

Exstem operadores `switch` complexos ou sequência de `ifs`.

## Razões

Geralmente o código de um `switch` pode ser espalhado em partes diferentes de um programa.
Como regra de ouro, pense em polimorfismo quando ver um `switch`.

## Soluções

### Solução 1: Isolar o operador `switch`

#### Extração do Método

Você tem um pedaço de código que pode ser agrupado.

```go
package order

// ...

func calculateTotal(products Products) (total float64) {
    for _, product := range products {
        total += product.Quantity * product.Price
    }

    // apply regional discounts
    switch user.Country {
        case "US": total *= 0.85
        case "RU": total *= 0.75
        case "CN": total *= 0.9
        // ...
    }

    return
}
```

Mova esse código para um novo método (ou função) e substitua o código antigo com uma chamada para o método.

```go
package order

// ...

func (u *User) calculateTotal(products Products) (total float64) {
    for _, product := range products {
        total += product.Quantity * product.Price
    }
    total = u.applyRegionalDiscounts(total)

    return
}

func (u *User) applyRegionalDiscounts(total float64) (result float64) {
    switch u.Country {
        case "US": result *= 0.85
        case "RU": result *= 0.75
        case "CN": result *= 0.9
        // ...
    }
    return
}
```

#### Mova o Método

Um método é usado mais em uma classe do que dentro da própria.

```go
package order

// ...

func (u *User) calculateTotal(products Products) (total float64){
    for _, product := range products {
        total += product.Quantity * product.Price
    }
    total = u.applyRegionalDiscounts(total)

    return
}

func (u *User) applyRegionalDiscounts(total float64) (result float64) {
    switch u.Country {
        case "US": result *= 0.85
        case "RU": result *= 0.75
        case "CN": result *= 0.9
        // ...
    }
    return
}
```

Crie um novo método na classe que use mais o método e mova o código para o método anterior. Transforme o código método original em referência para o novo método da outra classe ou remova-o completamente.

```go
package order

func (u *User) calculateTotal(d Discounts) (total float64) {
    // ...
    total = d.applyRegionalDiscounts(total, user.Country)
    total = d.applyCoupons(total)
    // ...
}
```

```go
package discounts

type Discounts struct {
    // ...
}

func (d *Discounts) applyRegionalDiscounts(total float64, country string) float64 {
    result := total
    switch country {
        case "US": result *= 0.85
        case "RU": result *= 0.75
        case "CN": result *= 0.9
        // ...
    }
    return result
}

func (d *Discounts) applyCoupons(total float64){
    // ...
}
```

### Solução 2: Livre-se dos códigos de tipo

Quando dados complicados são armazenados em variáveis.

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

### Solução 3: Substituir a condicional por polimorfismo

#### Substitua Condicional Por Polimorfismo

```go
package bird

// ...
func speed(type int) float64 {
    // ...
    switch type {
        case EUROPEAN:
            return baseSpeed()
        case AFRICAN:
            return baseSpeed() - loadFactor() * numberOfCoconuts
        case NORWEGIAN_BLUE:
            if isNailed {
                return 0
            }
            return baseSpeedByVoltage(voltage)
        // ...
    }
}
```

Coloque o método em comum dentro de uma interface e implemente em diferentes tipos.

```go
package bird

type Bird interface {
    Speed() float64
}

type European struct {
    // ...
}

func (e *European) Speed() float64 {
    return baseSpeed()
}

type African struct {
    // ...
}

func (a *African) Speed() float64 {
    return baseSpeed() - loadFactor() * numberOfCoconuts
}

type NorwegianBlue struct {
    // ...
}

func (n *NorwegianBlue) Speed() float64 {
    if isNailed {
        return 0
    }
    return baseSpeedByVoltage(voltage)
}

// somewhere in client code
speed = bird.Speed()
```

### Solução 4: Substituir os parâmetros do switch por métodos explícitos

Se não houverem muitas condições no operador e todos chamam o mesmo método com parâmetros diferentes, o polimorfismo não serve para muita coisa.

#### Substitua Parâmetros com Métodos Explícitos

Esse método é quebrado em partes, cada qual executada dependendo do valor do parâmetro.

```go
func setValue(name string, value int){
    if name == "height" {
        heigth = value
        return
    }
    if name == "width" {
        width = value
        return
    }
    // ...
}
```

Extraia as partes individuais do método em seus próprios.

```go
func setHeight(arg int){
    heigth = arg
}

func setWidth(arg int){
    width = arg
}
```

### Solução 5: Use objetos nulos

Se uma das opções for nula.

#### Use Objetos Nulos

Já que alguns métodos retornam `nil` ao invés de objetos reais, você pode precisar checar por eles no seu código.

```go
if customer == nil {
    plan = BillingPlan.Basic()
} else {
    plan = customer.Plan()
}
```

Ao invés de `nil`, retorne um objeto vazio que exibe o comportamento padrão.

```go
// replace null values with Null-object
if order.Customer != nil {
    order.Customer = &Customer{}
}

// use Null-object as if it's normal subclass
plan = customer.Plan()
```

## Recompensas

- Organização de código melhorada

## Quando ignorar

- Nã há necessidade de mudar nada se o operador `switch` fizer ações simples
- Quando os operadores `switch` são usadas por padrões de projeto `factory` (`Factory Method` ou `Abstract Method`) para selecionar uma classe criada.
