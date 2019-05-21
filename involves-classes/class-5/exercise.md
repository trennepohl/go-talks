## Batalha medieval

Crie uma jogo baseado em turnos onde o jogador enfrentará uma aventura medieval.

O jogo se dará em 5 rounds, no estilo arena. O jogador deverá iniciar o jogo informando o seu nome o estilo de luta do personagem, genero e poderá distribuir 5 pontos de atributos entre força, agilidade, magia e sorte.

O jogo criará randomicamente um oponente para o jogador. Cada round que o jogador vencer o jogo ficará mais difícil.

O jogador contará com um inventário de 2 espaços, onde ele poderá carregar equipamentos que o ajudarão na jornada.

Ao fim de cada round o jogdor será presenteado com um novo item, que será randomicamente escolhido dentro de uma lista de items.

Se o inventario estiver cheio o jogador deverá escolher se um item do inventário deve ser removido ou se ele não deseja o item que lhe foi presenteado.

Ao fim de cada batalha o jogador poderá atribuir 1 novo ponto em um dos atributos.

Durante a batalha o jogador poderá escolher qual ação seu personagem poderá fazer, estas ações incluem:

- Atacar
- Defender
- Usar item
- Fugir

As funções funcionarão da seguinte forma:

#### Guerreiro

O valor de ataque é a soma dos pontos de força com o valor do dano da arma, quando escolhido a opção defesa ele bloqueia 75% do dano.

Pontos de vida: 100 + 10 para cada ponto em força

#### Arqueiro

Ataque: O valor de ataque é a soma dos pontos de agilidade com o valor do dano da arma. Porém o arqueiro tem uma chance de dano crítico conforme a quantidade de pontos em sorte.

Defesa: Caso escolhida esta ação no proximo round o arqueiro tem 30% de chance de desviar do ataque "defendendo" 100% do dano.

Pontos de vida: 100 + 2% de evasão para cada ponto em agilidade

### Mago

o valor de ataque é a soma dos pontos de magia com o valor do dano da arma, e é o que 85 % de chances de ser bem sucedido em uma fuga.

Pontos de vida: 100 + 8% de chance de dano crítico para cada ponto em magia.



## TODO 

- Adicionar armadilhas

- Jogo online


