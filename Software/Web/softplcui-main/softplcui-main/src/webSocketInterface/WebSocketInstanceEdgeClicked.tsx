let socket: WebSocket | null = null;

export function setWebSocket(ws: WebSocket) {
    socket = ws;
}

export function sendEdgeClicked(data: {
    source: string;
    sourceHandle?: string | null;
    tool: string;
}) {
    if (socket && socket.readyState === WebSocket.OPEN) {
        socket.send(JSON.stringify({ type: 'edge_clicked', ...data }));
    } else {
        console.warn('WebSocket not connected');
    }
}
