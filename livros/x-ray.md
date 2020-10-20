# Software Design X-Rays: Fix Technical Debt with Behavioral Code Analysis

## Capítulo 1: Por que dívida técnica não é técnica

### A taxa de interesse é uma função do tempo

Só porque um código é ruim não quer dizer que é dívida técnica. Não é dívida técnica até termos que prestar atenção nele, e a taxa de interesse é uma função do tempo.
Precisaríamos saber qual a frequência de modificação (e leitura) de cada pedaço de código para separar a dívida que importa para a nossa habilidade de manter o sistema do código que pode ser importante, mas não nos impacta.

### Seus modelos mentais do código

Um dos aspectos mais desafiadores da programação é que precisamos servir duas audiências: a máquina que executa nosso programa, que não se importa com o estilo mas é bem chato em relação ao conteúdo; e os programadores mantendo o código, que tem um processo mental mais elaborado e precisa da nossa ajuda para usar esses processos de forma eficiente. É por isso que focamos em escrever códigos expressivos e bem organizados.

Um esquema é uma construção teórica usada para descrever a forma com que organizamos conhecimento na nossa memória e como usamos esse conhecimento para um evento em particular.

### A qualidade sofre com desenvolvimento paralelo

A estrutura de desenvolvimento de uma organização é um preditor mais forte de defeitos que qualquer métrica de código.
O risco de que um commit em específico introduz um efeito aumenta com o número de desenvolvedores que trabalharam anteriormente no código modificado.

### Mine a inteligência coletiva da sua organização

Como descobrir as áreas que precisam de melhoria? Idealmente, precisamos das seguintes informações:
- Onde está o código com maior taxa de interesse?
- Nossa arquitetura suporte a forma como nosso sistema evolui?
- Existem gargalos de produtividade para coordenação entre os times?
  Nossos dados de controle de versão são uma mina de ouro de informação. Mas é uma mina de ouro que raramente nos aprofundamos.

### Priorize melhorias guiadas por dados

Para melhorar, precisamos priorizar a forma com que realmente trabalhamos com o código, e priorizar dívida técnica requer uma dimensão de tempo na nossa base de código.

Nossa liberdade de código é fortemente restrita se tentarmos refatorar um módulo que está sob desenvolvimento constante com um grupo de programadores comparado com um pedaço de código que trabalha isolado. A não ser que levemos o lado social da nossa base de código em conta, vamos falhar em identificar custos de manutenção significativos.

https://martinfowler.com/bliki/TechnicalDebtQuadrant.html

## Capítulo 2: Identifique código com maior taxa de interesse

### Meça as taxas de interesse

Esses gráficos apresentam uma visão evolucionária de três bases de código diferentes. Organizamos os arquivos de cada base de código de acordo com suas frequências de alteração - ou seja, o número de commits feitos em cada arquivo gravado no controle de versão, com o eixo Y mostrando o número de commits.

[](imagens/1.png)

Eles mostram uma power law distribution.

A distribuição significa que a maioria do nosso código está na cauda longa. É código que raramente é tocado. De forma bem simplista, essa característica sugere que a maioria do nosso código não é importante da perspectiva de custo ou qualidade. Em comparação, a maior parte da atividade de desenvolvimento é focada em uma parte relativamente pequena da base de código. Isso nos dá uma ferramenta para priorizar melhorias, como a figura a seguir ilustra.

[](imagens/2.png)

Na prática, mais frequentemente que o contrário, arquivos com maior frequência de alterações sofrem de problemas de qualidade.

### Uma fórmula para a taxa de interesse

Você só precisa contar o número de vezes que cada arquivo foi referenciado no seu log do Git e organizar os resultados.

```bash
git log --format=format: --name-only | egrep -v '^S' | sort \ | unq -c | sort -r | head -5
```
O `--format=format:` dá uma lista de todos os arquivos já alterados. O `egrep -v` limpa nossos dados removendo linhas em branco do comando anterior e o resto conta as frequências de alteração e mostram os resultados em ordem.

Para salvar tudo em um arquivo, basta colocar `> arquivo.txt` no final.

Os arquivos que atraem a maior parte das alterações são aqueles que são centrais no sistema.

### A eficácia das frequências de alteração

Nossas frequências de alteração nos permite identificar o código que fazemos a maior parte do trabalho e também apontam possíveis problemas de qualidade.
Apesar dessas descobertas, o modelo ainda tem fraquezas, porque todo código é diferente. Há uma diferença enorme em aumentar o número da versão em uma linha e corrigri um bug em um módulo de cinco mil linhas.
A vantagem de usar linhas de código é a simplicidade.

