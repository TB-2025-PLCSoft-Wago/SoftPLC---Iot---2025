// CommentNode.tsx
import React from 'react';
import { NodeProps, useReactFlow } from 'reactflow';

export default function CommentNode({ id, data }: NodeProps) {
    const { setNodes } = useReactFlow();

    const handleChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        const newText = event.target.value;
        setNodes((nds) =>
            nds.map((node) =>
                node.id === id
                    ? {
                        ...node,
                        data: {
                            ...node.data,
                            text: newText,
                        },
                    }
                    : node
            )
        );
    };

    return (
        <div style={{ padding: 5, background: '#CFC9CE', borderRadius: 8 }}>
            <textarea
                value={data.text || ''}
                onChange={handleChange}
                style={{ width: 200, height: 30, background: '#FFFBDB'}}
                placeholder="Add a comment ..."
            />
        </div>
    );
}
