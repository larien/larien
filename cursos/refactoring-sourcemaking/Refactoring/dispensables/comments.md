# Comentários

Um método tem comentários explicativos.

## Razões

Comentários são criados com a melhor das intenções, quando a pessoa percebe que seu código não é intuitivo ou óbvio. Nesses casos, comentários são como um desodorante mascarando o cheiro de código pobre que poderia ser melhorado.
__O melhor comentário é um bom nome para um método ou classe.__
Tente mudar a estrutura de código de forma que torne os comentários  desnecessários.

## Soluções

### Solução 1: Extrair a variável

Se um comentário tem a intenção de explicar uma expressão complexa, a expressão pode ser separada em subexpressões compreensíveis.

#### Extrair Variável

Você tem uma expressão que é difícil de entender.

```go
func renderBanner(){
    if platform["MAC"] > -1 && browser["IE"] > -1 && wasInitialized() && resize > 0 {
        // do something
    }
}
```

Coloque o resultado da expressão ou suas partes em variáveis separadas que são autoexplicativas.

```go
func renderBanner(){
    isMacOS := platform["MAC"] > -1
    isIE := browser["IE"] > -1
    wasResized := resize > 0

    if (isMacOS & isIE && wasInitialized() && wasResized){
        // do something
    }
}
```

### Solução 2: Extrair método

Se um comentário explica uma parte do código que pode ser transformada em um método separado. O nome do novo método pode ser tirado do próprio comentário.

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

### Solução 3: Renomear o método

Se um método já foi extraído mas comentários ainda são necessários para explicar o que método faz, dê a ele um nome autoexplicativo.

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

### Solução 4: Usar asserção

Se você precisa verificar regras sobre um estado que é necessário para o sistema funcionar, asserte-as.

#### Use Asserção

Para uma parte do código funcionar corretamente, certas condições ou valores devem ser verdade.

```go
func expensiveLimit() int {
    // should have either limit or a primary project
    if expenseLimit != NULL_EXPENSE {
        return expenseLimit
    }
    return primaryProject.MemberExpenseLimit()
}
```

Substitua essas presunções com verificações específicas.

```go
func expensiveLimit() (int, error) {
    if expenseLimit != NULL_EXPENSE || primaryProject != nil {
        return 0, errors.New("empty value")
    }

    if expenseLimit != NULL_EXPENSE {
        return expenseLimit
    }
    return primaryProject.MemberExpenseLimit()
}
```

## Recompensas

- O código se torna mais intuitivo e óbvio

## Quando ignorar

- Comentários podem ser úteis quando explicam _por que_ alguma coisa está sendo implementada de forma específica ou explicando algoritmos complexos (quando todos os outros métodos para simplificar o algoritmo já foram testados)
