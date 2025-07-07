import React, { createContext, useContext, useState } from 'react';

export type Tool = 'default' | 'DisplayConnectionDebug' | 'Paint' | 'disconnect';

const emojiMap: Record<Tool, string> = {
    default: '🖱️',
    DisplayConnectionDebug: '🔍',
    Paint: '🔫',
    disconnect: '❌',
};

const ToolContext = createContext<{
    tool: Tool;
    setTool: (t: Tool) => void;
    emoji: string;
}>({
    tool: 'default',
    setTool: () => {},
    emoji: '🖱️',
});

export const ToolProvider = ({ children }: { children: React.ReactNode }) => {
    const [tool, setTool] = useState<Tool>('default');
    const emoji = emojiMap[tool];
    return (
        <ToolContext.Provider value={{ tool, setTool, emoji }}>
            {children}
        </ToolContext.Provider>
    );
};

export const useTool = () => useContext(ToolContext);
