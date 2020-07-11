# Web service in Go with Bill Kennedy

https://github.com/ardanlabs/service

Você pode não concordar com as decisões desse projeto, mas no que você acredita?

## Introdução

Ele veio de 10 anos de C e C# na plataforma Microsoftt e o primeiro baque foi não ter IDE nem padrões de como organizar o código. Isso é extremamente importante porque evita que seu projeto dê errado.

É importante dar mais importância à separação de camadas ao invés de agrupamento.

1. Fazer funcionar
2. Aplicar engenharia do que funciona

Não fazer as coisas fáceis de se fazer, mas fazê-las fácil de compreender.

Uma pessoa média não consegue lidar com mais de cinco coisas ao mesmo tempo na cabeça. Conforme desenvolvemos o projeto, é importante deixar uma quantidade pequena de camadas para lidar.

Go é uma tecnologia de duas camadas: a linguagem e o hardware. Java e C#, por exemplo, têm três: a linguagem, a máquina virtual e o hardware.

Um desenvolvedor não consegue entender mais do que 10 mil linhas de código ao mesmo tempo.

Linhas de código não são uma boa métrica de produtividade, e sim uma métrica de saúde mental da equipe.

Go é uma linguagem que permite que times grandes desenvolvam grandes projetos. Só que as pessoas não falam sobre como reduzir esses projetos.

## Estrutura de projeto

O objetivo desse projeto é ser um projeto para uma venda de garagem.

- `cmd` - aonde as aplicaçoes ficam, início, finalização ou entrypoint do serviço (CLI - flags, API - request/response), presentation layer
- `cmd/sales-api`
- `cmd/internal` - contém a camada de lógica de negócio, tudo o que é reutilizável dentro dos binários, diferentemente da camada de apresentação. O compilador não permite que códigos dentro dessa página sejam importados.
- `cmd/internal/platform` - código fundamental, onde o verdadeiro código de negócio está.
- `cmd/sales-api/internal`

Qual tipo de problema a sua empresa e você estão tentando resolver? Isso é importante porque toda a engenharia é baseada em entender a lógica de negócio. Se você não entender a lógica de negócio, você não entende seus dados nem as mudanças que estão por vir (ser capaz de criar uma arquitetura aberta pra mudança).

Um pacote define o que é uma unidade de código.

## Módulos

0. Limpar os caches do go modules para começar com um ambiente limpo
`go clean -modcache`

1. Inicializar o módulo com um nome único
`go mod init github.comm/larien/service`
O início de um módulo permite ao compilador executar seu programa onde quer que você esteja e mantém todas as dependências anotadas em um só lugar.

A versão 1.15 permite definir aonde os programas serão compilados em uma flag de ambiente com `GOCACHE`.

As ferramentas não são obrigadas a manter a backwards compatibility.

## Logs

É necessário focar desde o começo em como o projeto será mantido. Ou seja, é importante ser capaz de debugar e manter o código desde o começo.

É necessário saber responder:

1. O que seu projeto resolve?
2. Qual é o propósito dos meus logs?

Se houver discrepância entre as respostas dessas perguntas na sua equipe, o código provavelmente estará uma bagunça.

Logs são caros. Você está poluindo o heat com alocações. Eles servem majoritariamente para UI. É necessário tomar cuidado com o que está sendo logado e pode dificultar o debug do seu projeto. Você está usando logs para debugar a aplicação ou para outra coisa (ex.: armazenamento de dados)?

É necessário tratar os logs como algo que as pessoas vão ler.

Bill Kennedy não acredita em níveis de logs.

Nenhum código no nível de platform deve logar. Isso se deve ao fato de que os códigos em platfrom devem ser foundational e pode ser reutilizado. Como cada projeto utiliza sua própria forma de logar, é importante não inserir essa dependência no código de platform.

## Configuração

Só existe uma parte do projeto que pode importar as configurações: `main.go`.

Um desenvolvedor novo deve ser capaz de baixar o projeto e desenvolver com os valores `default`. Logo, os valores definidos como padrão deve funcionar com as dependências do projeto configuradas.

