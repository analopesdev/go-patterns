# 🚗 Problema do Projeto

## 📌 Contexto
O projeto visa a criação de um sistema para fabricar carros de diferentes marcas (**Fiat** e **Ford**). No entanto, queremos evitar que o código cliente dependa diretamente das classes concretas das marcas, garantindo maior flexibilidade e expansibilidade.

## ❌ O Problema
Sem um padrão de projeto adequado, o código cliente precisaria instanciar diretamente os objetos das marcas, o que geraria **alto acoplamento**. Isso significa que, ao adicionar uma nova marca, precisaríamos modificar o código existente, tornando-o difícil de manter e escalar.

## 🎯 Solução: Uso do Padrão Abstract Factory
O uso do padrão **Abstract Factory** permite criar fábricas especializadas para cada marca, garantindo que o código cliente possa criar carros **sem conhecer suas implementações concretas**. Isso facilita a manutenção e permite a adição de novas marcas sem alterar o funcionamento do sistema.

## 🏗️ Implementação

### 🔹 Definição da Interface do Carro
Criamos uma interface **Carro**, que define um método `GetMarca()`, garantindo que todas as marcas implementem essa função.

```go
// Interface do Carro
package main

type Carro interface {
    GetMarca() string
}
```

### 🔹 Implementação dos Carros Fiat e Ford
Cada marca implementa a interface **Carro**:

```go
// Implementação do Carro da Fiat
type Fiat struct{}
func (f Fiat) GetMarca() string { return "Fiat" }

// Implementação do Carro da Ford
type Ford struct{}
func (f Ford) GetMarca() string { return "Ford" }
```

### 🔹 Definição da Interface da Fábrica
Criamos a interface **CarFactory**, que define o método `CriarCarro()` para fabricar carros.

```go
// Interface da Fábrica
type CarFactory interface {
    CriarCarro() Carro
}
```

### 🔹 Implementação das Fábricas Fiat e Ford
Cada fábrica retorna um carro da marca correspondente:

```go
// Fábrica da Fiat
type FiatFactory struct{}
func (f FiatFactory) CriarCarro() Carro { return Fiat{} }

// Fábrica da Ford
type FordFactory struct{}
func (f FordFactory) CriarCarro() Carro { return Ford{} }
```

### 🔹 Função para Obter a Fábrica
Criamos uma função que retorna a fábrica correta de acordo com a marca escolhida:

```go
func ObterFabrica(marca string) CarFactory {
    if marca == "fiat" {
        return FiatFactory{}
    } else if marca == "ford" {
        return FordFactory{}
    }
    return nil
}
```

### 🔹 Código Cliente (main.go)
O código principal obtém a fábrica desejada e cria os carros sem depender das implementações concretas.

```go
package main

import "fmt"

func main() {
    fiatFactory := ObterFabrica("fiat")
    carroFiat := fiatFactory.CriarCarro()
    fmt.Println("Carro criado:", carroFiat.GetMarca()) // 🏎️ Fiat

    fordFactory := ObterFabrica("ford")
    carroFord := fordFactory.CriarCarro()
    fmt.Println("Carro criado:", carroFord.GetMarca()) // 🚗 Ford
}
```

## 🏆 Benefícios do Abstract Factory
✅ **Código desacoplado**: O cliente não precisa saber qual classe concreta está sendo usada.  
✅ **Fácil expansão**: Podemos adicionar novas marcas sem modificar o código existente.  
✅ **Organização e reutilização**: Cada classe tem uma responsabilidade clara.  

## 🚀 Como Executar
1. Clone o repositório ou copie os arquivos.
2. Execute o comando no terminal:
   ```sh
   go run main.go
   ```
3. A saída esperada será:
   ```sh
   Carro criado: Fiat
   Carro criado: Ford
   ```

## 📚 Conclusão
Este exemplo mostrou como **o Abstract Factory permite criar famílias de objetos relacionados** sem depender diretamente de suas implementações concretas. Isso facilita a manutenção e escalabilidade do código.  

Quer expandir o projeto? Adicione uma nova marca de carro criando uma nova fábrica! 🚗💨

