import React, { useRef } from 'react';
import { useTool } from './ToolContext.tsx';
import './ToolsMenu.css';

const toolLabels: Record<string, string> = {
    default: '🖱️ Default',
    DisplayConnectionDebug: '🔍 Display connection (debug)',
    Paint: '🔫 Paint connection',
    comment : '📝 add comment'
};

const ToolsMenu = () => {
    const { tool, setTool, emoji } = useTool();
    const selectRef = useRef<HTMLSelectElement>(null);

    return (
        <div className="dropdown">
            <div className="custom-select-container" onClick={() => selectRef.current?.focus()}>
                <span className="custom-label">🛠 Tool:</span>
                <span className="selected-emoji">{emoji}</span>
                <select
                    ref={selectRef}
                    className="native-select-overlay"
                    value={tool}
                    onChange={(e) => setTool(e.target.value as any)}
                >
                    {Object.entries(toolLabels).map(([value, label]) => (
                        <option key={value} value={value}>
                            {label}
                        </option>
                    ))}
                </select>
            </div>
        </div>
    );
};

export default ToolsMenu;
