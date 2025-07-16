import React from "react";
import { Position } from "reactflow";
import CustomHandle from "../CustomHandle";
import { LogicalNodeData } from "../types.ts";

interface Props {
    data: LogicalNodeData;
}

const FixedHandles: React.FC<Props> = ({ data }) => {
    return (
        <>
            {data.label && !['TON', 'TOF', 'Ftrig','Rtrig', 'RF_trig', 'Modbus Read Bool', 'Modbus Read Value'].includes(data.label) && (
                <>
                    {/* colored background above the line */}
                    <div className="node-top-background ntb-fixed" />

                    <div className="data-label dl-fixed">{data.label}</div>

                    {/* line of separation */}
                    <div className="node-separator ns-fixed" />
                </>
            )}

            {data.inputHandle.map((input, index) => (
                <CustomHandle
                    key={index}
                    type="target"
                    position={Position.Left}
                    id={input.name}
                    datatype={input.dataType}
                    isConnectable={1}
                    className="inputhandleClass"
                    style={{
                        height: 8,
                        width: 8,
                        top: `${(index + 1) * 100 / (data.inputHandle.length + 1)}%`
                    }}
                >
                    <div className="inputhandletext">{input.name}</div>
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

export default FixedHandles;
