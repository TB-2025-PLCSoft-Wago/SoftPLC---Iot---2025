import React, { useState, useRef, useEffect } from 'react';

interface TextInputSenderProps {
    irCode: number;
    ws: React.MutableRefObject<WebSocket | null>;
}

const TextInputSender: React.FC<TextInputSenderProps> = ({ irCode, ws }) => {
    const [value, setValue] = useState('');

    const sendValue = () => {
        if (!value.trim()) return;

        const payload = {
            irCode,
            value,
        };

        if (ws.current?.readyState === WebSocket.OPEN) {
            ws.current.send(JSON.stringify(payload));
            setValue('');
        } else {
            alert('WebSocket not connected.');
        }
    };

    return (
        <>
            <input
                type="text"
                value={value}
                placeholder="enter a value"
                onChange={(e) => setValue(e.target.value)}
                onKeyDown={(e) => {
                    if (e.key === 'Enter') sendValue();
                }}
                style={{ marginRight: '0.5rem' }}
            />
            <button onClick={sendValue}>Send</button>
        </>
    );
};

export default TextInputSender;
