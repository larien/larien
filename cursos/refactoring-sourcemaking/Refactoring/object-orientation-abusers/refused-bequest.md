# Herança desnecessária

Se uma subclasse usa poucas coisas herdadas dos pais, a hierarquia é desnecessária. Os métodosnão desejados podem só ficar sem ser usados ou ser redefinidos e gerar exceções inesperadas.

## Razões

Alguém criou herança entre classes para usar o código na superclasse, mas a superclasse e a subclasse são completamente diferentes.

## Soluções

### Solução 1: Substituir herança por delegação

Se a herança não fizer sentido e a subclasse não tiver nada em comum com a superclasse, troque a herança por delegação.

#### Substitua Herança Por Delegação

Você tem uma subclasse que usa apenas parte dos métodos da superclasse.

```go
type Vector interface {
    func isEmpty() bool
}

type Stack type {
    // ...
}

func (s *Stack) isEmpty() bool {
    // ...
}
```

Crie um campo no objeto da superclasse, delegue esses métodos para ela e se livre da herança.

```go
type Vector struct {
    Stack
}

type Stack struct {
    // ...
}

func (s *Stack) isEmpty() bool {
    // ...
}

// somewhere in the client code
return vector.Stack.isEmpty()
```

### Solução 3: Extrair a superclasse

Extraia campos e métodos indesejados da superclasse, coloque em uma nova classe e ambas herdam dessa.

#### Extraia a Superclasse

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

## Recompensas

- Melhora a organização e torna o código mais claro,já que você não precisa descobrir por que a classe `Cachorro` herda da classe `Cadeira` (apesar de ambos terem quatro pernas).
