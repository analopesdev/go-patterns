# 🏭 Padrão de Projeto: Abstract Factory em Go

## 📌 Introdução

O **Abstract Factory** é um padrão de projeto criacional que permite a criação de famílias de objetos relacionados sem especificar suas classes concretas. Esse padrão é útil quando queremos que um sistema seja flexível para suportar várias variantes de produtos sem modificar o código do cliente.

Neste exemplo, implementamos um sistema para uma loja esportiva que vende **tênis e camisetas** de diferentes marcas (**Nike e Adidas**). O Abstract Factory permite que o código do cliente solicite produtos sem precisar saber de qual marca eles são.

---

## 🎯 Problema a ser resolvido

Queremos criar um sistema onde seja possível fabricar **tênis e camisetas** de diferentes marcas sem precisar modificar o código ao adicionar uma nova marca.

📌 **Estrutura do problema:**

```
MARCA A  →  TÊNIS  |  CAMISETA
MARCA B  →  TÊNIS  |  CAMISETA
...
MARCA N  →  TÊNIS  |  CAMISETA
```

A **Factory** deve ser capaz de criar produtos de qualquer marca sem que o código do cliente precise se preocupar com os detalhes internos.

---

## 🏗️ Implementação do Abstract Factory em Go

### 1️⃣ Criamos as **interfaces** para os produtos

Definimos que todo tênis e camiseta precisam ter um **logo** e um **tamanho**:

```go
// Interface para Tênis
package main

type IShoe interface {
    getLogo() string
    getSize() int
}

// Interface para Camiseta
type IShirt interface {
    getLogo() string
    getSize() int
}
```

---

### 2️⃣ Criamos as **estruturas base** para os produtos

```go
// Estrutura base para tênis
type Shoe struct {
    logo string
    size int
}

func (s *Shoe) getLogo() string {
    return s.logo
}

func (s *Shoe) getSize() int {
    return s.size
}

// Estrutura base para camiseta
type Shirt struct {
    logo string
    size int
}

func (s *Shirt) getLogo() string {
    return s.logo
}

func (s *Shirt) getSize() int {
    return s.size
}
```

---

### 3️⃣ Criamos os **produtos concretos** para cada marca

#### 🔹 **Produtos da Nike**

```go
// Tênis da Nike
type NikeShoe struct {
    Shoe
}

// Camiseta da Nike
type NikeShirt struct {
    Shirt
}
```

#### 🔹 **Produtos da Adidas**

```go
// Tênis da Adidas
type AdidasShoe struct {
    Shoe
}

// Camiseta da Adidas
type AdidasShirt struct {
    Shirt
}
```

---

### 4️⃣ Criamos a **fábrica abstrata**

A fábrica deve ser capaz de criar **tênis** e **camisetas** sem se preocupar com a marca.

```go
// Interface para a Abstract Factory
type ISportsFactory interface {
    makeShoe() IShoe
    makeShirt() IShirt
}
```

---

### 5️⃣ Criamos as **fábricas concretas** para cada marca

#### 🔹 **Fábrica da Nike**

```go
// Fábrica da Nike
type NikeFactory struct{}

func (n *NikeFactory) makeShoe() IShoe {
    return &NikeShoe{Shoe{logo: "Nike", size: 42}}
}

func (n *NikeFactory) makeShirt() IShirt {
    return &NikeShirt{Shirt{logo: "Nike", size: 40}}
}
```

#### 🔹 **Fábrica da Adidas**

```go
// Fábrica da Adidas
type AdidasFactory struct{}

func (a *AdidasFactory) makeShoe() IShoe {
    return &AdidasShoe{Shoe{logo: "Adidas", size: 44}}
}

func (a *AdidasFactory) makeShirt() IShirt {
    return &AdidasShirt{Shirt{logo: "Adidas", size: 38}}
}
```

---

### 6️⃣ Criamos uma função para selecionar a fábrica correta

```go
func getSportsFactory(brand string) (ISportsFactory, error) {
    if brand == "nike" {
        return &NikeFactory{}, nil
    }
    if brand == "adidas" {
        return &AdidasFactory{}, nil
    }
    return nil, fmt.Errorf("Marca desconhecida")
}
```

---

### 7️⃣ Criamos o código do cliente

Agora podemos utilizar a **Abstract Factory** para criar produtos sem precisar se preocupar com a marca específica.

```go
func main() {
    nikeFactory, _ := getSportsFactory("nike")
    nikeShoe := nikeFactory.makeShoe()
    nikeShirt := nikeFactory.makeShirt()

    adidasFactory, _ := getSportsFactory("adidas")
    adidasShoe := adidasFactory.makeShoe()
    adidasShirt := adidasFactory.makeShirt()

    fmt.Println("Nike Shoe:", nikeShoe.getLogo(), "Size:", nikeShoe.getSize())
    fmt.Println("Nike Shirt:", nikeShirt.getLogo(), "Size:", nikeShirt.getSize())

    fmt.Println("Adidas Shoe:", adidasShoe.getLogo(), "Size:", adidasShoe.getSize())
    fmt.Println("Adidas Shirt:", adidasShirt.getLogo(), "Size:", adidasShirt.getSize())
}
```

---

## ✅ Benefícios do Abstract Factory

✔ **Desacoplamento**: O código do cliente não precisa saber qual marca está sendo usada.
✔ **Facilidade de manutenção**: Podemos adicionar novas marcas sem modificar o código existente.
✔ **Organização**: O código fica modular e mais fácil de entender.

---

## 🚀 Conclusão

O **Abstract Factory** é um padrão poderoso para criar **famílias de objetos relacionados** de forma flexível e organizada. Neste exemplo, demonstramos como aplicá-lo para fabricar **tênis e camisetas** de diferentes marcas (**Nike e Adidas**), mas esse conceito pode ser aplicado em diversos contextos.

Se quiser adicionar uma nova marca, basta criar uma nova fábrica seguindo o mesmo padrão!