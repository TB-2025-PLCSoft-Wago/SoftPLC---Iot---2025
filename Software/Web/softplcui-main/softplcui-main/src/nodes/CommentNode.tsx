// nodes/CommentNode.tsx
import React from 'react';
import { Handle, Position, NodeProps } from 'reactflow';
import './CommentNode.css'; // Style optionnel

const CommentNode = ({ data }: NodeProps) => {
    return (
        <div className="comment-node">
            <textarea
                className="comment-textarea"
                value={data.text}
                onChange={(e) => data.onChange(e.target.value)}
                placeholder="Add a comment ..."
            />
        </div>
    );
};

export default CommentNode;
