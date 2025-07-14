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
    const [parameterNames, setParameterNames] = useState<string[]>(data.parameterNameData || []);
    console.log("parameterNames : ",parameterNames)

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
                            {/* show settings, choose table parameterNames or inputValues */}
                            {(parameterNames.length < inputValues.length ? inputValues : parameterNames).map((val, index) => (
                                <div key={index} style={{ marginBottom: "8px" }}>
                                    <input
                                        type="text"
                                        value={inputValues[index] || ""}
                                        onChange={handleInputChange(index)}
                                        placeholder={
                                            data.parameterNameData?.[index] === "setting" || !data.parameterNameData?.[index]
                                                ? parameterNames[index] || `setting ${index + 1}`
                                                : data.parameterNameData[index]
                                        }
                                        style={{ padding: "4px", width: "80%" }}
                                    />
                                </div>
                            ))}
                    </div>

                    <div className="config-buttons">
                        {data.label === "HTTP Client" && (
                            <>
                                <button
                                    className="buttonNode"
                                    onClick={() => {
                                        setInputValues(prev => [...prev, ""]);
                                        setParameterNames(prev => [
                                            ...prev,
                                            `header Key ${Math.floor((prev.length)/2)}`,
                                            `header Value ${Math.floor((prev.length)/2)}`
                                        ]);
                                    }}
                                    style={{marginRight: "10px"}}
                                >
                                    + Add a header
                                </button>
                                <button
                                    className="buttonNode"
                                    onClick={() => {
                                        setInputValues(prev => prev.slice(0, -2));
                                        setParameterNames(prev => prev.slice(0, -2));
                                    }}
                                    disabled={inputValues.length <= data.parameterNameData.length && parameterNames.length <= data.parameterNameData.length}
                                    style={{marginRight: "10px"}}
                                >
                                    − Remove last header
                                </button>

                            </>
                        )}
                        <button className="buttonNode closeSetting" onClick={handleCloseConfig}>Close</button>
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