### Priorize dívida técnica com hotspots

Um hotspot é código complicado que precisa ser alterado constantemente. Hotspots são calculados com a combinação de duas métricas que exploramos:
1. Calculando a frequência de alteração de cada arquivo com a taxa de interesse
2. Usando linhas de código como medida simples de complexidade de código

Enclosure diagram

[](imagens/3.png)

Eles crescem com o tamanho da básica de código. É possível interpretar de duas formas:
- Hierárquica: segue a estrutura de pasta da sua base de código. Veja os círculos azuis grandes na figura. Cada um deles representa uma pasta da sua base de código. Círculos aninhadosrepresentam subpastas.
- Interativo: trabalhar com bases de código grandes precisa ser interativo. Isso significa que você pode aumentar a  área de interessee clicando em um dos círculos.

### Avalie os hotspots com tendências de complexidade

Podemos descobrir quão severo um problema em potencial é através da análise de tendências de complexidade, que procura por complexidade acumulada de um arquivo no decorrer do tempo.

[](imagens/4.png)

As linhas de código provavelmente crescem com o tempo, mas a comlexidade de cada linha cresce mais rápido.

### Use os Raios X para se aprofundar no código

Você pode capitalizar nesse aspecto executando uma análise de hotspot no nível d emétodo para identificar segmentos de código que contribuiem para a maioria do arquivo ser um hotspot. Vamos referenciar essa análise como Raio X para distinguir das análises a nível de arquivo.

Um Raio X te dá uma lista priorizada de métodos para inspecionar e, possivelmente, refatorar.

Uma análise Raio X envolve:
1. Obter o código fonte de cada revisão de histórico do nosso hotspot do Git
2. Executar um `git diff` de cada revisão subsequente do código. A saída nos mostra onde - no arquivo histórico - os desenvolvedores fazem modificações.
3. COmparamos os resultados as funções/métodos que existem naquela revisão em particular. Isso significa que precisamos explorar o código fonte para saber quais funções foram afetadas em um commit específico.
4. Fazer o cálculo de hotspot no conjunto resultando das funções alteradas em todas as revisões do hotspot. O algoritmo é idêntico ao usado no nível de arquivo, mas o escopo difere. A frequência de alteração representa o número de vezes que modificamos uma função, e o tamanho da função nos dá a dimensão de complexidade.

[](imagens/5.png)

### Inspecione o código

A vitória com uma análise hotspot é que ela nos permite minimizar os esforços manuais enquanto certificamos alta probabilidade de forcamos nas partes corretas do código.

Casos de acoplamento de controle através de booleados são um problema já que introduzem lógica condicional e diminuem a coesão forçando estado adicional. Tal controle também leva a duplicações de código sutis. Essas escolhas de design não são boas na manutenção.

### Escape da armadilha da dívida técnica

Se as pessoas conseguirem descifrar hieróglicos e a sequência do genoma humano, também deve ser possível entender bases de código legado.

Há vários motivos pelos quais código cresce em hotspots. O motivo mais comum é baixa coesão, o que significa que o hotspot contém diversas partes não relacionadas e falta modularidade. Eles atraem muitos commits porque têm muitas responsabilidades.

## Capítulo 3: Acoplamento no tempo: uma heurística para o conceito de surpresa

### Detecte arquivos que são alterados em conjunto

Com o hierarchical edge bundle:

[](imagens/6.png)

### O segredo do copia-e-cola

Copiar e colar pode não ser o problema se dois pedaços de código evoluem em direções diferentes. Se não o fizerem - se continuarmos modificando ambas as partes - temos um problema.

[](imagens/7.png)

Uma vez que identifiquemos os clones de software que importam, precisamos refatorá-los. A abordagem de refatoração é a extração do padrão repetido em um novo método parametrizando o conceito que varia. Isso torna o código um pouco mais barato de manter enquanto nossa dependência temporal desaparece.

## Capítulo 4: Pague sua dívida técnica

Mesmo com diversas técnicas de refatoração, precisamos considerar o lado das pessoas no código, também.

### Siga o princípio da proximidade

O princípio da proximidade foca em quão bem organizado seu código é em relação a legibilidade e mudança. A proximidade implica que funções que mudam juntas devem ser movidas para perto. A proximidade é tanto um princípio de design e uma heurística para refatorar hotspots para um código que é mais fácil de entender.

Melhoramos o código, como a figura ilustra, movendo os métodos semelhantes para perto uns dos outros.

[](imagens/8.png)

