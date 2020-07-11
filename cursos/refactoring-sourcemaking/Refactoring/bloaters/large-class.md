# Classe ou pacote grande

É uma classe que contenha muitos campos/métodos/linhas de código

## Razões

Classes ou pacotes costumam começar pequenos, mas se tornam abarrotados conforme o programa cresce.
Assim como no caso dos métodos longos, é mais fácil colocarem uma funcionalidade nova em uma classe existente do que criar uma nova classe para a funcionalidade.

## Soluções

Quando uma classe ou pacote tiver muitas responsabilidade, pense em separá-la.

### Solução 1: Extrair a classe ou pacote

Extrair a classe ou pacote ajuda se parte do comportamento da classe grande pode ser separado em um componente diferente.

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

#### Solução 2: Extrair a Subclasse

Extrair subclasses ajuda se parte do comportamento da classe grande for implementado de formas diferentes ou usado em casos raros.

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

##### Extrair Interface

Extrair interfaces ajuda se for necessário ter uma lista de operações e comportamentos que o cliente possa usar.

Vários clientes usam a mesma parte da interface de uma classe.

```go
type Employee struct {
    // ...
}

func (e *Employee) Rate()
func (e *Employee) hasSpecialSkill()
func (e *Employee) Name()
func (e *Employee) Department()
```

Move sua porção idêntica para a interface própria.

```go
type Billable interface {
    func Rate()
    func hasSpecialSkill()
}

type Employee struct {
    // ...
}

func (e *Employee) Rate()
func (e *Employee) hasSpecialSkill()
func (e *Employee) Name()
func (e *Employee) Department()
```

#### Solução 4: Separar os dados de usabilidade do domínio

Se uma classe grande é responsável pela interface gráfica, tente mover dados e comportamentos para um objeto de domínio diferente. Para tal, pode ser ncessário copiar dados em dois lugares para manter a consistência dos dados.

##### Dados Propositalmente Duplicados

```go
type IntervalWindow struct {
    startField string
    endField string
    lengthField string
}

func (i *IntervalWindow) StartField_FocusLost()
func (i *IntervalWindow) EndField_FocusLost()
func (i *IntervalWindow) LengthField_FocusLost()
func (i *IntervalWindow) calculateLength()
func (i *IntervalWindow) calculateEnd()
```

```go
type IntervalWindow struct {
    startField string
    endField string
    lengthField string
}

func (i *IntervalWindow) StartField_FocusLost()
func (i *IntervalWindow) EndField_FocusLost()
func (i *IntervalWindow) LengthField_FocusLost()

type Interval struct {
    start string
    end string
    length string
}

func (i *Interval) calculateLength()
func (i *Interval) calculateEnd()
```

## Recompensas

- Refatorar essas classes poupa pessoas desenvolvedoras de precisar lembrar vários atributos de uma classe.
- Em vários casos, separar classes grandes em partes menores evita duplicação de código e funcionalidade.