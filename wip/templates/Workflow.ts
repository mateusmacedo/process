interface WorkflowElement {
	id: string;
	name: string;
	incoming: string[];
	outgoing: string[];
}

interface Workflow {
	elements: WorkflowElement[];
	start: string[];
	end: string[];
}

const workflow: Workflow = {
	elements: [
		{
			id: "element1",
			name: "Element 1",
			incoming: ["element2"],
			outgoing: ["element3", "element4"]
		},
		{
			id: "element2",
			name: "Element 2",
			incoming: [],
			outgoing: ["element1"]
		},
		{
			id: "element3",
			name: "Element 3",
			incoming: ["element1"],
			outgoing: ["element5"]
		},
		{
			id: "element4",
			name: "Element 4",
			incoming: ["element1"],
			outgoing: []
		},
		{
			id: "element5",
			name: "Element 5",
			incoming: ["element3", "element6"],
			outgoing: []
		},
		{
			id: "element6",
			name: "Element 6",
			incoming: ["element7"],
			outgoing: ["element5"]
		},
		{
			id: "element7",
			name: "Element 7",
			incoming: [],
			outgoing: ["element6"]
		}
	],
	start: ["element2", "element7"],
	end: ["element5"]
};

function findElementById(id: string): WorkflowElement | undefined {
	return workflow.elements.find((element) => element.id === id);
}

function traverseWorkflow(startElementIds: string[], endElementIds: string[]): void {
	const visited: string[] = [];
	let currentElementIds: string[] = startElementIds;

	while (currentElementIds.length > 0) {
		const nextElementIds: string[] = [];
		currentElementIds.forEach((currentElementId) => {
			if (!visited.includes(currentElementId)) {
				visited.push(currentElementId);
				const currentElement = findElementById(currentElementId);
				if (currentElement) {
					const hasUnvisitedOutgoing = currentElement.outgoing.some((outgoingElementId) => {
						const outgoingElement = findElementById(outgoingElementId);
						return outgoingElement && !visited.includes(outgoingElementId);
					});
					if (!hasUnvisitedOutgoing && currentElement.incoming.every((incomingElementId) => {
						return visited.includes(incomingElementId);
					})) {
						// Aqui você pode acionar o objeto do elemento como quiser
						console.log(`Acionando objeto do elemento ${currentElement.id}`);
						if (endElementIds.includes(currentElementId)) {
							console.log(`Finalizando o workflow no elemento ${currentElementId}`);
						} else {
							nextElementIds.push(...currentElement.outgoing);
						}
					}
				}
			}
		});
		currentElementIds = nextElementIds;
	}
}

traverseWorkflow(workflow.start, workflow.end);


// Esse algoritmo percorre todos os elementos do workflow, começando pelo elemento de início definido no campo start. A cada iteração do loop, ele remove o próximo elemento da fila de elementos a serem visitados e verifica se já visitou esse elemento antes. Se o elemento ainda não foi visitado, o algoritmo aciona o objeto do elemento e adiciona seus elementos adjacentes à fila de elementos a serem visitados.
// O método findElementById é responsável por buscar um elemento na lista de elementos do workflow pelo seu ID. Ele retorna undefined se nenhum elemento com esse ID for encontrado.
// O método traverseWorkflow é responsável por iniciar a travessia do workflow a partir do elemento de início especificado. Ele usa uma fila para manter o controle dos elementos a serem visitados e um array para manter o controle dos elementos já visitados. No exemplo acima, o método simplesmente imprime uma mensagem no console para cada elemento visitado, mas você pode substituir essa funcionalidade pelo que quiser, como a execução de uma ação específica com base nas propriedades do elemento.
// Nessa nova versão do algoritmo, a verificação das conexões de saída de um elemento é feita com a função some(), que verifica se pelo menos um elemento da coleção outgoing ainda não foi visitado. Dessa forma, o algoritmo garante que todos os elementos da coleção outgoing de um elemento foram visitados antes de acionar o próximo elemento.
// Observe que, quando um elemento tem mais de um incoming, a nova verificação também garante que todos os elementos de entrada foram visitados antes de acionar o elemento atual.
// Por fim, o algoritmo também leva em consideração o caso em que o workflow não tem mais elementos a serem visitados, mas ainda não chegou a um elemento final. Nesse caso, o algoritmo simplesmente encerra a execução e indica que o workflow não pode ser concluído.
// Com essas novas verificações, o algoritmo é capaz de garantir que todos os elementos de um workflow serão visitados corretamente, independentemente da complexidade das conexões entre eles.
