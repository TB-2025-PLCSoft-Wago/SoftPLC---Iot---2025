import { useEffect, useRef } from "react";
import { Node, Edge } from "reactflow";

type UndoStack = React.MutableRefObject<{ nodes: Node[]; edges: Edge[] }[]>;

export default function useDebouncedUndo(
    nodes: Node[],
    edges: Edge[],
    pushToUndoStack: () => void,
    undoStack: UndoStack,
    manualPushRef: React.MutableRefObject<boolean>,
    delay: number = 500
) {
    const debounceTimeout = useRef<number | null>(null);

    useEffect(() => {
        if (manualPushRef.current) {
            manualPushRef.current = false; // Ignore this one, it's manual
            return;
        }

        if (debounceTimeout.current !== null) {
            clearTimeout(debounceTimeout.current);
        }

        debounceTimeout.current = window.setTimeout(() => {
            const lastState = undoStack.current[undoStack.current.length - 1];

            const stringify = (obj: Node[] | Edge[]) => JSON.stringify(obj);

            const isDifferent =
                !lastState ||
                stringify(lastState.nodes) !== stringify(nodes) ||
                stringify(lastState.edges) !== stringify(edges);

            if (isDifferent) {
                pushToUndoStack();
                console.log("pushToUndoStack");
            }
        }, delay);

        return () => {
            if (debounceTimeout.current !== null) {
                clearTimeout(debounceTimeout.current);
            }
        };
    }, [nodes, edges, pushToUndoStack, undoStack, delay, manualPushRef]);
}
