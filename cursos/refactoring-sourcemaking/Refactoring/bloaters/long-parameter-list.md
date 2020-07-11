# Lista de parâmetros longa

Mais de três ou quatro parâmetros por método.

## Razões

Pode ter havido uma mescla de vários algoritmos em um único método.
Também acontecem quando há esforço para tornar classes mais independentes umas das outras.
Um método pode usar os dados de seu próprio objeto. Se o objeto atual não tiver todos os dados necessários, outro objeto com esses dados pode ser passado como parâmetro.

## Soluções

### Solução 1: Substitua o parâmetro por chamada de método

Verifique quais valores são passados como parâmetro e substitua por chamadas de outro objeto.

#### Substitua o Parâmetro por Método

Chamar um método passando seus resultados como parâmetros de outro método, enquanto o método poderia chamar o método diretamente.

```go
basePrice := quantity * itemPrice
seasonDiscount := seasonalDiscount()
fees := Fees()
finalPrice := discountedPrice(basePrice, seasonDiscount, fees)
```

Ao invés de passar o valor por parâmetro, tente colocar uma chamada dentro do escopo do método.

```go
basePrice := quantity * itemPrice
finalPrice := discountedPrice(basePrice)
```

### Solução 2: Preserve o objeto inteiro

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

### Solução 3: Passe um objeto como parâmetro

#### Passe o Objeto Como Parâmetro

Você pode unir os campos em um único objeto.

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

## Recompensas

- Código menor e mais legível
- Pode revelar código duplicado que não havia sido percebido

## Quando ignorar

- Quando causar dependência desnecessária entre classes
