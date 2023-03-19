# Elementos

## Modelo de Dados

### Base Element

```json
{
 "id": "id",
 "name": "Name",
 "category": "FlowObjects/ConnectingObjects/SwimLanes/Artifacts",
 "type": "Event/Activities/Gateways/SequenceFlow/MessageFlow/Association/Pool/Lane/DarkPool/DataObject/Group/Annotation",
 "subType": "Start/Intermediate/End",
 "incoming": [
  "flow_id_1",
  "flow_id_2"
 ],
 "outgoing": [
  "flow_id_3"
 ],
 "catching": true,
 "catchDefinition": {
  "message": "CatchMessage"
 },
 "throwing": true,
 "throwDefinition": {
  "signal": "ThrowMessage"
 },
 "catchs": [
  "catch_id_1",
  "catch_id_2"
 ],
 "throws": [
  "throw_id_1"
 ],
 "refs": [
  "id_1",
  "id_2"
 ],
 "triggeredBy": "id_6",
 "timestamp": {
  "created": "2021-10-23T12:00:00Z",
  "updated": "2021-10-24T13:30:00Z",
  "executed": "2021-10-25T15:00:00Z",
  "triggered": "2021-10-26T10:30:00Z"
 },
 "version": {
  "major": 1,
  "minor": 0,
  "patch": 0
 },
 "state": "Created/Pending/InProgress/Completed/Canceled/Approved/Rejected/Running/Failed/Paused/ScheduledRetry/ExecutedRetry",
 "stateReason": "",
 "attempts": 0,
 "maxAttempts": 3,
 "retryDelay": 30,
 "transitionCondition": [
  {
   "id": "transition_id_1",
   "name": "Transition 1",
   "from": "state_id_1",
   "to": "state_id_2",
   "transitionCondition": {
    "operator": "<=",
    "operand": "10"
   },
   "expression": "this.status == 'approved'"
  }
 ]
}

```

Este código é um exemplo de um objeto em JSON que pode ser usado para representar uma instância de atividade ou evento em um fluxo de trabalho (workflow). Cada chave representa uma propriedade dessa instância e seus valores são as respectivas informações está na explicação de cada chave:

- `"id"`: ID único do elemento.
- `"name"`: Nome do elemento.
- `"category"`: Categoria do elemento, que pode ser "FlowObjects" (objetos de fluxo), "ConnectingObjects" (objetos que conectam elementos de fluxo), "SwimLanes" (pistas) ou "Artifacts" (artefatos).
- `"type"`: Tipo do elemento, que varia dependendo de qual categoria ele pertence. Por exemplo, se o valor da categoria for "FlowObjects", o valor possível deste campo poderá ser "Event" (evento), "Activities" (atividades), "Gateways" (portões) ou qualquer outro tipo relacionado ao fluxo de processos.
- `"subType"`: Subtipo do elemento. Dependendo do tipo especificado no campo acima, essa propriedade pode ter valores como "Start" (início), "Intermediate" (intermediário) ou "End" (fim).
- `"incoming"`: Lista de IDs das conexões de entrada para o elemento.
- `"outgoing"`: Lista de IDs das conexões de saída do elemento.
- `"catching"`: Define se o elemento esta capturando alguma mensagem/evento.
- `"catchDefinition"`: Definição da mensagem/evento que está sendo capturado quando a propriedade "catching" for verdadeira.
- `"throwing"`: Define se o elemento esta lançando alguma mensagem/evento.
- `"throwDefinition"`: Definição da mensagem/evento que está sendo lançado quando a propriedade "throwing" for verdadeira.
- `"catchs"`: Lista de IDs dos elementos que capturam mensagens/eventos desse elemento.
- `"throws"`: Lista de IDs dos elementos que lançam mensagens/eventos a partir deste elemento.
- `"refs"`: Lista de IDs para elementos que este elemento refere.
- `"triggeredBy"`: ID de um elemento que disparou este elemento.
- `"timestamp"`: Registro dos timestamps de criação, atualização, execução e acionamento do elemento.
- `"version"`: Número de versão do elemento.
- `"state"`: Estado atual do elemento, que pode ser alterado ao longo de sua vida útil.
- `"stateReason"`: Descrição do motivo pelo qual o estado atual do elemento foi definido (opcional).
- `"attempts"`: Número de tentativas para executar o elemento.
- `"maxAttempts"`: Número máximo de tentativas permitidas para executar o elemento.
- `"retryDelay"`: Intervado de tempo em segundos para uma nova tentativa ser realizada.
- `"transitionCondition"`: Array que contem os IDs das transições que devem ser executadas quando o elemento for executado.

