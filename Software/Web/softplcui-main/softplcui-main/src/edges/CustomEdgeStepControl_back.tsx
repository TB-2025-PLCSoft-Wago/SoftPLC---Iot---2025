import React, {useCallback, useEffect, useMemo, useRef, useState} from 'react';
import { BaseEdge, EdgeProps, useReactFlow } from 'reactflow';

type Point = { x: number; y: number };

//most important
function getZigzagPath(source: Point, target: Point, controls: Point[]): string {
    const allPoints = [source, ...controls, target];
    //console.log("allPoints :", allPoints)
    return allPoints.reduce(
        (acc, point, i) =>
            i === 0 ? `M${point.x},${point.y}` : `${acc} L${point.x},${point.y}`,
        ''
    );
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
    const getDefaultControls = (sourceX: number, sourceY: number, targetX: number, targetY: number): Point[] => [
        { x: (sourceX + targetX) / 2, y: sourceY },
        { x: (sourceX + targetX) / 2, y: targetY },
    ];


    // const controls: Point[] = data?.controls ?? [
    //     { x: (sourceX + targetX) / 2, y: sourceY },
    // ];
    const controls: Point[] = data?.controls ?? getDefaultControls(sourceX, sourceY, targetX, targetY);


    const fullPoints: Point[] = [{ x: sourceX, y: sourceY }, ...controls, { x: targetX, y: targetY }];
    const segments = getSegmentsFromPoints(fullPoints);
    const [isNotMouseUp, setIsNotMouseUp] = useState(false);
    const isNotMouseUpRef = useRef(false); //référence synchronisée

    useEffect(() => {
        isNotMouseUpRef.current = isNotMouseUp;
    }, [isNotMouseUp]);

    const onMouseDown = (index: number) => (event: React.MouseEvent) => {
        event.stopPropagation();
        console.log("setDragIndex , index : ", index)
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

                    const  newControls = [...(edge.data?.controls ?? getDefaultControls(sourceX, sourceY, targetX, targetY))];
                    let adjustedDragIndex = dragIndex;

                    newControls[adjustedDragIndex] = flowPoint;
                    const prevPoint = newControls[adjustedDragIndex];
                    const dx = flowPoint.x - prevPoint.x;
                    const dy = flowPoint.y - prevPoint.y;

                    // Recrée les points complets avec source, controls, target
                    const controls = [...(edge.data?.controls ?? getDefaultControls(sourceX, sourceY, targetX, targetY))];
                    const fullPoints = [{ x: sourceX, y: sourceY }, ...controls, { x: targetX, y: targetY }];

                    // Recalcule les segments
                    const segments = getSegmentsFromPoints(fullPoints);
                    let isHorizontal;
                    // Vérifie si le dragIndex correspond à un segment
                    if (adjustedDragIndex < segments.length) {
                        const segment = segments[adjustedDragIndex];
                        isHorizontal = segment.from.y === segment.to.y;

                        console.log("isHorizontal :", isHorizontal);
                    }


                    // 💡 Ajout d’un nouveau point si on déplace l’un des deux extrêmes
                    const isEnd = dragIndex === 0 || adjustedDragIndex === controls.length;
                    console.log("controls.length :",controls.length)
                    console.log("newControls.length :",newControls.length)
                    console.log("dragIndex :",dragIndex)
                    console.log("dragIndex :",adjustedDragIndex)
                    console.log("condition isEnd :",isEnd)
                    console.log("newControls : ", newControls);
                    const distance = 10;

                    if (isEnd && !isNotMouseUpRef.current) {
                        setIsNotMouseUp(true);
                        const dir = adjustedDragIndex === 0 ? -1 : 1;



                        console.log("custom edge - dir :", dir);
                        if (dir === -1) {
                            // Ajout en début
                            const addedPoints = [
                                { x: sourceX - 10 * dir, y: sourceY },
                                { x: sourceX - 10 * dir, y: flowPoint.y },
                            ];
                            newControls.unshift(...addedPoints);

                            // Met à jour le dragIndex car on a ajouté au début
                            adjustedDragIndex += addedPoints.length;
                        } else {
                            const addedPoints = [
                                { x: targetX - 10 * dir, y: flowPoint.y },
                                { x: targetX - 10 * dir, y: targetY },
                            ];
                            newControls.push(...addedPoints);

                            // Pas besoin de modifier adjustedDragIndex si on ajoute en fin
                        }
                    }/*else{
                        // déplacer les voisins (si ils existent)
                        if (!isHorizontal){
                            if (adjustedDragIndex > 0) {
                                console.log("custom edge > 0 source side",dy)
                                newControls[adjustedDragIndex - 1] = {
                                    x: flowPoint.x,
                                    y: newControls[adjustedDragIndex - 1].y,
                                };
                            }
                            if (adjustedDragIndex < newControls.length - 1) {
                                console.log("custom edge target side :", dy)
                                newControls[adjustedDragIndex + 1] = {
                                    x: flowPoint.x,
                                    y: newControls[adjustedDragIndex + 1].y,
                                };
                            }
                        }else{
                            if (adjustedDragIndex > 0) {
                                console.log("custom edge > 0 source side",dy)
                                newControls[adjustedDragIndex - 1] = {
                                    x: newControls[adjustedDragIndex - 1].x,
                                    y: flowPoint.y,
                                };
                            }
                            if (adjustedDragIndex < newControls.length - 1) {
                                console.log("custom edge target side :", dy)
                                newControls[adjustedDragIndex + 1] = {
                                    x:  newControls[adjustedDragIndex + 1].x,
                                    y: flowPoint.y,
                                };
                            }

                        }

                    }
*/





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
        setIsNotMouseUp(false);
        console.log("onMouseUp")

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
    {/*// Calcul des milieux pour les cercles interactifs
    const segments = [
        { from: { x: sourceX, y: sourceY }, to: controls[0] },
        ...controls.map((pt, i) =>
            i < controls.length - 1
                ? { from: pt, to: controls[i + 1] }
                : { from: controls[i], to: { x: targetX, y: targetY } }
        ),
    ];*/}

    const visibleControlPoints =
        segments.map(({ from, to }, index) => {
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
