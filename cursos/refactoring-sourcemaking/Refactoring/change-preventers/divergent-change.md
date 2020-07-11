# Mudança que atinge vários lugares

Você precisa mudar muitos métodos sem reclação quando faz alteração em uma classe. Por exemplo, quando precisa adicionar um novo tipo de produto, você precisa mudar os métodos para encontrar, mostrar e ordenar produtos.

## Razões

Esses grupos costumam aparecer por estruturação de dados ruim ou "programação copypasta".

## Soluções

### Solução 1: Extraia a classe

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

### Solução 2: Combinar as classe através de herança

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

##### Extraia a Subclasse

Uma classe tem funcionalidades usadas em certos casos.

```go
// Job Item
// getTotalPrice()
// getUnitPrice()
// getEmployee()
```

Crie uma subclasse e use em outros casos.

```go
// Job Item
// getTotalPrice()
// getUnitPrice()
// getEmployee()

// Labor Item
// getUnitPrice()
// getEmployee()
```

## Recompensas

- Melhora a organização do código
- Reduz duplicação do código
- Simplifica o suporte
