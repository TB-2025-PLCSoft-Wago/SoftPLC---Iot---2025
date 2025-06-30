// ConnectionLine.tsx
import { getBezierPath, type ConnectionLineComponentProps } from 'reactflow';

export default function ConnectionLine({
                                           fromX,
                                           fromY,
                                           toX,
                                           toY,
                                           fromPosition,
                                           toPosition,
                                           connectionLineStyle,
                                       }: ConnectionLineComponentProps) {
    const [edgePath] = getBezierPath({
        sourceX: fromX,
        sourceY: fromY,
        sourcePosition: fromPosition,
        targetX: toX,
        targetY: toY,
        targetPosition: toPosition,
    });

    return (
        <path
            d={edgePath}
            style={connectionLineStyle}
            className="react-flow__connection-path"
        />
    );
}
