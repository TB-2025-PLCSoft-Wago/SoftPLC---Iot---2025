import React, { useCallback, useEffect, useRef, useState } from 'react';
import { BaseEdge, EdgeLabelRenderer, type EdgeProps, useReactFlow } from 'reactflow';

type ControlData = {
    x: number;
    y: number;
};

function getStepPathWithControl(
    sourceX: number,
    sourceY: number,
    targetX: number,
    targetY: number,
    control: ControlData
): string {
    return `M${sourceX},${sourceY} L${control.x},${sourceY} 
            L${control.x},${sourceY+50} L${control.x+50},${sourceY+50} 
            L${control.x+50},${targetY} L${targetX},${targetY}`;
}

const CustomEdgeStepControl: React.FC<EdgeProps> = ({
                                                        id,
                                                        sourceX,
                                                        sourceY,
                                                        targetX,
                                                        targetY,
                                                        style,
                                                        data,
                                                    }) => {
    const { setEdges } = useReactFlow();

    const [dragging, setDragging] = useState(false);
    const controlRef = useRef<SVGCircleElement | null>(null);

    const controlX = data?.control?.x ?? (sourceX + targetX) / 2;
    const controlY = sourceY; // For vertical step edges, Y stays aligned with source

    const onMouseDown = (event: React.MouseEvent) => {
        console.log("custom Edge Click down")
        event.stopPropagation();
        setDragging(true);
    };

    const { project } = useReactFlow();

    const onMouseMove = useCallback(
        (event: MouseEvent) => {
            if (!dragging) return;

            const x = event.clientX;
            const y = event.clientY;

            const flowPoint = project({ x, y });

            setEdges((eds) =>
                eds.map((edge) =>
                    edge.id === id
                        ? {
                            ...edge,
                            data: {
                                ...edge.data,
                                control: {
                                    x: flowPoint.x,
                                    y: flowPoint.y,
                                },
                            },
                        }
                        : edge
                )
            );
        },
        [dragging, id, setEdges, project]
    );


    const onMouseUp = () => {
        console.log("custom Edge Click onMouseUp")
        setDragging(false);
    };

    useEffect(() => {
        if (dragging) {
            window.addEventListener('mousemove', onMouseMove);
            window.addEventListener('mouseup', onMouseUp);
        } else {
            window.removeEventListener('mousemove', onMouseMove);
            window.removeEventListener('mouseup', onMouseUp);
        }

        return () => {
            window.removeEventListener('mousemove', onMouseMove);
            window.removeEventListener('mouseup', onMouseUp);
        };
    }, [dragging, onMouseMove]);

    const edgePath = getStepPathWithControl(sourceX, sourceY, targetX, targetY, {
        x: controlX,
        y: controlY,
    });

    return (
        <>
            <BaseEdge id={id} path={edgePath} style={style} />
            <circle
                ref={controlRef}
                cx={controlX}
                cy={controlY}
                r={6}
                fill="#fff"
                stroke="#000"
                strokeWidth={1.5}
                onMouseDown={onMouseDown}
                style={{ cursor: 'grab',pointerEvents: 'all', }}
                className=""
            />
            {/*
            <EdgeLabelRenderer>
                <div
                    style={{
                        position: 'absolute',
                        transform: `translate(-80%, -50%) translate(${targetX}px,${targetY}px)`,
                        background: 'white',
                        padding: '4px 8px',
                        fontSize: 12,
                        pointerEvents: 'all',
                        border: '1px solid #aaa',
                        borderRadius: 4,
                    }}
                    className="nodrag nopan"
                >
                    {data?.label ?? ''}
                </div>
            </EdgeLabelRenderer>
            */}
        </>
    );
};

export default CustomEdgeStepControl;
