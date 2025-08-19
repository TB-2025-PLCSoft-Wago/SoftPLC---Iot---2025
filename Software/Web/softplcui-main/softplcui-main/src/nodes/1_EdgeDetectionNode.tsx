// EdgeDetectionNode.tsx
import React, { useState, useEffect, useMemo } from "react";
import { getParameterElementUsingNumber } from "./utils/getParameterCount.ts";
import FixedHandles from "./handles/FixedHandles.tsx";
import { LogicalNodeData } from "./types.ts";

interface EdgeDetectionNodeProps {
    data: LogicalNodeData;
    numberOfConnectedTargetHandles: number;
}

const EdgeDetectionNode: React.FC<EdgeDetectionNodeProps> = ({ data, numberOfConnectedTargetHandles }) => {
    const [selectedType, setSelectedType] = useState(data.type);

    const onChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        const value = event.target.value;
        if (value === "RFtrigNode") {
            data.label = "RF_trig";
            data.type = "RFtrigNode";
            setSelectedType("RFtrigNode");
        } else if(value === "RtrigNode") {
            data.label = "Rtrig";
            data.type = "RtrigNode";
            setSelectedType("RtrigNode");
        } else {
            data.label = "Ftrig";
            data.type = "FtrigNode";
            setSelectedType("FtrigNode");
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
                        <option value="RFtrigNode">RF_trig</option>
                        <option value="RtrigNode">Rtrig</option>
                        <option value="FtrigNode">Ftrig</option>
                    </select>
                </div>
            )}
            {content}
        </div>
    );
};

export default EdgeDetectionNode;
