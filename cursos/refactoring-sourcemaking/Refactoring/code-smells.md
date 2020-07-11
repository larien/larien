# Code Smells

Code smells são sinais claros de que uma refatoração é necessária. A falta da refatoração pode levar à paralização completa de um projeto, desperdiçando alguns anos de desenvolvimento ou recriação dele todo.
Logo, é bom se livrar deles quando ainda estão pequenos.

## Bloaters

São códigos, métodos e classes que aumentaram em proporções tão grandes que são difíceis de se trabalhar com.

- [Método longo](bloaters/long-method.md)
- [Classe ou pacote grande](bloaters/large-class.md)
- [Obsessão por primitivas](bloaters/primitive-obsession.md)
- [Lista de parâmetros longa](bloaters/long-parameter-list.md)
- [Falta de encapsulamento de dados](bloaters/data-clumps.md)

## Aplicação incorreta da orientação a objeto

Tudo isso acontece pela aplicação incorreta dos princípios da orientação a objeto.

- [Declarações switch](object-orientation-abusers/switch-statements.md)
- [Campo temporário](object-orientation-abusers/temporary-field.md)
- [Herança desnecessária](object-orientation-abusers/refused-bequest.md)
- [Classes alternativas com interfaces diferentes](object-orientation-abusers/alternative-classes-with-different-interfaces.md)

## Manutenibilidade ruim

Acontece quando você precisa mudar uma coisa no seu código e tem que fazer mudanças em outros lugares. Isso torna o desenvolvimento muito mais complicado e caro.

- [Mudança que atinge vários lugares](change-preventers/divergent-change.md)
- [Tiro de espingarda](change-preventers/shotgun-surgery.md)
- [Hierarquias de herança paralelas](change-preventers/parallel-inheritance-hierarchies.md)

## Dispensáveis

É algo inútil cuja remoção torna o código mais limpo, mais eficiente e mais fácil de entender.

- [Comentários](dispensables/comments.md)
- [Código duplicado](dispensables/duplicated-code.md)
