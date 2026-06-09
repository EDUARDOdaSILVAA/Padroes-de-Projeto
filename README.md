# Padrões de Projeto (Design Patterns)

Instituição: Universidade Tecnológica Federal do Paraná (UTFPR) - Campus Campo Mourão  
Curso:Bacharelado em Ciência da Computação  
Aluno: Eduardo da Silva

Este repositório contém a implementação prática de três Padrões de Projeto (um Criacional, um Estrutural e um Comportamental), desenvolvidos como requisito avaliativo.

---

##  Declarações Obrigatórias

* Direitos Autorais e Referência: Todos os conceitos, fundamentações teóricas e ilustrações baseiam-se no catálogo do [Refactoring Guru](https://refactoring.guru/pt-br/design-patterns).
* Uso de Inteligência Artificial: A estruturação inicial das pastas de arquivos, trechos de códigos e auxílio na sistaxe das inguagens, revisão da documentação, além da criação dos makefiles deste repositório contaram com o auxílio da LLM Google Gemini 3.1 Pro, sendo posteriormente testados, adaptados e organizados para os fins desta entrega. Foi também utilizado o copilot free para correção de erros de importação.

O Factory Method é um padrão criacional que fornece uma interface para criar objetos numa superclasse, mas permite que as subclasses alterem o tipo de objetos que serão criados.

* Contexto: Um sistema corporativo de gestão documental que processa ficheiros PDF e Word.
* Problema: Instanciar objetos diretamente acopla rigidamente o código, enchendo o sistema de condicionais e violando o Princípio Aberto/Fechado.
* Solução: A responsabilidade de instanciar o objeto é delegada para subclasses especialistas (`PdfCreator` e `WordCreator`), interagindo apenas com a abstração.
* Código e Execução: Consulta a pasta `factory-method-java/`.

---



---

## 3. Padrão Comportamental: Strategy (Go)

O Strategy é um padrão comportamental que permite definir uma família de algoritmos, colocá-los em estruturas separadas e tornar os seus objetos intercambiáveis em tempo de execução.

* Contexto: Geração de propostas comerciais para clientes onde o cálculo pode ser feito por Preço Fixo ou Por Hora.
* Problema: Colocar toda a matemática numa única estrutura geraria um código confuso. Adicionar novos modelos exigiria alterar a classe principal.
* Solução: Extraímos os algoritmos para estratégias independentes que assinam uma mesma interface (`PricingStrategy`), permitindo trocar o modelo dinamicamente.
* Código e Execução: Consulta a pasta `strategy-go/`.