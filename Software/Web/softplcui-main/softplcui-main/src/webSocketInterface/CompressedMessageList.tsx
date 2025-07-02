import React, { useMemo } from 'react';

const CompressedMessageList = ({ messages, timestamps }: { messages: string[]; timestamps: number[] }) => {
    const processedMessages = useMemo(() => {
        const result = [];
        let lastMsg = null;
        let lastTime = null;
        let count = 0;

        for (let i = 0; i < messages.length; i++) {
            const msg = messages[i];
            const time = timestamps[i];

            if (msg === lastMsg) {
                count++;
            } else {
                if (lastMsg !== null) {
                    const timeDiff =
                        lastTime != null && time != null
                            ? `${Math.round((time - lastTime) / 1000)}s`
                            : null;

                    result.push({
                        text: lastMsg,
                        count,
                        timeDiff,
                    });
                }
                lastMsg = msg;
                lastTime = time;
                count = 1;
            }
        }

        // last message
        if (lastMsg !== null) {
            result.push({
                text: lastMsg,
                count,
                timeDiff: null,
            });
        }

        return result;
    }, [messages, timestamps]);


    return (
        <div style={{ border: '1px solid #ccc', padding: '1rem', height: '300px', overflowY: 'auto' }}>
            <h4>📨 Messages :</h4>
            {processedMessages.map((entry, idx) => (
                <div key={idx}>
                    • {entry.text}
                    {entry.count > 1 && ` (×${entry.count})`}
                    {/*entry.timeDiff && <span style={{ color: '#888' }}> – {entry.timeDiff} later</span>*/}
                </div>
            ))}
        </div>
    );
};

export default CompressedMessageList;
