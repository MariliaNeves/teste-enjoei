

## Enunciado

Dada uma lista de números inteiros que representa uma sequência parcial de números de 1 a \( N \) (onde \( N \) é o maior número na lista), escreva um algoritmo em Go que calcule a soma dos números que estão faltando na sequência.

### Especificações

1. O algoritmo deve receber um slice de inteiros não negativos.
2. A lista pode conter números repetidos, mas esses números devem ser considerados apenas uma vez.
3. O algoritmo deve calcular a soma dos números que estão faltando na sequência de 1 a \( N \).
4. Se todos os números de 1 a \( N \) estiverem presentes, a soma dos números faltantes deve ser 0.

### Exemplo

Considere o slice `[]int{1, 2, 4, 6}`. A sequência completa de 1 a 6 é `1, 2, 3, 4, 5, 6`. Os números que estão faltando são `3` e `5`.