Existem várias categorias de elementos que são utilizados para representar diferentes tipos de conceitos. A categoria de Flow Objects, podemos ter Event, Activities e Gateways. Estes elementos são usados principalmente para definir o fluxo dos processos de negócios.

Dentro da categoria de Event, podemos ter os seguintes sub-elementos:

- `Start Event`: representa algo que desencadeia um processo de negócio;
- `Intermediate Event`: representa algo que ocorre durante o processo de negócio;
- `End Event`: representa o final do processo de negócio.

Já dentro da categoria de Activities, podemos ter os seguintes sub-elementos:

- `Task`: representa uma atividade que precisa ser realizada;
- `Sub-process`: representa um processo que pode ser dividido em sub-processos menores;
- `Call Activity`: representa uma atividade que chama outro processo ou sub-processo.

Por fim, na categoria de Gateways, temos os seguintes sub-elementos:

- `Exclusive Gateway`: representação de uma condição exclusiva para cada caminho;
- `Inclusive Gateway`: representação de múltiplas condições verdadeiras para passar por ele;
- `Parallel Gateway`: permite que várias sequências possam ocorrer ao mesmo tempo, criando forks e joins.

Na categoria Connecting Objects temos SequenceFlow, MessageFlow, Association.

- `SequenceFlow` é usado para conectar dois objetos de fluxo e sempre
   tem direção, pois a seta indica a direção de progressão do processo.
- `MessageFlow` é utilizado para mostrar o relacionamento entre dois elementos de pool ou entidades externas.
- `Association` é usado para conectar informações adicionais a um
   objeto de fluxo, como anotações, documentos, arquivos.

`SwimLanes` são utilizados para organizar e estruturar as atividades e responsabilidades dentro de um diagrama de processo. As categorias aqui são Pools e Lanes.

`Pool` representação de uma entidade externa, organização ou papel envolvido no processo de negócio;
Lane: Subset do elemento de Pool. Utilizado para detalhar melhor uma organização quanto aos usuários e processos.

Em Artifacts a categoria única é DataObject onde se destacam os subelemntos que são:

- `Data Object` representação de dados utilizados em um processo;
- `Data Store object` representa um lugar físico para armazenar esses dados, interagindo diretamente com o Armazenamento de dados do processo.
- `Data Input/Output Association` Representa a entrada ou saída de dados em um fluxo de processo.

A categoria `DarkPool` é utilizado quando se deseja dar foco em eventos, tarefas ou decisões específicas em um processo maior, escondendo outros detalhes dos objetos sob o dark pool.

Para `Annotations` o elemento única tem o nome Anotation que é um elemento adicional usado para fornecer mais informações sobre um processo ou objeto específico dentro de um diagrama de processos.

### Transições

Para representar transitionCondition como um objeto em JSON, você pode usar a notação de chave: valor para descrevê-lo dentro do objeto que representa uma determinada transição. Um exemplo seria:

```json
{
 "id": "transition_id_1",
 "name": "Transition 1",
 "from": "state_id_1",
 "to": "state_id_2",
 "transitionCondition": {
  "operator": "<=",
  "operand": "10"
 }
}
```

Se expression for uma representação para um atributo do próprio elemento, basta substituir o valor da propriedade value pelo nome do atributo. Por exemplo:

```json
{
 "id": "transition_id_1",
 "name": "Transition 1",
 "from": "state_id_1",
 "to": "state_id_2",
 "transitionCondition": {
  "operator": "<=",
  "operand": "10"
 },
 "expression": "this.status == 'approved'"
}
```

Este código é um exemplo de uma condição de transição que pode ser encontrada no campo "transitionCondition" do objeto descrito anteriormente. Essa condição contém as seguintes informações:

- `id`: o identificador único da condição de transição;
- `name`: o nome da condição de transição;
- `from`: o estado de origem da transição;
- `to`: o estado de destino da transição;
- `transitionCondition`: um objeto que contém informações sobre a condição de transição (neste caso, o operador `<=` e o operando `10`);
- `expression`: uma expressão que deve ser avaliada para determinar se a transição deve ser executada (neste caso, a expressão é `this.status == 'approved'`).

#### Exemplo de um processo de negócios com transições condicionais

Suponha que temos um processo de negócios para aprovação de pedidos de compra. O objeto de processo de negócios pode ser definido da seguinte forma:

```json
{
 "id": "pedido_compra_001",
 "name": "Pedido de Compra 001",
 "category": "FlowObjects/Activities",
 "type": "Activities",
 "subType": "Start",
 "incoming": [],
 "outgoing": [
  "transicao_001"
 ],
 "catching": true,
 "catchDefinition": {},
 "throwing": false,
 "throwDefinition": {},
 "catchs": [],
 "throws": [],
 "refs": [],
 "triggeredBy": "",
 "timestamp": {},
 "version": {},
 "state": "Created",
 "stateReason": "",
 "attempts": 0,
 "maxAttempts": 3,
 "retryDelay": 30,
 "transitionCondition": []
}
```

