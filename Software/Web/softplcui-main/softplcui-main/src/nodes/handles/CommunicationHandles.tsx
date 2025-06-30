import React, { useState } from "react";
import { Position } from "reactflow";
import CustomHandle from "../CustomHandle";
import { LogicalNodeData } from "../types.ts";

interface Props {
    data: LogicalNodeData;
    inputValues: string[];
    handleInputChange: (index: number) => (e: React.ChangeEvent<HTMLInputElement>) => void;
    setInputValues: React.Dispatch<React.SetStateAction<string[]>>;
    onResize?: (width: number, height: number) => void;
}

const CommunicationHandles: React.FC<Props> = ({
                                                   data,
                                                   inputValues,
                                                   handleInputChange,
                                                   setInputValues,
                                                   onResize
                                               }) => {
    const [showConfig, setShowConfig] = useState(false);
    const handleOpenConfig = () => {
        setShowConfig(true);
        onResize?.(550, (data.inputHandle.length + 2) * 40); // setting size
    };

    const handleCloseConfig = () => {
        setShowConfig(false);
        onResize?.(225, (data.inputHandle.length + 2) * 40); // normal size
    };

    return (
        <>
            {/* colored background above the line */}
            <div className="node-top-background ntb-Communication" />

            {data.label && <div className="data-label dl-Communication">{data.label}</div>}

            {/* line of separation */}
            <div className="node-separator ns-Communication"/>

            {/*settings button */}
            <button className={"buttonNode"}
                    onClick={handleOpenConfig}
                    style={{
                        position: "absolute",
                        top: "20px",
                        right: "5px",
                        transform: "translateY(-50%)",
                        zIndex: 10
                    }}
            >
                settings
            </button>

            {/* settings pannel */}
            {showConfig && (
                <div className="config-panel">
                    <h4>Settings configuration</h4>

                    <div className="config-inputs">
                        {inputValues.map((val, index) => (
                            <div key={index} style={{marginBottom: "8px"}}>
                                <input
                                    type="text"
                                    value={val}
                                    onChange={handleInputChange(index)}
                                    placeholder={
                                        data.parameterNameData?.[index] === "setting" || !data.parameterNameData?.[index]
                                            ? `setting ${index - data.parameterNameData.length + 1}`
                                            : data.parameterNameData[index]
                                    }
                                    style={{padding: "4px", width: "80%"}}
                                />
                            </div>
                        ))}
                    </div>

                    <div className="config-buttons">
                        <button onClick={() => setInputValues([...inputValues, ""])} style={{marginRight: "10px"}}>
                            + Add a setting
                        </button>
                        <button onClick={handleCloseConfig}>Close</button>
                    </div>
                </div>
            )}

            {/* input */}
            {data.inputHandle.map((input, index) => (
                <React.Fragment key={index}>
                    <CustomHandle
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
                        {!showConfig && <div className="inputhandletext">{input.name}</div>}
                    </CustomHandle>
                </React.Fragment>
            ))}

            {/* output */}
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
                    {!showConfig && <div className="outputhandletext">{output.name}</div>}
                </CustomHandle>
            ))}

        </>
    );
};

export default CommunicationHandles;
