import { useMemo } from 'react';
import { getConnectedEdges, Handle, useNodeId, useStore } from 'reactflow';

type Node = {
    id: string;
    type: string;
};
type Edge = {
    source: string;
    target: string;
    sourceHandle: string;
    targetHandle: string;
};

const selector = (s: { nodeInternals: Map<string, Node>; edges: Edge[]; }) => ({
    nodeInternals: s.nodeInternals,
    edges: s.edges,
});

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-expect-error
const CustomHandle = (props) => {
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-expect-error
    const { nodeInternals, edges } = useStore(selector);
    const nodeId = useNodeId();

    const isHandleConnectable = useMemo(() => {
        if (nodeId !== null) {
        if (typeof props.isConnectable === 'function') {
            const node = nodeInternals.get(nodeId);
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-expect-error
            const connectedEdges = getConnectedEdges([node], edges);

            // Filter the connected edges for this specific handle
            const connectedEdgesForThisHandle = connectedEdges.filter(edge =>
                (edge.source === nodeId && edge.sourceHandle === props.id) ||
                (edge.target === nodeId && edge.targetHandle === props.id)
            );

            return props.isConnectable({node, connectedEdges: connectedEdgesForThisHandle});
        }

        if (typeof props.isConnectable === 'number') {
            const node = nodeInternals.get(nodeId);
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-expect-error
            const connectedEdges = getConnectedEdges([node], edges);

            // Filter the connected edges for this specific handle
            const connectedEdgesForThisHandle = connectedEdges.filter(edge =>
                (edge.source === nodeId && edge.sourceHandle === props.id) ||
                (edge.target === nodeId && edge.targetHandle === props.id)
            );

            return connectedEdgesForThisHandle.length < props.isConnectable;
        }

        return props.isConnectable;
    } else {
        return props.isConnectable(false);
        }
    }, [nodeInternals, edges, nodeId, props.isConnectable, props.id]);

    return (
        <Handle {...props} isConnectable={isHandleConnectable}></Handle>
    );
};

export default CustomHandle;