Ctrl+Shift+P - Go: Restart Language Server: para reiniciar o serviço caso ele não funciona corretamente

## Como o go mod tidy funciona

Digamos que um projeto A que é importado utiliza uma dependência D cuja versão mais recente está na 1.9, mas a dependência A utiliza sua versão 1.9. O go mod, ao invés de pegar a versão mais recente, utiliza um algoritmo de MVS (minimum version selection) para obter a versão 1.2 para evitar quebrar a compatibilidade da dependência original.
Caso uma dependência B apareça com a versão D 1.4, o algoritmo selecionará a versão mais recente (greatest but not latest), logo 1.4, ao invés de manter a D 1.2.
Se por ventura a dependência B for removida, a versão maior ainda é mantida.

Existe um serviço no Google que é escalável chamado Module Mirror/Proxy Server. Ele contém todas as versões de todos os módulos gerados até então. Quando uma aplicação requisita um módulo com o go modules, ele verifica com a origem do módulo (ex. GitHub) e atualiza as versões, disponibilizando uma versão compilada para a aplicação.
Caso exista alguma questão de segurança e os pacotes não possam passar pelo Google, é possível utilizar o pacote diretamente com o GitHub com a diretiva `direct` ou crioar seu próprio Module Mirror (ex. Athens).

GOPROXY="https://proxy.golang.org,direct": se o servidor de proxy retornar 401 ou 404 ou algum erro interno com 500, a aplicação procura direto na fonte do código.

É possível definir o servidor privado diretamente na variável de ambiente GOPRIVATE, que copia para GONOPROXY e GOSUMDB. Essa variável representa a ideia de que os links que estiverem nessa lista vão ser verificados antes do servidor de proxy. É importante saber disso para evitar vazamento de dados e para definir o que pode ser compartilhado com o Google e o que não pode.

Existe outro servidor chamado `Checksum DB` que trabalha diretamente com o go.sum. Para cada módulo requisitado no servidor de proxy pela primeira vez. Caso seja a primeira vez requisitando um módulo, a aplicação gera um checksum do pacote gerado. Ela faz, então, uma segunda requisição para esse servidor para verificar se essa hash gerada é a mesma hash que existe no servidor de proxy.
Isso é extremamente importante para segurança no caso de a versão da dependência ter sido alterada (ex. uma segunda versão da mesma versão é lançada). Logo, se o código for diferente o hash será diferente e a aplicação saberá que o código foi alterado através do hash diferente.
Isso acontece somente com os pacotes públicos que passam pelo servidor de proxy e não acontecem com os pacotes privados, já que não há acesso de pacotes privados nesse servidor.
Já deve ser o suficiente atualizar os pacotes, rodar os testes e todos passarem.

`go mod vendor` é uma boa opção para manter o código das dependências utilizadas no caso delas serem deletadas no repositório remoto.

É possível utilizar duas versões diferentes no mesmo pacote se utilizarmos nomes diferentes para cada versão. Então, por exemplo, para usarmos a versão 4.1 e 5.6 de uma dependência, é necessário definir seu nome no go.mod para que ele possa ser importado apropriadamente no projeto. Por exemplo:

- v4.1

```go
module github.com/organizacao/pacote/v4
```

- v5.6

```go
module github.com/organizacao/pacote/v5
```

Logo, no momento de ser importado:

```go
import (
    pacotev4 "github.com/organizacao/pacote/v4"
    pacotev5 "github.com/organizacao/pacote/v5"
)
```


## Definir variáveis em tempo de compilação

É possível alterar valores de variáveis definidas no código através do comando abaixo:

go build -ldflags "-X main.build=$VCS_REF"

Nesse caso, existe uma variável global declarada conforme exemplificado abaixo que é alterada durante a compilação.

```go
var build = "develop"
```

1. A ordem de inicialização da variável global não pode impactar na execução da main
2. A inicialização dessa variável não pode depender do sistema de configuração
3. O único código fonte que pode utilizar essa variável global é o arquivo que a declara

## Métricas

É importante adicionar métricas desde o início e o Go já dá um conjunto de ferramentas nativamente. Isso é disponibilizado através com /debug/vars e /debug/pprofs/ que são definidos no início da aplicação e cada um contém um endpoint diferente para ser possível acessá-los com um servidor em execução.