Ao ordenarmos nossas funções e métodos de acordo com nossos padrões de mudança, comunicamos a informação que não é expressa na sintaxe da linguagem de programação.

A principal vantagem de uma refatoração por proximidade é que tem risco baixo. Se você detectar código copia-e-cola no dia de um prazo crítico, pode não ser a hora certa de abstrair a duplicação.

Para manter uma estrutura amigável para as pessoas, você precisa manter essas funções relacionadas próximas umas das outras no seu código fonte.

### Refatore código congestionado com o padrão splinter

O padrão splinter providencia uma forma estruturada de quebrar os hotspots em pedaços gerenciáveis que podem ser divididos entre várias pessoas para que possam trabalhar, ao invés de ter um grupo de desenvolvedores trabalhando em um pedaço grande de código.

O principal motivo pelo qual um pedaço de código cresce em hotspots é porque acumula diversas responsabilidades centrais. Como consequência, o hotspot tem muitas formas de mudar.

### O desenvolvimento em paralelo conflita com a refatoração

O padrão splinter resolve esse dilema reconhecendo que refatorar um hotspot é um processo iterativo que é aplicado em diversas encarnações do código. Em uma refatoração splinter você não precisa melhorar a qualidade de código dessa forma, mas ao invés disso transformar o código em uma estrutura onde divesas pessoas podem trabalhar juntas em paralelo em direção ao objetivo geral de refatoração.

### Separe um arquivo hotspot em responsabilidades

A intenção do padrão splinter é quebrar um hotspot em partes menores em relação a suas responsabilidades enquanto mantém a API original no período de mudança.

Segue as etapas por trás da refatoração splinter iterativa:

[](imagens/9.png)

1. Tenha certeza que seus testes cobrem o candidato a splinter. Se não tiver uma suite de testes adequada - poucos hotspots têm - você precisa criar uma;
2. Identifique os comportamentos dentro do seu hotspot. Essa etapa é um exercício de leitura de código onde você analisa os nomes dos métodos dentro do hotspot e identifica códigos que formem grupos de comportamentos;
3. Refatore por proximidade. Agora você forma grupos de funções com comportamento relacionado dentro do arquivo maior, baseado em comportamentos que você identificou antes. Essa refatoração por proximidade torna sua próxima etapa muito mais fácil;
4. Extraia um novo módulo para o comportamento com maior atividade de desenvolvimento. Use uma análise Raio X para decidir por onde começar, então copie e cole seu grupo de métodos em uma nova classe deixando o original intocado. Lembre de colocar um nome descritivo no seu novo módulo para capturar seu propósito;
5. Delegue para o novo módulo. Substitua o corpo dos métodos originais com delegações para seu novo módulo. Isso permite que você evolua mais rapidamente, o que limita o risco de mudanças conflitarem para outras pessoas;
6. Execute os testes de regressão necessários para certificar que você não tenha alterado o comportamento do sistema. Commite suas mudanças quando os testes passarem;
7. Selecione o próximo comportamento para refatorar e comece na etapa 4. Repita as etapas splinter até que tenha extraído todos os hotspots críticos que identificou na análise Raio X.

A chave para uma refatoração splinter de sucesso é priorizar sua próxima etapa com dados evolucionários, porque não há como refatorar um hotspot principal de uma só vez.

### Saiba as consequências dos splinters

Uma refatoração splinter cria um novo contexto onde você lida com um problema maior quebrando-o em partes menores.

### Crie testes temporários como rede de segurança

Para um hotspot gigante é difícil construir uma suite de testes que cubra boa parte dos casos. Em situações como essa, você precisa construir uma rede de segurança baseada em testes end-to-end. Esses testes focam em capturar cenários performados no nível no sistema. Isso significa que são executados com uma base de dados, conexões de rede, UI, e outros componentes do sistema de verdade.
Os testes end-to-end te dão uma cobertura de testes boa o suficiente para servirem como suite de regressão, e essa suite te permite fazer a refatoração inicial sem quebrar um comportamento fundamental.

O truque é tratar o código como uma caixa preta e só focar no comportamento visível.

Depois, meça a cobertura de teste da sua suite e procure por caminhos de execução que não foram cobertos com alta complexidade. Use essa informação de cobertura como feedback para completar seus testes e grave testes adicionais para cobrir caminhos de execução faltano. Você também pode fazer uma nova mental de extrair esse comportamento em seu próprio módulo splinter.

### Reduza a dívida removendo custos

É comum encontrar hotspots com cobertura de testes inadequada. Isso não significa que não existam testes, só que não existem testes aonde eles precisam estar.
É melhor apagar os testes nesse caso.