Neste exemplo, o objeto de processo de negócios representa o início do processo de aprovação de um pedido de compra. Ele tem um identificador único, um nome, uma categoria, um tipo, um subtipo e outras informações relevantes sobre o processo. Além disso, ele tem uma transição de saída chamada "transicao_001".

Agora, vamos adicionar um objeto de transição para esta transição:

```json
{
 "id": "transicao_001",
 "name": "Transição 001",
 "from": "Start",
 "to": "Aprovacao Gerente",
 "transitionCondition": {
  "operator": "==",
  "operand": "true"
 },
 "expression": "this.valor <= this.limite"
}
```

Neste exemplo, o objeto de transição representa a transição que leva o processo do estado "Start" para o estado "Aprovacao Gerente". Ele tem um identificador único, um nome, o estado de origem e destino, uma condição de transição (neste caso, o operador "==" e o operando "true") e uma expressão que deve ser avaliada para determinar se a transição deve ser executada.

Para usar esses objetos em um sistema de controle de processo de negócios, podemos adicionar lógica booleana, um índice de progresso, parâmetros críticos e interação com outras variáveis / atributos. Por exemplo, podemos criar uma variável booleana chamada "autorizado" que indica se o pedido de compra foi autorizado ou não. Podemos então definir a condição de transição da seguinte forma:

```json
{
 "operator": "==",
 "operand": "true",
 "expression": "this.autorizado == true"
}
```

Neste caso, a transição só será executada se a variável "autorizado" for verdadeira. Além disso, podemos adicionar um índice de progresso para monitorar o status geral do processo, parâmetros críticos para definir limites e metas e interação com outras variáveis / atributos para permitir que o processo seja adaptado com base em mudanças em seu ambiente ou entrada.

### Sub-processos

Sim, em BPMN, é possível utilizar um array para armazenar uma lista de referências para outros elementos. Na verdade, essa é uma abordagem comum para descrever a sequência de atividades em um processo.

Por exemplo, em um sub-processo, podemos usar um array para armazenar uma lista de referências para outras atividades, como tarefas (tasks) ou outros sub-processos. Cada referência pode ser uma string que identifica o ID do elemento correspondente.

Segue um exemplo de como poderíamos representar um sub-processo com um array de referências para atividades:

```json
{
 "id": "subprocess_1",
 "name": "Processo de vendas",
 "category": "FlowObjects/Subprocesses",
 "type": "Subprocesses",
 "subType": "",
 "incoming": [
  "flow_id_1"
 ],
 "outgoing": [
  "flow_id_2"
 ],
 "catching": false,
 "catchDefinition": {},
 "throwing": false,
 "throwDefinition": {},
 "catchs": [],
 "throws": [],
 "refs": [],
 "triggeredBy": "",
 "timestamp": {
  "created": "2022-03-14T12:00:00Z",
  "updated": "",
  "executed": "",
  "triggered": ""
 },
 "version": {
  "major": 1,
  "minor": 0,
  "patch": 0
 },
 "state": "Created",
 "stateReason": "",
 "attempts": 0,
 "maxAttempts": 3,
 "retryDelay": 30,
 "transitionCondition": [],
 "activities": [
  "task_1",
  "subprocess_2",
  "task_2"
 ]
}
```

Nesse exemplo, temos uma lista de elementos que representam as atividades do workflow. Cada elemento tem um ID e um nome, além das coleções incoming e outgoing, que representam as conexões do elemento com outros elementos do workflow.

O campo start representa o elemento de início do workflow, enquanto o campo end representa o elemento de fim do workflow. No exemplo acima, o elemento de início é o element2 e o elemento de fim é o element5.

Essa estrutura pode ser facilmente representada em uma GUI lowcode como o Node-RED, permitindo que os usuários criem e modifiquem workflows conectando elementos entre si por meio das coleções incoming e outgoing.

```json
{
  "elements": [
    {
      "id": "element1",
      "name": "Element 1",
      "incoming": ["element2"],
      "outgoing": ["element3", "element4"]
    },
    {
      "id": "element2",
      "name": "Element 2",
      "incoming": [],
      "outgoing": ["element1"]
    },
    {
      "id": "element3",
      "name": "Element 3",
      "incoming": ["element1"],
      "outgoing": ["element5"]
    },
    {
      "id": "element4",
      "name": "Element 4",
      "incoming": ["element1"],
      "outgoing": []
    },
    {
      "id": "element5",
      "name": "Element 5",
      "incoming": ["element3"],
      "outgoing": []
    }
  ],
  "start": ["element2"],
  "end": ["element5"]
}
```
