# Método longo

Um método que contém linhas demais (começa a ser preocupante a partir de 10 linhas).

## Razões

Coisas são colocadas no método, mas nada é tirado. É mais fácil colocar coisa em um método do que criar um novo, criando um código espaguete.

## Soluções

A regra de ouro é que se você acha que precisa explicar alguma coisa, essa coisa deve ir para um novo método, mesmo se for uma só linha. Se o método tiver um nome descritivo, ninguém vai precisar entrar no código para ver o que ele faz.

### Solução 1: Extração do método

#### Extraia o Método

Pedaço de código que pode ser agrupado.

```go
func printOwing(){
    printBanner();

    // print details
    fmt.Println("name:", nome)
    fmt.Println("amount:", getOutstanding())
}
```

Mova esse código para um método separado e chame na função original.

```go
func printOwing(){
    printBanner()
    printDetails(outstanding float64)
}

func printDetails(outstanding float64){
    fmt.Println("name:", nome)
    fmt.Println("amount:", getOutstanding())
}
```

### Solução 2: Reduzir variáveis locais e parâmetros antes de extrair um método

Utilize as opções abaixo caso variáveis locais e parâmetros estejam interferindo com a extração do método.

#### Substitua a Variável Por um Método

Você coloca o resultado de uma expressão em uma variável local para usar no código depois.

```go
func calculateTotal() float64 {
    basePrice := quantity * itemPrice
    if (basePrice > 1000){
        return basePrice * 0.95
    }
    else {
        return basePrice * 0.98
    }
}
```

Mova a expressão para um método separado e retorne o resultado dela.

```go
func calculateTotal() float64 {
    if basePrice() > 1000 {
        return basePrice() * 0.95
    } else {
        return basePrice() * 0.98
    }
}

func basePrice(){
    return quantity * itemPrice
}
```

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

### Solução 3: Substitua o método com o método de um objeto ou estrutura

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

### Solução 4: Condicionais e iterações

Operadores condicionais e de iteração são boas pistas de que o código pode ser movido para um método separado. Para condicionais, use a Separação da Condicional. Se o problema for as iterações, use o Método de Extração.

#### Separação da Condicional

Se tiver uma condicional complexa (`if-then`/`else` ou `switch`).

```go
if date.Before(SUMMER_START) || date.After(SUMMER_END) {
    charge = quantity * winterRate + winterServiceCharge;
} else {
    charge = quantity * summerRate;
}
```

Separe as partes complicadas da condicional em métodos diferentes: a condição, `then` e `else`.

```go
if isSummer(date) {
    charge = summerCharge(quantity)
} else {
    charge = winterCharge(quantity)
}
```

#### Extração de Método

Você tem um pedaço de código que pode ser agrupado.

```go
func printProperties(users []User) {
    for _, user := range users {
        result := ""
        result += user.Name
        result += " "
        result += user.Age

        fmt.Println(result)

        // ...
    }
}
```
Mova o código para um novo método, se possível contido na estrutura.

```go
func printProperties(users []User) {
    for _, user := range users {
        fmt.Println(user.Properties())
    }
}

func (u *User) Properties(){
    return u.Name + " " + u.Age;
}
```

## Recompensas

- Projetos com métodos menores duram mais. Quanto maior um método ou função for, mais difícil é para entendê-lo e mantê-lo.
- Além disso, métodos maiores escondem código duplicado não desejado mais facilmente.

## A performance é melhorada?

Na maioria dos casos, a quantidade de métodos não causa impacto significante. E agora que você tem um código limpo e compreensível, é mais fácil encontrar métodos eficazes para reestruturar código e obter ganho de performance real.