`expvarmon` gera uma interface com as métricas geradas pelo seu sistema.

## Routers

Não é bom se basear em roteadores que diferem em nanossegundos, e sim o que te dá as funcionalidades que você precisa.

## Shutdown

É importante definir uma forma única de começar e iniciar a apliucação porque hoje em dia cada formato de manter a aplicação utiliza uma forma de fazer ambas as coisas (k8s, docker etc) e manter um entrypoint que essas aplicações possam utilizar pode poupar dores de cabeça em relação a definir uma rotina de finalização de execução por nós mesmos que essas ferramentas devem esperar. Uma forma de definir isso é com timeouts.

Tudo o que é feito com canais é para definir sinais. Os sinais são utilizados para comunicação com o hardware para inicializar ou finalizar a aplicação. Isso é importante porque não queremos perder informações caso a aplicação deva ser finalizada por algum motivo inesperado. Por isso fazemos o clean shutdown: ao receber o sinal de desligamento, o sistema guarda as informações que não foram salvas no contexto para poder voltar a ser executado após o serviço reiniciar.

O contexto é uma forma de compartilhar informações entre as funções e pode mudar de função para função.

A função bloqueante é o canal de shutdown. Quando o shutdown é ativado, a função api.ListenAndServe imediatamente retorna erro e não o contrário. É importante certificar de que ambos estejam rodando em goroutines diferentes caso esse comportamento não seja desejado.

Você não é responsável por desacoplar código de quem o chama. Retorne o tipo concreto e deixe quem está chamando a função definir se a interface será inserida ou não. As exceções são o tratamento de erro e quando lidamos com `reflection`.

## Health

Health é o primeiro endpoint para certificar que o sistema está funcionando como esperado. A ideia principal é chamar esse handler toda vez que determinado endpoint for chamado para ter certeza de que a requisição pode continuar seu processamento, já que o sistema está funcionando corretamente.

## Consistência

É um dos maiores problemas no desenvolvimento e engenharia de software. Manter o mesmo padrão de modelo mental no decorrer do desenvolvimento.
É necessário definir quem é responsável pelo quê e aplicar acordos que respeitem isso.
Ex.: o request precisa ser decodificado, é necessário validar os dados que chegam e a aplicação da lógica de negócio. Request e validação deveriam estar no handler, mas a lógica de negócio deve estar em seu próprio pacote. Além disso, deve haver uma resposta para a request, que deve ser sucesso ou falha. Não é tão bom deixar a responsabilidade de lidar com o erro para quem está escrevendo o handler, nem de logar. Essas não são responsabilidades de handlers.

É necessário encontrar uma forma de estender o Mux existente para lidar com esses casos, ou seja, criar um middleware a partir do mux que já temos.

## Context

O contexto nos dá os seguintes benefícios/possibilidades:
1. Cancelamento e timeouts
2. Possibilidade de debugar e acompanhar as requisições enquanto elas acontecem

No nosso caso atual não é possível inserir o contexto porque o handler espera uma assinatura específica que não o contém. Logo, precisamos expandir o handler para que possa conter nosso contexto. Isso pode ser feito acoplando o handler em uma estrutura:

```go
type App struct {
	*httptreemux.TreetMux
	shutdown chan os.Signal
}
```

Dessa forma, App tem tudo o que ContextMux tem e também o canal de sinal para shutdown.

Criando esse App que contém `httptreemux.TreetMux` é possível trocar o tipo concreto retornado em handlers.API. Um tipo concreto continua sendo retornado, mas o mux utilizado não é vazado para o caller:

```go
func API(build string, shutdown chan os.Signal, log *log.Logger) *web.App { // retorna tipo concreto mesmo assim

	app := web.NewApp(shutdown)

	app.Handle(http.MethodGet, "/test", health)

	return app
}
```

## Middleware

O boilerplate presente no handler deve evitar ter códigas relacionadas a lógica de negócio. Digamos que, por exemplo, um erro é recebido no handler. O que fazer? Não dá para logar lá pelo acordo feito inicialmente. Ignorar é burrice. Como reagir a casos assim dentro do handler?
Inserindo um middleware.

