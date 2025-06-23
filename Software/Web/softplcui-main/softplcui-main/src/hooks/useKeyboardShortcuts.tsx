// hooks/useKeyboardShortcuts.ts
import { useEffect, useRef } from "react";
import { Node, Edge } from "reactflow";

type UseKeyboardShortcutsProps = {
    nodes: Node[];
    edges: Edge[];
    setNodes: React.Dispatch<React.SetStateAction<Node[]>>;
    setEdges: React.Dispatch<React.SetStateAction<Edge[]>>;
    getId: () => string;
};



export default function useKeyboardShortcuts({
                                                 nodes,
                                                 edges,
                                                 setNodes,
                                                 setEdges,
                                                 getId,
                                             }: UseKeyboardShortcutsProps) {
    const copiedDataRef = useRef<{ nodes: Node[]; edges: Edge[] }>({ nodes: [], edges: [] });
    const undoStack = useRef<{ nodes: Node[]; edges: Edge[] }[]>([]);
    const redoStack = useRef<{ nodes: Node[]; edges: Edge[] }[]>([]);
    const pushToUndoStack = () => {
        undoStack.current.push({
            nodes: structuredClone(nodes),
            edges: structuredClone(edges),
        });
        // On vide la redo stack à chaque nouvelle action
        redoStack.current = [];
    };

    useEffect(() => {
        const handleKeyDown = (event: KeyboardEvent) => {
            const activeElement = document.activeElement;
            const isInputFocused =
                activeElement instanceof HTMLInputElement ||
                activeElement instanceof HTMLTextAreaElement ||
                (activeElement && (activeElement as HTMLElement).isContentEditable);
            if (isInputFocused) return;
            // Copy ctrl + c
            if ((event.ctrlKey || event.metaKey) && event.key === "c" || event.key === "x") {
                const selectedNodes = nodes.filter((n) => n.selected);
                const selectedNodeIds = new Set(selectedNodes.map((n) => n.id));

                const selectedEdges = edges.filter(
                    (e) => selectedNodeIds.has(e.source) && selectedNodeIds.has(e.target)
                );

                copiedDataRef.current = {
                    nodes: selectedNodes,
                    edges: selectedEdges,
                };

                //print in the developer console Ctrl + Maj + J
                console.log("copy :", {
                    nodes: selectedNodes,
                    edges: selectedEdges,
                });

                //delete node after copy when x
                if (event.key === "x"){
                    pushToUndoStack();
                    setNodes((prev) => prev.filter((n) => !selectedNodeIds.has(n.id)));
                    setEdges((prev) =>
                        prev.filter(
                            (e) => !selectedNodeIds.has(e.source) && !selectedNodeIds.has(e.target) //keep if is not selecting
                        )
                    );
                }
            }

            // Paste ctrl + v
            if ((event.ctrlKey || event.metaKey) && event.key === "v") {
                pushToUndoStack();
                const { nodes: copiedNodes, edges: copiedEdges } = copiedDataRef.current;

                const idMap = new Map<string, string>();

                const newNodes = copiedNodes.map((node) => {
                    const newId = getId();
                    idMap.set(node.id, newId);

                    return {
                        ...structuredClone(node), // Clone profond
                        id: newId,
                        position: {
                            x: node.position.x + 40,
                            y: node.position.y + 40,
                        },
                        selected: true,
                        draggable: true,
                    };
                });


                const newEdges = copiedEdges.map((edge) => ({
                    ...edge,
                    id: getId(),
                    source: idMap.get(edge.source) || edge.source,
                    target: idMap.get(edge.target) || edge.target,
                }));

                setNodes((prev) => {
                    const deselected = prev.map((n) => ({ ...n, selected: false }));
                    return [...deselected, ...newNodes]; // Deselect old, add new
                });
                setEdges((prev) => [...prev, ...newEdges]);

                //print in the developer console  Ctrl + Maj + J
                console.log("paste :", {
                    nodes: newNodes,
                    edges: newEdges,
                });
            }

            // Undo (Ctrl + Z)
            if ((event.ctrlKey || event.metaKey) && event.key === "z") {
                const lastState = undoStack.current.pop();
                if (lastState) {
                    redoStack.current.push({
                        nodes: structuredClone(nodes),
                        edges: structuredClone(edges),
                    });
                    setNodes(lastState.nodes);
                    setEdges(lastState.edges);
                }
                return;
            }

            // Redo (Ctrl + Y)
            if ((event.ctrlKey || event.metaKey) && event.key === "y") {
                const nextState = redoStack.current.pop();
                if (nextState) {
                    undoStack.current.push({
                        nodes: structuredClone(nodes),
                        edges: structuredClone(edges),
                    });
                    setNodes(nextState.nodes);
                    setEdges(nextState.edges);
                }
                return;
            }

        };

        window.addEventListener("keydown", handleKeyDown);
        return () => window.removeEventListener("keydown", handleKeyDown);
    }, [nodes, edges, setNodes, setEdges, getId]);
}
