import { useEffect, useRef } from "react";
import { Node, Edge } from "reactflow";
import useDebouncedUndo from "./useDebouncedUndo";

type UseKeyboardShortcutsProps = {
    nodes: Node[];
    edges: Edge[];
    setNodes: React.Dispatch<React.SetStateAction<Node[]>>;
    setEdges: React.Dispatch<React.SetStateAction<Edge[]>>;
    getId: () => string;
    isDragging: boolean;
};

export default function useKeyboardShortcuts({
                                                 nodes,
                                                 edges,
                                                 setNodes,
                                                 setEdges,
                                                 getId,
                                                 isDragging,
                                             }: UseKeyboardShortcutsProps) {
    const copiedDataRef = useRef<{ nodes: Node[]; edges: Edge[] }>({ nodes: [], edges: [] });
    const undoStack = useRef<{ nodes: Node[]; edges: Edge[] }[]>([]);
    const redoStack = useRef<{ nodes: Node[]; edges: Edge[] }[]>([]);
    const manualPushRef = useRef(false);

    const nodesRef = useRef<Node[]>(nodes);
    const edgesRef = useRef<Edge[]>(edges);

    useEffect(() => {
        nodesRef.current = nodes;
        edgesRef.current = edges;
    }, [nodes, edges]);

    const pushToUndoStack = () => {
        const currentNodes = nodesRef.current;
        const currentEdges = edgesRef.current;
        undoStack.current.push({
            nodes: structuredClone(currentNodes),
            edges: structuredClone(currentEdges),
        });
        //redoStack.current = [];
    };

    useEffect(() => {
        if (!isDragging) {
            console.log("UseEffectKeyboard");
        }

        const handleKeyDown = (event: KeyboardEvent) => {
            const activeElement = document.activeElement;
            const isInputFocused =
                activeElement instanceof HTMLInputElement ||
                activeElement instanceof HTMLTextAreaElement ||
                (activeElement && (activeElement as HTMLElement).isContentEditable);
            if (isInputFocused) return;

            // Copy / Cut
            if ((event.ctrlKey || event.metaKey) && (event.key === "c" || event.key === "x")) {
                const selectedNodes = nodesRef.current.filter((n) => n.selected);
                const selectedNodeIds = new Set(selectedNodes.map((n) => n.id));

                const selectedEdges = edgesRef.current.filter(
                    (e) => selectedNodeIds.has(e.source) && selectedNodeIds.has(e.target)
                );

                copiedDataRef.current = {
                    nodes: selectedNodes,
                    edges: selectedEdges,
                };

                console.log("copy :", copiedDataRef.current);

                if (event.key === "x") {
                    manualPushRef.current = true;
                    pushToUndoStack();
                    setNodes((prev) => prev.filter((n) => !selectedNodeIds.has(n.id)));
                    setEdges((prev) =>
                        prev.filter(
                            (e) =>
                                !selectedNodeIds.has(e.source) &&
                                !selectedNodeIds.has(e.target)
                        )
                    );
                }
            }

            // Paste
            if ((event.ctrlKey || event.metaKey) && event.key === "v") {
                manualPushRef.current = true;
                pushToUndoStack();

                const { nodes: copiedNodes, edges: copiedEdges } = copiedDataRef.current;
                const idMap = new Map<string, string>();

                const newNodes = copiedNodes.map((node) => {
                    const newId = getId();
                    idMap.set(node.id, newId);
                    return {
                        ...structuredClone(node),
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
                    ...structuredClone(edge),
                    id: getId(),
                    source: idMap.get(edge.source) || edge.source,
                    target: idMap.get(edge.target) || edge.target,
                }));

                setNodes((prev) => {
                    const deselected = prev.map((n) => ({ ...n, selected: false }));
                    return [...deselected, ...newNodes];
                });
                setEdges((prev) => [...prev, ...newEdges]);

                console.log("paste :", { nodes: newNodes, edges: newEdges });
            }

            // Undo (Ctrl + Z)
            if ((event.ctrlKey || event.metaKey) && event.key === "z") {
                event.preventDefault();
                manualPushRef.current = true;
                const lastState = undoStack.current.pop();
                if (lastState) {
                    redoStack.current.push({
                        nodes: structuredClone(nodesRef.current),
                        edges: structuredClone(edgesRef.current),
                    });
                    setNodes(lastState.nodes);
                    setEdges(lastState.edges);
                }
                return;
            }

            // Redo (Ctrl + Y)
            if ((event.ctrlKey || event.metaKey) && event.key === "y") {
                event.preventDefault();
                manualPushRef.current = true;
                const nextState = redoStack.current.pop();
                if (nextState) {
                    undoStack.current.push({
                        nodes: structuredClone(nodesRef.current),
                        edges: structuredClone(edgesRef.current),
                    });
                    setNodes(nextState.nodes);
                    setEdges(nextState.edges);
                }
                return;
            }
        };

        window.addEventListener("keydown", handleKeyDown);
        return () => window.removeEventListener("keydown", handleKeyDown);
    }, [setNodes, setEdges, getId]);

    useDebouncedUndo(nodes, edges, pushToUndoStack, undoStack, manualPushRef);
}
