// hooks/useKeyboardShortcuts.ts
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
    const pushToUndoStack = () => {
        undoStack.current.push({
            nodes: JSON.parse(JSON.stringify(nodes)),
            edges: JSON.parse(JSON.stringify(edges)),
        });
        redoStack.current = [];
    };


    useEffect(() => {
        if(!isDragging){
            //console.log("UseEffectKeyboard")
        }else{
            //console.log("UseEffectKeyboard error")
        }
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

                    // Supprimer onChange du clone
                    const { onChange, ...cleanData } = node.data || {};

                    return {
                        ...node,
                        id: newId,
                        position: {
                            x: node.position.x + 40,
                            y: node.position.y + 40,
                        },
                        data: cleanData,
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

                    const newNodesWithHandlers = newNodes.map((node) => {
                        if (node.type === "commentNode") {
                            return {
                                ...node,
                                data: {
                                    ...node.data,
                                    onChange: (newText: string) => {
                                        setNodes((nds) =>
                                            nds.map((n) =>
                                                n.id === node.id
                                                    ? {
                                                        ...n,
                                                        data: {
                                                            ...n.data,
                                                            text: newText,
                                                        },
                                                    }
                                                    : n
                                            )
                                        );
                                    },
                                },
                            };
                        }
                        return node;
                    });

                    return [...deselected, ...newNodesWithHandlers];
                });

                setEdges((prev) => [...prev, ...newEdges]);

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

    useDebouncedUndo(nodes, edges, pushToUndoStack, undoStack);



}



