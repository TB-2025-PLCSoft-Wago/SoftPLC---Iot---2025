import { useEffect, useRef } from 'react';
import {Edge} from "reactflow";
import { setWebSocket } from './WebSocketInstanceEdgeClicked';
type Props = {
    setEdges: (edges: any[]) => void;
};

const WebSocketManager = ({ setEdges }: Props) => {
    const ws = useRef<WebSocket | null>(null);

    useEffect(() => {
        ws.current = new WebSocket('ws://localhost:8890/ws');
        console.log('🔌 WebSocket opened from WebSocketManager');
        setWebSocket(ws.current); //recup ws for tools
        ws.current.onmessage = (event) => {
            try {
                const data = JSON.parse(event.data);
                if (data.edges) {
                    console.log('🔄 Updating edges from WebSocket:', data.edges);

                    const newEdgesFromWS: Edge[] = data.edges.map((edge: any) => ({
                        ...edge,
                        label: edge.label ?? '???',
                        labelBgPadding: [8, 4],
                        labelBgBorderRadius: 4,
                        labelBgStyle: { fill: 'white', color: '#333', fillOpacity: 0.5 },
                        style: {
                            ...edge.style,
                            strokeWidth: 1
                        }
                    }));

                    setEdges((prevEdges: Edge[]): Edge[] => {
                        // Remplacer les anciens edges qui ont le même ID
                        const edgeMap = new Map<string, Edge>();

                        // Ajouter tous les anciens
                        for (const e of prevEdges) {
                            edgeMap.set(e.id, e);
                        }

                        // Remplacer ou ajouter ceux reçus du backend
                        for (const newEdge of newEdgesFromWS) {
                            edgeMap.set(newEdge.id, newEdge);
                        }

                        // Retourner la fusion
                        return Array.from(edgeMap.values());
                    });
                }
            } catch (err) {
                console.error('❌ Error parsing WebSocket message:', err);
            }
        };


        return () => {
            console.log('🔌 WebSocketManager closing connection');
            ws.current?.close();
        };
    }, [setEdges]);

    return null; // invisible
};

export default WebSocketManager;
