# Aprendizado em GO, durante 100 dias!

## Este é um repositório para acompanhar minha jornada de aprendizado na linguagem Go

### Nível - Básico

__day 1__ - Comecei a aprender go através de videos aulas do canal [Aprenda Golang](https://youtu.be/bOlnyWOjVIo?si=kdFFeu08PQa5Leu1) - assisti apenas a primeira aula

__day 2__ - Me cadastrei no site [Exercism](https://exercism.org) escolhi a trilha Go, fiz o 'Hello World' e o primeiro exercício básico, comecei a usar funções

__day 3__ - Comecei a aprender operações aritméticas, tipos da variável número e converter tipos, ex int -> float, usando funções 

__day 4__ - Aprendi, através da plataform [Exercism](https://exercism.org) operações lógicas utilizando AND, NOT, OR (&&, !, ||) respectivamente, e como trabalhar com variáveis do tipo booleano (bool), bem como funções que recebem como parâmetros e retornam
este tipo de dado. Aprendi também a comentar meu código GO utilizando as regrás da linguagem.

__day 5__ - Continuando meu aprendizado no [Exercism](https://exercism.org) agora sobre Strings, [pacote strings](https://pkg.go.dev/strings#pkg-functions) ; Concatenação e tratamento utilizando algumas funções como .ReplaceAll, .TrimSpace, ToUpper, pra fazer precisei utilizar a documentação. 

__day 6__ - Comecei o assunto sobre "conditionais If", que envolve if, else, if else. Li a teoria e fiz exercicios.

__day 7__ - Estou muito feliz de ter aprendido, de forma definitiva, o conceito de packages e imports no GO, dessa forma pude organizar os diretorios de um pequeno programa
onde eu pude criar pacotes e importar as funções. 
_obs_: Porém é preciso dizer que packages criados devem ser transferidos onde aponta o GOPATH, no meu caso _/usr/local/go/src_ para que ele o programa funcione.
Também pude me aprofundar nas funções do pacote __fmt__ Sprintf principalmente, ela me permite fazer interpolação de strings e formatar números decimais e float.
ex:
```
fmt.Sprintf("03%d",7)
=> 007
```
O código acima informa que iremos formatar um decimal __%d__ adicionando _zeros_ a esquerda e que o número terá no mínimo 3 dígitos.
_obs_: A função _Sprintf()_ não escreve nada na tela, ela apenas gera um retorno.

__day 8__ - Comecei a aprender e avancei muito bem acerca do conhecimento de _slices_ e _variadic functions_, espero que no day 9 eu possa concluir este assunto.

__day 9__ - Concluído o assunto de _slices_ e _variadic functions_

__day 10__ - Aprendido e praticado _for loops_ (um assunto que precisa ser mais praticado), resolvi refatorar meu exercício que está no diretório _/exercism/for_loos_ mas deixei meu código antigo comentado para que, em consultas futuras eu possa comparar a evolução.

__day 11__ - Revisão sobre _for loops_ e _slices_ e visto alguma coisa sobre _arrays_. Concluído assunto sobre _Randomness_ e _Conditionals Switch_

__day 12__ - Revisão sobre _for loops_ é um pouco dificil adaptar-se a um linguagem que o único laço de repetição é o _for_. Neste sentido, revisei mais um pouco, e li algumas coisas a respeito da linguagem GO, como nasceu, de quais linguagens herdou, qual o propósito, etc.
Implantei o algoritmo de árvore binária em _./algoritmos/arvore_binaria/main.go_

__day 13__ - Revisto conceitos de __array__, __slices__, __loops__. Aprendendo um novo conceito __map__. Fiz alguns exercícios básicos para treinar o conceito de __map__. Falta refatorar e organizar o código em _/praticas/maps_

__day 14__ - Fiz uma lista de exercício com 5 questões de nível fácil criada pelo chatgpt sobre __maps__ em __Go__. São 5 funções correspondentes a esta lista, que se encontra no pacote _chatgptlistafacil1_ em /praticas/maps/chatgpt_lista_facil1.
Concluído o curso Learn Go: Loops, Arrays, Maps, and Structs da plataforma [codecademy](https://www.codecademy.com/learn)

__day 15__ - Fiz as primeiras 3 questões do livro _A linguagem de programação Go_ sobre laços _for_, _range_ e _join_ este último um conceito novo que aprendi , o método _join_ faz parte do pacote _strings_ e é utilizado para concatenar strings usando um separador _sep_ informado. Também comecei aprender sobre o pacote _time_ que é bem interessante para medir o tempo decorrido entre cada programa. Precisei dar um passo atrás e aprender com cuidado e lentidão _structs_ encontrei uma certa dificuldade e por isso estou fazendo esta escolha.

__day 16__ - Continuando a leitura do livro _A linguagem de programação Go_ aprendi coisas interessantes de como ler da entrada padrão e tratar as mensagens digitadas pelo usuário, por exemplo, criar um programa que mostra na tela as linhas adjacentes repetidas, incrementando o programa, aprendi que posso ter tanto da entrada padrão (os.Stdin) quanto de arquivos (os.Open(<nome-do-arquivo>)). OBS.: É preciso abrir e fechar o arquivo os.Close()

__day 17__ - Ainda estou no capítulo 1, referente ao livro _A linguagem de programação Go_ fazendo os exercícios e tentando entendê-los minimamente, hoje eu modifiquei um programa que cria um gif e aprendi um pouco sobre o pacote net/http fazendo um pequeno programa que consulta URLs passadas na linha de comando e mostra seu código fonte, muito semelhante ao que o __curl__ faz.

__day 18__ - Cap 1 - _A linguagem GO_ exercicio concluido a respeito de um programa que le urls da entrada padrao passadas como argumentos e retorna o tempo de pesquisa e o tempo total. Ideia sobre _goroutines_ e _channels_.

__day 19__ - Cap 1 - _A linguagem GO_ tendo uma ideia geral das pacotes __net/http__ para criar servidores web que tratam requisições http, eco e mostra cabeçalhos. Falta ainda fazer o exercício deste tópico (Exercício 1.12).

__day 20__ - Cap 1 - _A linguagem GO_ __concluído__.

__day 21__ - Entrei no Cap 2 do livro, onde puder aprender sobre nomes reservados em GO e como declarar variaveis. Voltei para a plataforma exercism e fiz um exercicio que envolve funcoes com unico e multiplos retornos.

__day 22__ - [Exercism](https://exercism.org/tracks/go) - Exercício sobre _maps_ concluído, 10% do curso concluído.
