// TimerNode.tsx
import React, { useState, useEffect, useMemo } from "react";
import { getParameterElementUsingNumber } from "./utils/getParameterCount.ts";
import FixedHandles from "./handles/FixedHandles.tsx";
import { LogicalNodeData } from "./types.ts";

interface TimerNodeProps {
    data: LogicalNodeData;
    numberOfConnectedTargetHandles: number;
}

const TimerNode: React.FC<TimerNodeProps> = ({ data, numberOfConnectedTargetHandles }) => {
    const [selectedType, setSelectedType] = useState(data.type);

    const onChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        const value = event.target.value;
        if (value === "TOFNode") {
            data.label = "TOF";
            data.type = "TOFNode";
            setSelectedType("TOFNode");
        } else {
            data.label = "TON";
            data.type = "TONNode";
            setSelectedType("TONNode");
        }
    };

    const content = <FixedHandles data={data} />;

    return (
        <div
            className="react-flow__node-default logicalNode"
            style={{
                height: `${(Math.max(numberOfConnectedTargetHandles, getParameterElementUsingNumber(data.parameterValueData ?? [])) + 3) * 40}px`,
                width: "80px",
                position: "relative",
            }}
        >
            {data.label && (
                <div>
                    <select className="custom-select" value={selectedType} onChange={onChange}>
                        <option value="TOFNode">TOF</option>
                        <option value="TONNode">TON</option>
                    </select>
                </div>
            )}
            {content}
        </div>
    );
};

export default TimerNode;
