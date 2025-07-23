import React, { useEffect, useRef, useState } from 'react';
import TextInputSender from "./TextInputSender.tsx";
import {useNavigate} from "react-router-dom";
import CompressedMessageList from "./CompressedMessageList.tsx";

type Appliance = {
    name: string;
    inputs: {
        textInput: boolean;
        text: string; irCode: number }[];
};
type Output = {
    id: number;
    name: string;
    applianceName: string;
    type: string; // "bool", "value", "float"
    value: string;
};

const WebSocketView = () => {
    const [messages, setMessages] = useState<string[]>([]);
    const [timestamps, setTimestamps] = useState<number[]>([]);
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
                    if (data.edges) { /* empty */ }else {
                        setMessages((prev) => [...prev, event.data]);
                        setTimestamps((prev) => [...prev, Date.now()]); //capture timestamp
                    }

                }
            } catch {
                setMessages((prev) => [...prev, event.data]);
                setTimestamps((prev) => [...prev, Date.now()]);
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
    const navigate = useNavigate();
    const goBackView = () => {
        navigate(-1);
    };
    //console.log("appliances:",appliances);
    return (
        <>
            {/* button */}
            <div style={{
                position: 'fixed',
                top: '2rem',
                right: '2rem',
                zIndex: 1000
            }}>
                <button className={"button button1"} onClick={goBackView}>Close view</button>
            </div>

            {/* main content
 */}
        <div style={{height: '100vh', overflowY: 'auto', padding: '1rem'}}>
            <h2>User Interface</h2>
            <div style={{marginBottom: '1rem'}}>
                <input
                    type="text"
                    placeholder="Enter message"
                    value={input}
                    onChange={(e) => setInput(e.target.value)}
                    onKeyDown={(e) => {
                        if (e.key === 'Enter') sendMessage();
                    }}
                    style={{marginRight: '0.5rem'}}
                />
                <button onClick={sendMessage}>Send</button>
            </div>
            {appliances.length >= 0 && (
                <div>
                    {/*<h4>Devices :</h4>*/}
                    {appliances.map((device, idx) => {
                        const deviceOutputs = outputs[device.name] || [];

                        return (
                            <fieldset key={idx} style={{marginBottom: '1rem'}}>
                                <legend>{device.name}</legend>

                                {/* Command Inputs*/}
                                {device.inputs.map((btn, bidx) => {
                                    if (btn.textInput) {
                                        return (
                                            <div key={bidx} style={{marginBottom: '0.5rem'}}>
                                                <label style={{marginRight: '0.5rem'}}>{btn.text}:</label>
                                                <TextInputSender irCode={btn.irCode} ws={ws}/>
                                            </div>
                                        );
                                    }

                                    return (
                                        <button className={"button button1"}
                                            key={bidx}
                                            style={{marginRight: '0.5rem', marginBottom: '0.5rem'}}
                                            onClick={() => sendIRCode(btn.irCode)}
                                        >
                                            {btn.text}
                                        </button>
                                    );
                                })}


                                {/* Outputs (states) */}
                                {deviceOutputs.length > 0 && (
                                    <div style={{marginTop: '0.5rem'}}>
                                        <h5>📊 States :</h5>
                                        <ul style={{marginTop: '0.2rem'}}>
                                            {deviceOutputs.map(output => (
                                                <li key={output.id}>
                                                    {output.name} :{' '}
                                                    <strong>
                                                        {output.type === 'bool'
                                                            ? output.value
                                                                ? '🟢 ON'
                                                                : '🔴 OFF'
                                                            : output.value.toString()}
                                                    </strong>
                                                </li>
                                            ))}
                                        </ul>
                                    </div>
                                )}
                            </fieldset>
                        );
                    })}

                    {/* Outputs (without associated button) */}
                    {Object.entries(outputs)
                        .filter(([applianceName]) => !appliances.some(dev => dev.name === applianceName))
                        .map(([applianceName, deviceOutputs]) => (
                            <fieldset key={applianceName} style={{marginBottom: '1rem', borderColor: 'gray'}}>
                                <legend>{applianceName}</legend>

                                <div style={{marginTop: '0.5rem'}}>
                                    <h5>📊 States :</h5>
                                    <ul style={{marginTop: '0.2rem'}}>
                                        {deviceOutputs.map(output => (
                                            <li key={output.id}>
                                                {output.name} :{' '}
                                                <strong>
                                                    {output.type === 'bool'
                                                        ? output.value
                                                            ? '🟢 ON'
                                                            : '🔴 OFF'
                                                        : output.value.toString()}
                                                </strong>
                                            </li>
                                        ))}
                                    </ul>
                                </div>
                            </fieldset>
                        ))}
                </div>
            )}


            <CompressedMessageList messages={messages} timestamps={timestamps} />

        </div>
        </>
    );
};

export default WebSocketView;
