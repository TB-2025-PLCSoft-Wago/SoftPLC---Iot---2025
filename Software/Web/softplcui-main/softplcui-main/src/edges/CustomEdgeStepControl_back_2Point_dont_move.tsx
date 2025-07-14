import React, { useCallback, useEffect, useState } from 'react';
import { BaseEdge, EdgeProps, useReactFlow } from 'reactflow';

type Point = { x: number; y: number };

function buildStepPath(points: Point[]): string {
    return points.reduce((path, point, i) => {
        const prefix = i === 0 ? 'M' : 'L';
        return `${path} ${prefix}${point.x},${point.y}`;
    }, '');
}

function getSegmentsFromPoints(points: Point[]) {
    const segments: { from: Point; to: Point }[] = [];
    for (let i = 0; i < points.length - 1; i++) {
        segments.push({ from: points[i], to: points[i + 1] });
    }
    return segments;
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
        { x: (sourceX + targetX) / 2, y: targetY },
    ];

    const fullPoints: Point[] = [{ x: sourceX, y: sourceY }, ...controls, { x: targetX, y: targetY }];
    const segments = getSegmentsFromPoints(fullPoints);

    const onMouseDown = (index: number) => (event: React.MouseEvent) => {
        console.log("onMouseDown")
        event.stopPropagation();
        setDragIndex(index);
    };

    const onMouseMove = useCallback(
        (event: MouseEvent) => {
            if (dragIndex === null) return;

            const flowPoint = project({ x: event.clientX, y: event.clientY });

            setEdges((edges) =>
                edges.map((edge) => {
                    if (edge.id !== id) return edge;

                    const currentControls = [...(edge.data?.controls ?? [])];

                    const isFirst = dragIndex === 0;
                    const isLast = dragIndex === currentControls.length - 1;
                    const control = currentControls[dragIndex];

                    const prevPoint = isFirst
                        ? { x: sourceX, y: sourceY }
                        : currentControls[dragIndex - 1];

                    const nextPoint = isLast
                        ? { x: targetX, y: targetY }
                        : currentControls[dragIndex + 1];

                    const isHorizontal = control.y === prevPoint.y && control.y === nextPoint.y;
                    const isVertical = control.x === prevPoint.x && control.x === nextPoint.x;

                    const newPoint: Point = { ...control };

                    if (isHorizontal) {
                        // Drag vertical seulement
                        newPoint.y = flowPoint.y;

                        if (isLast) {
                            // On ajoute un virage vertical + horizontal
                            const preTarget = { x: control.x, y: flowPoint.y };
                            const toTarget = { x: targetX, y: flowPoint.y };
                            const newControls = [...currentControls];
                            newControls[dragIndex] = preTarget;
                            newControls.splice(dragIndex + 1, 0, toTarget);

                            return {
                                ...edge,
                                data: {
                                    ...edge.data,
                                    controls: newControls,
                                },
                            };
                        }
                    } else if (isVertical) {
                        // Drag horizontal seulement
                        newPoint.x = flowPoint.x;
                    }

                    currentControls[dragIndex] = newPoint;

                    return {
                        ...edge,
                        data: {
                            ...edge.data,
                            controls: currentControls,
                        },
                    };
                })
            );
        },
        [dragIndex, id, setEdges, project, sourceX, sourceY, targetX, targetY]
    );

    const onMouseUp = () => {
        console.log("custom Edge Click onMouseUp")
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

    const path = buildStepPath(fullPoints);

    // On affiche des cercles uniquement sur les segments horizontaux sauf le dernier vers le node
    const visibleControlPoints = segments
        .map(({ from, to }, index) => {
            const isLast = index === segments.length - 2;
            const isHorizontal = from.y === to.y;
            //if (!isHorizontal) return null;
            //if (isLast) return null;

            const midX = (from.x + to.x) / 2;
            const midY = (from.y + to.y)/2;
            return { x: midX, y: midY, index };
        })
        .filter(Boolean) as { x: number; y: number; index: number }[];

    return (
        <>
            <BaseEdge id={id} path={path} style={style} />
            {visibleControlPoints.map(({ x, y, index }) => (
                <circle
                    key={index}
                    cx={x}
                    cy={y}
                    r={6}
                    fill="#fff"
                    stroke="#000"
                    strokeWidth={1.5}
                    style={{ cursor: 'ns-resize', pointerEvents: 'all' }}
                    onMouseDown={onMouseDown(index)}
                />
            ))}
        </>
    );
};

export default CustomEdgeStepControl;
