import React from "react";
import { Position } from "reactflow";
import CustomHandle from "../CustomHandle";
import { LogicalNodeData } from "../types.ts";

interface Props {
    data: LogicalNodeData;
    numberOfConnectedTargetHandles: number;
}

const StretchableHandles: React.FC<Props> = ({ data, numberOfConnectedTargetHandles }) => {
    const totalHandles = numberOfConnectedTargetHandles + 1;

    return (
        <>
            {/* colored background above the line */}
            <div className="node-top-background ntb-stretchable" />

            {data.label && <div className="data-label dl-stretchable">{data.label}</div>}

            {/* line of separation */}
            <div className="node-separator ns-stretchable"/>
            {Array.from({ length: totalHandles }).map((_, index) => (
                <CustomHandle
                    key={index}
                    type="target"
                    position={Position.Left}
                    id={`${data.inputHandle[0].name}${index}`}
                    datatype={data.inputHandle[0].dataType}
                    isConnectable={1}
                    className="inputhandleClass"
                    style={{
                        height: 8,
                        width: 8,
                        top: `${(index + 1) * 100 / (totalHandles + 1)}%`
                    }}
                >
                    <div className="inputhandletext">{`${data.inputHandle[0].name}${index}`}</div>
                </CustomHandle>
            ))}

            {data.outputHandle.map((output, index) => (
                <CustomHandle
                    key={index}
                    type="source"
                    position={Position.Right}
                    id={output.name}
                    datatype={output.dataType}
                    className="inputhandleClass"
                    style={{
                        height: 8,
                        width: 8,
                        top: `${(index + 1) * 100 / (data.outputHandle.length + 1)}%`
                    }}
                >
                    <div className="outputhandletext">{output.name}</div>
                </CustomHandle>
            ))}
        </>
    );
};

export default StretchableHandles;
