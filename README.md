# Aprendizado em GO, durante 100 dias!

## Este é um repositório para acompanhar minha jornada de aprendizado na linguagem Go

day 1 - Comecei a aprender go através de videos aulas do canal [Aprenda Golang](https://youtu.be/bOlnyWOjVIo?si=kdFFeu08PQa5Leu1) - assisti apenas a primeira aula

day 2 - Me cadastrei no site [Exercism](https://exercism.org) escolhi a trilha Go, fiz o 'Hello World' e o primeiro exercício básico, comecei a usar funções

day 3 - Comecei a aprender operações aritméticas, tipos da variável número e converter tipos, ex int -> float, usando funções 

day 4 - Aprendi, através da plataform [Exercism](https://exercism.org) operações lógicas utilizando AND, NOT, OR (&&, !, ||) respectivamente, e como trabalhar com variáveis do tipo booleano (bool), bem como funções que recebem como parâmetros e retornam
este tipo de dado. Aprendi também a comentar meu código GO utilizando as regrás da linguagem.

day 5 - Continuando meu aprendizado no [Exercism](https://exercism.org) agora sobre Strings, [pacote strings](https://pkg.go.dev/strings#pkg-functions) ; Concatenação e tratamento utilizando algumas funções como .ReplaceAll, .TrimSpace, ToUpper, pra fazer precisei utilizar a documentação. 

day 6 - Comecei o assunto sobre "conditionais If", que envolve if, else, if else. Li a teoria e fiz exercicios.

day 7 - Estou muito feliz de ter aprendido, de forma definitiva, o conceito de packages e imports no GO, dessa forma pude organizar os diretorios de um pequeno programa
onde eu pude criar pacotes e importar as funções. 
__obs__: Porém é preciso dizer que packages criados devem ser transferidos onde aponta o GOPATH, no meu caso _/usr/local/go/src_ para que ele o programa funcione.
Também pude me aprofundar nas funções do pacote __fmt__ Sprintf principalmente, ela me permite fazer interpolação de strings e formatar números decimais e float.
ex:
```
fmt.Sprintf("03%d",7)
=> 007
```
O código acima informa que iremos formatar um decimal __%d__ adicionando _zeros_ a esquerda e que o número terá no mínimo 3 dígitos.
__obs__: A função _Sprintf()_ não escreve nada na tela, ela apenas gera um retorno.
