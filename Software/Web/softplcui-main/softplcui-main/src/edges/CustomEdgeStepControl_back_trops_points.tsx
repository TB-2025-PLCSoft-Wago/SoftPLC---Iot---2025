import React, { useCallback, useEffect, useState } from 'react';
import { BaseEdge, EdgeProps, useReactFlow } from 'reactflow';

type Point = { x: number; y: number };

function getZigzagPath(source: Point, target: Point, controls: Point[]): string {
    const allPoints = [source, ...controls, target];
    return allPoints.reduce(
        (acc, point, i) =>
            i === 0 ? `M${point.x},${point.y}` : `${acc} L${point.x},${point.y}`,
        ''
    );
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
    const { setEdges, project } = useReactFlow();
    const [dragIndex, setDragIndex] = useState<number | null>(null);

    const controls: Point[] = data?.controls ?? [
        { x: (sourceX + targetX) / 2, y: sourceY },
    ];

    const onMouseDown = (index: number) => (event: React.MouseEvent) => {
        event.stopPropagation();
        setDragIndex(index);
    };

    const onMouseMove = useCallback(
        (event: MouseEvent) => {
            if (dragIndex === null) return;

            const flowPoint = project({
                x: event.clientX,
                y: event.clientY,
            });

            setEdges((edges) =>
                edges.map((edge) => {
                    if (edge.id !== id) return edge;

                    const newControls = [...(edge.data?.controls ?? [])];
                    newControls[dragIndex] = flowPoint;

                    // 💡 Ajout d’un nouveau point si on déplace l’un des deux extrêmes
                    const isEnd = dragIndex === 0 || dragIndex === newControls.length - 1;
                    const distance = 30;

                    if (isEnd) {
                        const dir = dragIndex === 0 ? -1 : 1;
                        const newPoint = {
                            x: flowPoint.x + distance * dir,
                            y: flowPoint.y + distance * dir,
                        };
                        newControls.splice(dragIndex === 0 ? 0 : newControls.length, 0, newPoint);
                    }

                    return {
                        ...edge,
                        data: {
                            ...edge.data,
                            controls: newControls,
                        },
                    };
                })
            );
        },
        [dragIndex, id, setEdges, project]
    );

    const onMouseUp = () => {
        setDragIndex(null);
    };

    useEffect(() => {
        if (dragIndex !== null) {
            window.addEventListener('mousemove', onMouseMove);
            window.addEventListener('mouseup', onMouseUp);
        }
        return () => {
            window.removeEventListener('mousemove', onMouseMove);
            window.removeEventListener('mouseup', onMouseUp);
        };
    }, [dragIndex, onMouseMove]);

    const path = getZigzagPath(
        { x: sourceX, y: sourceY },
        { x: targetX, y: targetY },
        controls
    );

    // Calcul des milieux pour les cercles interactifs
    const segments = [
        { from: { x: sourceX, y: sourceY }, to: controls[0] },
        ...controls.map((pt, i) =>
            i < controls.length - 1
                ? { from: pt, to: controls[i + 1] }
                : { from: controls[i], to: { x: targetX, y: targetY } }
        ),
    ];

    return (
        <>
            <BaseEdge id={id} path={path} style={style} />
            {segments.map(({ from, to }, index) => {
                const midX = (from.x + to.x) / 2;
                const midY = (from.y + to.y) / 2;
                return (
                    <circle
                        key={index}
                        cx={midX}
                        cy={midY}
                        r={6}
                        fill="#fff"
                        stroke="#000"
                        strokeWidth={1.5}
                        style={{ cursor: 'grab', pointerEvents: 'all' }}
                        onMouseDown={onMouseDown(index)}
                    />
                );
            })}
        </>
    );
};

export default CustomEdgeStepControl;