O middleware recebe um handler e retorna um handler. O nosso middleware em questão recebe uma lista de middlewares e aplica cada um deles do último para o primeiro para o handler em questão.

Já que middlewares costumam ser pequenos, é overkill criar um pacote por middleware. Logo, vamos criar um pacote para todas e separá-las em arquivos.

Só que isso tudo é muito bonito para os casos em que todos os handlers utilizam todos os middlewares. Mas e para handlers específicos? Por exemplo, nem todo handler precisa logar. O que fazer?

Utiliza-se clojures para acoplar as dependências do middleware e injetar os middlewares que queremos. Por exemplo:

```go
func Logger(health web.Handler) web.Handler {
	h := func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
		health(ctx, w, r, params)

		// TODO - log

		return nil
	}
	return h
}

```

O problema com essa abordagem é que a função Logger recebe apenas o health que é utilizado para o healthcheck, mas não há uma forma fácil de injetar os dados de log nesse código. O Logger é um middleware, então não dá para injetar um novo parâmetro na função, que deve receber um handler e devolver outro handler.
Uma solução para isso é encapsular esse middleware em uma função que vai apenas disponibilizar o contexto do log. Ou seja, a função em si vai usar clojures para conter o código acima e só vai servir para receber a estrutura do log e disponibilizá-la no contexto do middleware da seguinte forma:

```go
func Logger(log *log.Logger) web.Middleware {
	m := func(health web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
			health(ctx, w, r, params)

			log.Println("I have a logger")

			return nil
		}
		return h
	}
	return m
}
```

Nota: não é recomendado colocar logs ou dados de database no contexto porque as coisas podem dar errado caso eles não estejam lá. Caso seja o caso, é importante lidar com isso. Se não houver dado de log ou de db, há algo muito errado com aquela rotina e é melhor acabar com ela:

```go
v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return web.NewShutdownError("web value missing from context")
	}
```

## Request ID

O request ID é útil para entender o que aconteceu com aquela requisição. Além disso, também é útil para entender o que um usuário fez no decorrer de alguma ação atrelando o request ID ao ID do usuário.

## Erros

Os métodos em Go geram um tipo de pilha. Por padrão, praticamente toda função deve retornar um erro que represente uma falha no meio da sua execução. Retornando um erro, ele vai retornando para as funções da pilha até que alguma função no meio dela lide com ele ou ele chegue na origem da pilha, gerando uma falha de execução.
O problema disso é que até alguma função lidar com esse erro, todo o contexto das funções que o erro já passou é perdido. Uma forma de resolver isso é envolvendo "wrap" o erro nos erros dessas funções. Dessa forma o contexto não se perde.

Criamos um middleware para lidar com erros.

Existem três tipos de erros que podemos começar.
- erros de aplicação, que retornamos para o cliente porque sabemos que é limpo
- erros de fora do nível da aplicação, que geralmente retorna 500 e devemos evitar
- erros de shutdown, porque algo muito ruim está acontecendo

A diferença de colocar esses tipos de coisas (tratamentos de erros, logs) é gerar um contrato entre todas as aplicações que utilizam essas bibliotecas. Se colocarmos esse mesmo tratamento de erro na camada fundamental, todos os projetos são obrigados a seguir esse mesmo formato de tratamento de erros. A parte boa é que todos vão usar a mesma língua, mas a parte ruim é que os projetos ficam atrelados a apenas esse tratamento de erro.

## Autenticação e autorização

Utiliza JWT - JSON Web Token
A autenticação vai validar se esse usuário é válido para o sistema e ele é reconhecido.
A autorização dá acesso ou não ao usuário em relação aos acessos do sistema.

## Para estudar

- Modelos mentais
- MVS - minimum version selection
- expvarmon
- dimfield/httptreemux
- diamond dependency
- github.com/pkg/errors
- jwt.io

## Referências

go please - `github.com/golang/tools/blob/master/gopls/README.md`
ardan conf - `github.com/ardanlabs/conf`