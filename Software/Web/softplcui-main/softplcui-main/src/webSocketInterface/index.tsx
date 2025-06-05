import React, { useEffect, useRef, useState } from 'react';

type Appliance = {
    name: string;
    buttons: { text: string; irCode: number }[];
};
type Output = {
    id: number;
    name: string;
    applianceName: string;
    type: string; // "bool", "string", "float"
    value: string;
};


const WebSocketView = () => {
    const [messages, setMessages] = useState<string[]>([]);
    const [input, setInput] = useState('');
    const [appliances, setAppliances] = useState<Appliance[]>([]);
    const ws = useRef<WebSocket | null>(null);
    const [outputs, setOutputs] = useState<Record<string, Output[]>>({});

    useEffect(() => {
        ws.current = new WebSocket('ws://localhost:8890/ws');

        ws.current.onopen = () => {
            console.log('✅ WebSocket connected');
        };

        ws.current.onmessage = (event) => {
            console.log('📨 Message received:', event.data);
            try {
                const data = JSON.parse(event.data);
                if (data.type === 'appliances') {
                    setAppliances(data.appliances);
                } else if (data.type === 'update') {
                    setOutputs((prev) => ({
                        ...prev,
                        [data.appliance]: data.outputs
                    }));
                } else {
                    setMessages((prev) => [...prev, event.data]);
                }
            } catch {
                setMessages((prev) => [...prev, event.data]);
            }
        };


        ws.current.onerror = (err) => {
            console.error('❌ WebSocket error:', err);
        };

        ws.current.onclose = () => {
            console.warn('🔌 WebSocket closed');
        };

        return () => {
            ws.current?.close();
        };
    }, []);

    const sendMessage = () => {
        if (ws.current?.readyState === WebSocket.OPEN) {
            ws.current.send(input);
            setInput('');
        } else {
            alert('WebSocket not connected.');
        }
    };

    const sendIRCode = (irCode: number) => {
        if (ws.current?.readyState === WebSocket.OPEN) {
            ws.current.send(JSON.stringify({ type: 'irCommand', irCode }));
        }
    };

    return (
        <div style={{ padding: '1rem' }}>
            <h2>🧠 WebSocket Interface</h2>

            <div style={{ marginBottom: '1rem' }}>
                <input
                    type="text"
                    placeholder="Enter message"
                    value={input}
                    onChange={(e) => setInput(e.target.value)}
                    onKeyDown={(e) => { if (e.key === 'Enter') sendMessage(); }}
                    style={{ marginRight: '0.5rem' }}
                />
                <button onClick={sendMessage}>Send</button>
            </div>

            {appliances.length > 0 && (
                <div>
                    <h4>📺 Appareils :</h4>
                    {appliances.map((device, idx) => (
                        <fieldset key={idx} style={{ marginBottom: '1rem' }}>
                            <legend>{device.name}</legend>
                            {device.buttons.map((btn, bidx) => (
                                <button
                                    key={bidx}
                                    style={{ marginRight: '0.5rem', marginBottom: '0.5rem' }}
                                    onClick={() => sendIRCode(btn.irCode)}
                                >
                                    {btn.text}
                                </button>
                            ))}
                        </fieldset>
                    ))}
                </div>
            )}

            {Object.entries(outputs).map(([applianceName, outs]) => (
                <div key={applianceName} style={{ marginBottom: '1rem' }}>
                    <h4>📊 État de {applianceName}</h4>
                    <ul>
                        {outs.map(output => (
                            <li key={output.id}>
                                {output.name} :&nbsp;
                                <strong>
                                    {output.type === 'bool'
                                        ? output.value ? '🟢 ON' : '🔴 OFF'
                                        : output.value.toString()}
                                </strong>
                            </li>
                        ))}
                    </ul>
                </div>
            ))}


            <div style={{ border: '1px solid #ccc', padding: '1rem', height: '300px', overflowY: 'auto' }}>
                <h4>📨 Messages :</h4>
                {messages.map((msg, idx) => (
                    <div key={idx}>• {msg}</div>
                ))}
            </div>
        </div>
    );
};

export default WebSocketView;
