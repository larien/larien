# As Leis Fundamentais do Projeto de Software

## Prefácio

A diferença entre um mau e um bom programador é o entendimento.

## 1. Introdução

Programar tem que se tornar o ato de reduzir complexidade à simplicidade.

As partes complexas de um programa têm que ser organizadas de algum modo simples para que um programador possa trabalhar nele sem ter capacidades mentais divinas.

Algumas pessoas acreditam que escrever de modo simples leva mais tempo do que escrever rapidamente algo que "faz o serviço". Na verdade, gastar um pouco mais de tempo escrevendo códgio simples revela-se mais rápido do que escrever muito código rapidamente no início e então gastar muito tempo tentando entendê-lo mais tarde.

Muitos programas ótimos tiveram seu desenvolvimento interrompido no decorrer dos anos só porque levou muito tempo para adicionar recursos às complexas bestas que elas haviam se tornado.

## 2. A ciência ausente

A ciência de projeto de software ajuda pessoas a tomar decisões como:
- Qual deveria ser a estrutura do código de nosso programa?
- É mais importante focar em ter um programa rápido ou um programa cujo código seja de fácil leitura?
- Para nossas necessidades, qual linguagem de programação deveríamos usar?

Projeto de software não aborda assuntos como:
- Qual deveria ser a estrutura da empresa?
- Quando deveríamos ter reuniões de equipe?
- Quais horas do dia os programadores deveriam trabalhar?
- Como deveríamos avaliar o desempenho de nossos programadores?

## 3. As forças motoras do projeto de sotware

Há um único propósito para todo software: ajudar as pessoas.

Nossos objetivos como desenvolvedores são:
- Permitir que escrevamos software que seja o mais útil possível
- Permitir que nosso  software continue a ser o mais útil possível
- Projetar sistemas que possam ser criados e submetidos à manutenção o mais facilmente possível por seus programadores, para que eles possam ser - e continuem a ser - o mais úteis possível

## 4. O futuro

A equação do projeto de software é:

D = V/E
Aonde:
D = desejabilidade por uma alteração
V = valor de uma alteração
E = esforço envolvido
Logo:
A desejabilidade por qualquer alteração é diretamente proporcional ao valor da alteração e inversamente proporcional ao esforço envolvido em realizar a alteração.

## 5. Alteração

| | Arquivo 1 | Arquivo 2 | Arquivo 3 | Arquivo 4 |
| --- | --- | --- | --- | --- |
| Período analisado | 5 anos, 2 meses | 8 anos, 3 meses | 13 anos, 3 meses | 13 anos, 4 meses |
| Linhas originais | 423 | 192 | 227 | 309 |
| Linhas inalteradas | 271 | 101 | 4 | 8 |
| Linhas agora | 664 | 948 | 388 | 414 |
| Cresceram | 241 | 756 | 161 | 105 |
| Vezes alteradas | 47 | 99 | 194 | 459 |
| Linhas adicionadas | 396 | 1026 | 913 | 3828 |
| Linhas eliminadas | 155 | 270 | 752 | 3723 |
| Linhas modificadas | 124 | 413 | 1382 | 3556 |
| Total de alterações | 675 | 1709 | 3047 | 11107 |
| Razão de alteração | 1.6x | 8.9x | 13x | 36x |

Em que:
- Período analisado: período de tempo no decorrer do qual o arquivo existiu
- Linhas originais: quantas linhas havia no arquivo quando ele foi originalmente escrito
- Linhas inalteradas: quantas linhas são as mesmas agora em relação às linhas quando o arquivo foi originalmente escrito
- Linhas agora: quantas linhas há no arquivo agora, no final do período de análise
- Cresceram: a diferença entre "linhas agora" e "linhas originais"
- Vezes alteradas: o número total de vezes que um programador fez um conjunto de alterações no arquivo. Normalmente, um conjunto de alterações representará a correção de um erro, um novo recurso etc.
- Linhas adicionadas: quantas vezes, no histórico do arquivo, uma nova linha foi adicionada
- Linhas eliminadas: quantas vezes, no histórico do arquivo, uma linha existente foi eliminada
- Linhas modificadas: quantas vezes, no histórico do arquivo, uma linha existente foi alterada (mas não recém-adicionada ou eliminada)
- Total de alterações: a soma das "linhas adicionadas", "linhas eliminadas" e "linhas modificadas" para cada arquivo
- Razão de alteração: o quanto "total de alterações" é maior do que "linhas originais"

