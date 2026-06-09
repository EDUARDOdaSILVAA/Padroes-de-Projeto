# Padrões de Projeto (Design Patterns)

Instituição: Universidade Tecnológica Federal do Paraná (UTFPR) - Campus Campo Mourão  
Curso: Bacharelado em Ciência da Computação  
Aluno: Eduardo  

Este repositório contém a implementação prática de três Padrões de Projeto (um Criacional, um Estrutural e um Comportamental), desenvolvidos como requisito avaliativo.

---

## Declarações Obrigatórias

* Direitos Autorais e Referência: Todos os conceitos, fundamentações teóricas e ilustrações baseiam-se no catálogo do [Refactoring Guru](https://refactoring.guru/pt-br/design-patterns).
* Uso de Inteligência Artificial: A estruturação inicial dos códigos de exemplo e a revisão da documentação deste repositório contaram com o auxílio da LLM Google Gemini, sendo posteriormente testados, adaptados e organizados para os fins desta entrega.

---

## 1. Padrão Criacional: Factory Method (Java)

O Factory Method é um padrão criacional que fornece uma interface para criar objetos em uma superclasse, mas permite que as subclasses alterem o tipo de objetos que serão criados.

### Contexto
Imagine um sistema de gestão documental para um ambiente corporativo (ou para os clientes de uma Empresa Júnior, por exemplo). Inicialmente, o sistema precisa processar e abrir apenas documentos em PDF. Com a evolução do projeto, surge a necessidade de suportar documentos do Word.

### O Problema
Se instanciarmos os objetos diretamente no código principal usando `new PdfDocument()` ou `new WordDocument()`, o código ficará rigidamente acoplado a essas classes específicas. Adicionar um novo formato no futuro exigiria modificar todo o código que interage com os documentos, enchendo o sistema de condicionais (`if/else`) e violando o Princípio Aberto/Fechado (Open/Closed Principle).

### A Solução
O padrão Factory Method sugere substituir as chamadas diretas de construção de objetos por chamadas a um método fábrica especial. 
Criamos uma classe criadora abstrata (`DocumentCreator`) que delega a responsabilidade de instanciar o objeto para suas subclasses (`PdfCreator` e `WordCreator`). O código cliente passa a interagir apenas com as abstrações, sem saber os detalhes de qual documento específico está sendo criado.

### Explicação do Código

A implementação em Java (localizada na pasta `factory-method-java/src/main/java/com/`) está dividida nas seguintes partes:

1. A Abstração (Produto): A interface `Document.java` define as operações padrão (`open()` e `close()`) que todos os documentos devem ter.
2. As Implementações Concretas: As classes `PdfDocument.java` e `WordDocument.java` implementam a interface `Document`, cada uma com a sua lógica específica para abrir e fechar ficheiros.
3. O Criador (Fábrica Base): A classe abstrata `DocumentCreator.java` declara o método de fábrica `createDocument()`. Ela também contém a lógica de negócio principal (`manageDocument()`), que usa o documento gerado sem se importar com o tipo dele.
4. Os Criadores Concretos: As classes `PdfCreator.java` e `WordCreator.java` estendem o criador base e sobrescrevem o método de fábrica para retornar, respectivamente, um `PdfDocument` ou um `WordDocument`.
5. O Cliente: Em `Main.java`, simulamos a utilização. Configuramos o criador desejado em tempo de execução e chamamos o processamento, demonstrando que o código flui independentemente do tipo de documento.

---

## 2. Padrão Estrutural: Adapter (C++)
*(Em desenvolvimento - será atualizado em breve)*

---

## 3. Padrão Comportamental: Strategy (Go)
*(Em desenvolvimento - será atualizado em breve)*