### Transforme os métodos do hotspot em pedaços amigáveis para pessoas

Nomear nossos construtores de programação é uma técnica poderosa que lida com o fator mais limitador da programação - nossa memória funcional.
A memória funcional é uma construção cognitiva que serve como espaço de trabalho do seu cérebro. Ela te permite integrar e manipular informação na sua cabeça. A memória funcional também é limitada em recurso e tarefas de programação a levam ao máximo.
Otimizar código para programadores entenderem é uma das escolhas mais importantes que fazemos.

### Pedaços

Mestres de xadrez não se lembram de peças individuais. Eles lembram de padrões, que são representados como grupos de peças. Psicólogos cognitivos chamam esses grupos de chunks, e chunks também funcionam para código.

Quando separar um método hotspot em um grupo de chunks, considere deixar o código como está e siga com uma análise Raio X no seu código refatorado um mês depois. As chances são que a maioria dos seus chunks continuaram estáveis, o que significa que você pode ignorá-los e focar seus esforços de refatoração nas partes que continuam a evoluir.

### A maldição do sistema de sucesso

Escrever código sempre envolve explorar e enetender ambos os domínios do problema e da solução. Logo, é inevitável que sigamos o caminho errado de vez em quando, e a pressão de completar uma funcionalidade torna difícil parar e voltar atrás.

Isso não deve acontecer se prestarmos atenção na nossa base de código e tomar as medidas certas quando necessário.

## Capítulo 5: Os princípios da idade de código

### Estabilize código por idade

Desenvolver com código antigo como um guia significa que:
1. Organizamos nosso código por tempo de existência;
2. Transformamos pacotes estáveis em bibliotecas;
3. Movemos e refatoramos código que falhamos em estabilizar.
Seguir esses princípios nos dá uma série de vantagens:
- Promove modelos de memória a longo prazo;
- Diminuir a carga cognitiva já que há menos código ativo;
- Prioriza suites de testes para diminuir tempo útil.

Para identificar a idade do arquivo, usamos o seguinte comando:

```bash
git log -1 --format="%ad" --date=short \-- <arquivo>
```

[](imagens/10.png)

Obtemos uma lista de todos os arquivos no repositório com sua última modificação e finalmente calculamos a idade de cada arquivo.

### As três gerações de código

Dan North diz que queremos que nosso código ou seja muito recente ou muito antigo, e o tipo de código que é difícil de entender fica entre esses dois extremos.

Tentemos a esquecer em uma taxa exponencial.

A figura a seguir mostra a curva do esquecimento de Ebbinghaus, onde esquecemos rapidamente informações aprendidas desde o primeiro dia. Para reter informaçõs precisamos repeti-la, e com cada repetição somos capazes de melhorar nossa performance lembrando mais.

[](imagens/11.png)

Código recente é o que estendemos e modificamos agora, o que significa que temos um modelo mental fresco do código e sabemos como ele funciona. Em comparação, código antigo é por definição estável, o que significa que não precisamos modificá-lo, nem precisamos manter informação detalhada sobre seu funcionamento interior. É uma caixa preta.
Essa curva também explica por que o código que não é nem antigo nem novo é problemático; tal código é onde esquecemos muitos detalhes e precisamos revisitá-lo de vez em quando. Cada vez que revisitamos esse meio termo precisamos reaprender seu funcionamento, o que se torna um curto de tempo e esforço.
Também há um lado social da idade do código no sentido de que quanto mais antigo o código, mais provável que o programador original tenha deixado a organização. Isso é particularmente problemático para o código intermediário - o código que falhamos em estabilizar - porque significa que, como organização, temos que modificar código que não conhecemos.
A primeira cirurgia ignorante (David Parnas) é um convite para outros fazerem o mesmo. Com o tempo o código se tornar mais difícil de entender, o que nos deixa com dívida técnica que vem do fator organizacional de falha em manter domínio do sistema. Tal código também se torna frágil, o que significa que é importante estabilizar código do ponto de vista de qualidade também.

### Refatore código de idades parecidas

Vamos colocar os números em uma planilha e gerar um histograma como o que vemos nessa figura.

[](imagens/12.png)

A idade do código é uma heurística. Isso significa que os resultados da análise não vão tomar decisões por nós, mas vai nos guiar para perguntar as questões corretas.

O ponto positivo de análises de software evolucionárias é que elas nos dão feedback que nos ajudam a lidar com o gap entre o estado atual do código e onde queremos estar.

## Capítulo 6: Identifique o ponto de inflexão do seu sistema