Há três grandes erros que os projetistas de software cometem quando tentam confrontar a Lei da Alteração:
1. Escrever código que não é necessário
2. Não tornar o código fácil de alterar
3. Ser genérico demais

## 6. Erros e projeto

A chance de introduzir um erro em seu programa é proporcional à quantidade de alterações que você faz nele.

O melhor projeto é o que permite a máxima alteração no ambiente com a mínima alteração no software.

Nunca "corrija" nada a menos que seja um problema e você tenha evidência mostrando que o problema realmente existe.

## 7. Simplicidade

A facilidade de manutenção de qualquer software é proporcional à simplicidade de suas partes individuais.

É importante ter seções na documentação de seu código como "Novo Nesse Código?". Essas devem ser escritas como se o leitor nada soubesse sobre o programa.

Muitos programadores supõe que outros programadores estarão dispostos a passar muito tempo aprendendo tudo sobre seu código porque, afinal de contas, levou muito tempo escrevê-lo! O código é importante para eles, então,  não será importante para todo mundo?

Não é uma questão de inteligência - é uma questão de conhecimento. Programadores que são novos ao seu código nada sabem sobre ele; eles têm que aprender. Quanto mais fácil você tornar para eles aprenderem, mais rápido eles vão entendê-lo e mais fácil será para eles o usarem.

Há muitos modos de tornar seu código fácil de aprneder: documentação simples, projeto simples, tutoriais passo a passo etc.

Mas, se seu código não for extremamente simples de aprender, as pessoas vão usá-lo incorretamente, criar erros e geralmente bagunçar as coisas.

Quando você torna seu produto ou código extremamente simples, você está permitindo que as pessoas o entendam. Isso as faz sentir-se inteligentes, permite que elas façam o que estão tentando fazer e absolutamente não relete de forma negativa sobre você. Na verdade, as pessoas provavelmente o admirarão mais se você tornar as coisas simples do que se você as tornar complexas.

## 8. Complexidade

Coisas que aumentam a complexidade de um projeto:
- Expandir o propósito do software
- Acrescentar programadores
- Alterar coisas que não precisam ser alteradas
- Ficar preso a tecnologias ruins
- Mal-entendido
- Projeto ruim ou nenhum projeto
- Reinventando a roda

Se alguém chega para você e diz algo como: "Como faço esse pônei voar até a lua?, a pergunta que você precisa fazer é: "Qual problema você está tentando resolver?"

Descarte suas suposições. Realmente olhe para o problema que você está tentando resolver. Certifique-se de que você entende plenamente todos os aspectos dele e então descubra o modo mais simples de resolvê-lo.

Pergunte-se: "Como, em geral, em um mundo perfeito, esse tipo de problema deve ser resolvido?"

A maioria dos problemas difíceis de projeto pode ser resolvida simplesmente desenhando ou escrevendo em um papel.

Alguns projetistas, quando confrontados com um sistema muito complexo, jogam-no fora e começam tudo de novo. Contudo, reescrever um sistema a partir do zero é essencialmente uma admissão de fracasso como projetista.

É possível projetar um sistema que nunca precise ser jogado fora.

Você pode apenas reescrever um projeto se todos os seguintes pontos forem verdadeiros:
1. Você desenvolveu uma estimativa precisa que mostra que reescrever o sistema será um uso mais eficiente do tempo do que reprojetar o sistema existente.
2. Você tem muito tempo para passa criando um novo sistema.
3. Você é de algum modo um projetista melhor do que o projetista original do sistema.
4. Você pretende projetar esse novo sistema em uma série de passos simples e ter usuários que podem dar-lhe feedback para cada passo ao longo do caminho.
5. Você tem os recursos disponíveis tanto para fazer manutenção no sistema existente quanto para projetar um novo ao mesmo tempo.

## 9. Testes

O grau em que você sabe como seu software se comporta é o grau em que você o testou com precisão.

## A. As leis do projeto de software
## B. Fatos, leis, regras e definições