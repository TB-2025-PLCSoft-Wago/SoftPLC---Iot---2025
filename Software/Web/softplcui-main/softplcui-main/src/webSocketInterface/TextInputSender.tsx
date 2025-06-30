import React, { useState, useRef, useEffect } from 'react';

interface TextInputSenderProps {
    irCode: number;
    ws: React.MutableRefObject<WebSocket | null>;
}

const TextInputSender: React.FC<TextInputSenderProps> = ({ irCode, ws }) => {
    const [value, setValue] = useState('');
    const [isSent, setIsSent] = useState(false);

    const sendValue = () => {
        const payload = {
            irCode,
            value,
        };

        if (ws.current?.readyState === WebSocket.OPEN) {
            ws.current.send(JSON.stringify(payload));
            setIsSent(true); // Mark as sent
        } else {
            alert('WebSocket not connected.');
        }
    };

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setValue(e.target.value);
        if (isSent) {
            setIsSent(false); // Remove the italic as soon as we modify
        }
    };

    return (
        <>
            <input
                type="text"
                value={value}
                placeholder="enter a value"
                onChange={handleChange}
                onKeyDown={(e) => {
                    if (e.key === 'Enter') sendValue();
                }}
                style={{
                    marginRight: '0.5rem',
                    fontStyle: isSent ? 'italic' : 'normal',
                }}
            />
            <button onClick={sendValue}>Send</button>
        </>
    );
};

export default TextInputSender;
