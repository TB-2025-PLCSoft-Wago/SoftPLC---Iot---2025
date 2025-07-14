import React, { FC } from 'react';
import {
    BaseEdge,
    EdgeLabelRenderer,
    type EdgeProps,
} from 'reactflow';

// Custom function to make a "step" edge
function getCustomStepPath({
                               sourceX,
                               sourceY,
                               targetX,
                               targetY,
                           }: {
    sourceX: number;
    sourceY: number;
    targetX: number;
    targetY: number;
}): string {
    const midX = (sourceX + targetX) / 2;
    return `M${sourceX},${sourceY} L${midX},${sourceY} L${midX},${targetY} L${targetX},${targetY}`;
}

const CustomEdgeStartEndDebug: FC<EdgeProps> = ({
                                                    id,
                                                    sourceX,
                                                    sourceY,
                                                    targetX,
                                                    targetY,
                                                    style = {},
                                                    data,
                                                }) => {
    const edgePath = getCustomStepPath({ sourceX, sourceY, targetX, targetY });

    return (
        <>
            <BaseEdge id={id} path={edgePath} style={style} />
            <EdgeLabelRenderer>
                <div
                    style={{
                        position: 'absolute',
                        background: 'rgba(255, 255, 255, 0.25)',
                        padding: '5px 10px',
                        color: '#000',
                        fontSize: 12,
                        fontWeight: 700,
                        transform: `translate(-100%, -80%) translate(${targetX}px,${targetY}px)`,
                        whiteSpace: 'nowrap',
                        pointerEvents: 'all',
                    }}
                    className="nodrag nopan"
                >
                    {data?.label ?? '???'}
                </div>
            </EdgeLabelRenderer>
        </>
    );
};

export default CustomEdgeStartEndDebug;
