import React from "react";
import { Position } from "reactflow";
import CustomHandle from "../CustomHandle";
import { LogicalNodeData } from "../types.ts";
import { getParameterElementUsingNumber } from "../utils/getParameterCount.ts";

interface Props {
    data: LogicalNodeData;
    numberOfConnectedTargetHandles: number;
    inputValues: string[];
    handleInputChange: (index: number) => (e: React.ChangeEvent<HTMLInputElement>) => void;
}

const StringToBoolHandles: React.FC<Props> = ({
                                                  data,
                                                  numberOfConnectedTargetHandles,
                                                  inputValues,
                                                  handleInputChange,
                                              }) => {
    const maxHandles = Math.max(
        numberOfConnectedTargetHandles + 1,
        getParameterElementUsingNumber(data.parameterValueData ?? []) + 1
    );

    return (

        <>
        {data.inputHandle.map((input, index) => (
            <CustomHandle
                key={index}
                type="target"
                position={Position.Left}
                id={input.name}
                datatype={input.dataType}
                isConnectable={1}
                style={{
                    height: 8,
                    width: 8,
                    top: `${(index + 1) * 100 / (data.inputHandle.length + 1)}%`
                }}
            >
                <div className="inputhandletext">{input.name}</div>
            </CustomHandle>
        ))}

        {Array.from({length: maxHandles}).map((_, index) => (
            <React.Fragment key={index}>
                <CustomHandle
                    key={index}
                    type="source"
                    position={Position.Right}
                    id={`${data.outputHandle[0].name}${index}`}
                    datatype={data.outputHandle[0].dataType}
                    style={{
                        height: 8,
                        width: 8,
                        top: `${(index + 1) * 100 / (maxHandles + 1)}%`
                    }}
                >
                    <div className="outputhandletext">{`${data.outputHandle[0].name}${index}`}</div>
                </CustomHandle>

                <input
                    type="text"
                    className="inputNodeSelect"
                    value={inputValues[index] || ""}
                    onChange={handleInputChange(index)}
                    id={`${data.id}-input-${index}`}
                    style={{
                        position: "absolute",
                        top: `${(index + 1) * 100 / (maxHandles + 1)}%`,
                        left: "20px",
                    }}
                />
            </React.Fragment>
        ))}

</>
)
    ;
};

export default StringToBoolHandles;
