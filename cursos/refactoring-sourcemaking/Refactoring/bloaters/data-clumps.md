# Falta de encapsulamento de dados

Às vezes partes diferentes do código contém grupos idênticos de variáveis (como a conexão à base de dados). Essa repetição deveria ser transformada na própria classe.

Regra de ouro: Deletar um dos valores de dados e veja se a outra ainda faz sentido. Se não for o caso, é um bom sinal de que esse grupo de variáveis deve ser combinado dentro de um objeto.

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

### Solução 2: Passe um objeto como parâmetro

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

### Solução 3: Preserve o objeto inteiro

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

### Solução 4: Mova o código para uma classe

## Recompensas

- Melhora o entendimento e a organização do código
- Diminui o tamanho do código

## Quando ignorar

Passar um objeto inteiro como parâmetro de um método ao invés de passar apenas os valores (tipos primitivos) pode criar uma dependência indesejada entre as duas classes.
