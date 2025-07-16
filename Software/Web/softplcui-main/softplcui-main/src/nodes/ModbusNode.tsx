// ModbusNode.tsx
import React, { useState, useEffect, useMemo } from "react";
import { getParameterElementUsingNumber } from "./utils/getParameterCount.ts";
import FixedHandles from "./handles/FixedHandles.tsx";
import { LogicalNodeData } from "./types.ts";
import CommunicationHandles from "./handles/CommunicationHandles.tsx";
import CommunicationHandles_select from "./handles/CommunicationHandles_select.tsx";

interface ModbusNodeProps {
    data: LogicalNodeData;
    inputValues: string[];
    handleInputChange: (index: number) => (e: React.ChangeEvent<HTMLInputElement>) => void;
    setInputValues: React.Dispatch<React.SetStateAction<string[]>>;
    onResize?: (width: number, height: number) => void;
    nodeSize: any;
}

const ModbusNode: React.FC<ModbusNodeProps> = ({ data,
                                                   inputValues,
                                                   handleInputChange,
                                                   setInputValues,
                                                   onResize,
                                                   nodeSize}) => {
    const [selectedType, setSelectedType] = useState(data.type);

    useEffect(() => {
        setSelectedType(data.type);
    }, [data.type]);


    const onChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        const value = event.target.value;
        if (value === "ConfigurableNodeModbusReadBool") {
            data.label = "Modbus Read Bool";
            data.type = "ConfigurableNodeModbusReadBool";
            const newValuesEntry = data.inputHandle.find(h => h.name === "NewValues");
            if (newValuesEntry) {
                data.inputHandle = data.inputHandle.map(h =>
                    h.name === "NewValues" ? { ...newValuesEntry, name: "Quantity" } : h
                );
            }
            setSelectedType("ConfigurableNodeModbusReadBool");
        } else if(value === "ConfigurableNodeModbusReadValue") {
            data.label = "Modbus Read Value";
            data.type = "ConfigurableNodeModbusReadValue";
            const newValuesEntry = data.inputHandle.find(h => h.name === "NewValues");
            if (newValuesEntry) {
                data.inputHandle = data.inputHandle.map(h =>
                    h.name === "NewValues" ? { ...newValuesEntry, name: "Quantity" } : h
                );
            }
            setSelectedType("ConfigurableNodeModbusReadValue");
        }else if(value === "ConfigurableNodeModbusWriteBool") {
            data.label = "Modbus Write Bool";
            data.type = "ConfigurableNodeModbusWriteBool";
            const newValuesEntry = data.inputHandle.find(h => h.name === "Quantity");
            if (newValuesEntry) {
                data.inputHandle = data.inputHandle.map(h =>
                    h.name === "Quantity" ? { ...newValuesEntry, name: "NewValues" } : h
                );
            }
            setSelectedType("ConfigurableNodeModbusWriteBool");
        }
        else if(value === "ConfigurableNodeModbusWriteValue") {
            data.label = "Modbus Write Value";
            data.type = "ConfigurableNodeModbusWriteValue";
            const newValuesEntry = data.inputHandle.find(h => h.name === "Quantity");
            if (newValuesEntry) {
                data.inputHandle = data.inputHandle.map(h =>
                    h.name === "Quantity" ? { ...newValuesEntry, name: "NewValues" } : h
                );
            }
            setSelectedType("ConfigurableNodeModbusWriteValue");
        }
    };

    const content = <CommunicationHandles_select
                data={data}
                inputValues={inputValues}
                handleInputChange={handleInputChange}
                setInputValues={setInputValues}
                onResize={onResize}
            />;


    return (
        <div
            className="react-flow__node-default logicalNode"
            style={{
                ...nodeSize,
                position: "relative",
            }}
        >
            {data.label && (
                <div className="data-label dl-Communication">
                    <select className="custom-select" value={selectedType} onChange={onChange}>
                        <option value="ConfigurableNodeModbusReadBool">Modbus Read Bool (0x02)</option>
                        <option value="ConfigurableNodeModbusReadValue">Modbus Read Value (0x04)</option>
                        <option value="ConfigurableNodeModbusWriteBool">Modbus Write Bool (0x15)</option>
                        <option value="ConfigurableNodeModbusWriteValue">Modbus Write Value (0x06)</option>
                    </select>
                </div>
            )}
            {content}
        </div>
    );
};

export default ModbusNode;
