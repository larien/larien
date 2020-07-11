# Campo temporário

São utilizados apenas em situações específicas e ficam vazios fora delas.

## Razões

Costumam ser usados em um algoritmo que precisa de várias entradas e não é usada no restante do projeto. É o tipo de código difícil de entender.

## Soluções

### Solução 1: Extrair campos temporários e comportamentos relacionados

Campos temporários e todo o código podem ser colocados em uma classe separada.

#### Extraia a Classe ou Programa

Quando uma classe faz o trabalho de duas, coisas estranhas acontecem.

```go
package person

type Person struct {
    name string
    officeAreaCode int
    officeNumber int
}

func (p *Person) TelephoneNumber() string {
    // ...
}
```

Crie uma nova classe e coloque os campos e métodos responsáveis pela funcionalidade relevante nele.

```go
package person

type Person struct{
    name string
    number Telephone
}

func (p *Person) telephoneNumber() string {
    return p.number.TelephoneNumber()
}
```

```go
package telephone

type Telephone struct {
    officeAreaCode int
    officeNumber int
}

func (t *Telephone) TelephoneNumber() string {
    // ...
}
```

#### Substitua o Método Com o Método de um Objeto ou Estrutura

Há um método em que as variáveis locais estão tão interligadas que não é possível extrair um método.

```go
package order

func price() float64 {
    var (
        primaryBasePrice float64
        secondaryBasePrice float64
        tertiaryBasePrice float64
    )
    // perform long computation
}
```

Crie uma estrutura ou classe para o método e mantenha seu escopo lá.

```go
package order

type PriceCalculator {
    primaryBasePrice float64
    secondaryBasePrice float64
    tertiaryBasePrice float64
}

func price(order Order) float64 {
    pc := &PriceCalculator{
        // copy relevant information from the order received
    }
    return pc.compute()
}

func (p *PriceCalculator) compute(){
    // perform long computation
}
```

### Solução 3: Use objetos vazios

Use objetos vazios e integre no luga da condicional que era usada para verificar a existência dos campos temporários.

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

- Organização de código melhor e mais clara.
