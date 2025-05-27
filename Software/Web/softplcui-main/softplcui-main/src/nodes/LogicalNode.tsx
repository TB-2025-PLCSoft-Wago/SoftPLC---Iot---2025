import {getConnectedEdges, Handle, NodeProps, Position, useStore} from "reactflow";
import CustomHandle from "./CustomHandle.tsx";
import React, {useEffect, useMemo, useState} from "react";
import {useUpdateNodeInternals} from "reactflow";

interface LogicalNodeData {
    inputHandle: { dataType: string; name: string }[];
    outputHandle: { dataType: string; name: string }[];
    stretchable: boolean;
    label: string;
    id: string;
    selectedFriendlyNameData?: string;
    selectedServiceData?: string;
    selectedSubServiceData?: string;
    valueData?: string;
    dataType?: string;
    parameterValueData?: string[];
}


const LogicalNode: React.FC<NodeProps<LogicalNodeData>> = (props) => {
    const {data = {inputHandle: [], outputHandle: [], stretchable: false, label: "BUG", id: "BUG",}} = props;
    data.selectedServiceData = "";
    data.selectedSubServiceData = "";
    data.valueData = "";
    data.selectedFriendlyNameData = "";

    const edges = useStore((state) => state.edges);
    const updateNodeInternals = useUpdateNodeInternals();

    const numberOfConnectedTargetHandles = useMemo(() => {
        if (data.stretchable) {
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-expect-error
            const connectedEdges = getConnectedEdges([props], edges);
            const connectedTargetHandles = connectedEdges
                .filter(edge => edge.target === props.id)
                .map(edge => edge.targetHandle);
            return connectedTargetHandles.length;
        }
        return -1;
    }, [data.stretchable, data.inputHandle, props.id, edges]);

    useEffect(() => {
        updateNodeInternals(props.id);
    }, [numberOfConnectedTargetHandles]);

    useEffect(() => {
        const initialValues = Array.from({ length: Math.max(numberOfConnectedTargetHandles + 1,getParameterElementUsingNumber(data.parameterValueData ?? [])+1)}, (_, i) => inputValues[i] || "");
        setInputValues(initialValues);
        data.parameterValueData = initialValues;
        console.log("parameterValueData initialValues : ", initialValues);
    }, [numberOfConnectedTargetHandles]);

    const [inputValues, setInputValues] = useState<string[]>(data.parameterValueData ?? []);
    const handleInputChange = (index: number) => (event: React.ChangeEvent<HTMLInputElement>) => {
        const newValues = [...inputValues];
        newValues[index] = event.target.value;
        setInputValues(newValues);
        data.parameterValueData = newValues; // update data.parameterValueData
        console.log("handleInputChange newValues : ", newValues);
    };
    const getParameterElementUsingNumber = (arr: string[]): number => {
        let count : number = 0;
        let countEmptyBetween  : number = 0;
        arr.forEach((str, index) => {
            //console.log(`Index ${index}: ${count}`);
            if (!str){
                countEmptyBetween += 1;
            }else{
                count += 1 + countEmptyBetween;
                countEmptyBetween = 0;
            }
        });
        console.log("Count parameter : ", count);
        return count;
    }

    let content;

    //console.log("logical node parameter value data :", data.parameterValueData);
    if (data.label === "bool to string") {
        //console.log("BToSNode logical node");
        content = (
            <>
                {Array.from({length: Math.max(numberOfConnectedTargetHandles + 1,getParameterElementUsingNumber(data.parameterValueData ?? [])+1)}).map((_, index) => (
                    <React.Fragment key={index}>
                        <CustomHandle
                            key={index}
                            type="target"
                            position={Position.Left}
                            id={data.inputHandle[0].name + index}
                            datatype={data.inputHandle[0].dataType}
                            isConnectable={1}
                            style={{
                                height: 8,
                                width: 8,
                                top: `${(index + 1) * 100 / (Math.max(numberOfConnectedTargetHandles,getParameterElementUsingNumber(data.parameterValueData ?? [])) + 2)}%`
                            }}
                        >
                            <div className="inputhandletext">{data.inputHandle[0].name + index}</div>
                        </CustomHandle>
                        <input
                            type="text"
                            className="inputNodeSelect"
                            value={inputValues[index] || ""}
                            onChange={handleInputChange(index)}
                            id={`${data.id}-input-${index}`} // Rends l'ID unique
                            style={{ position: 'absolute', top: `${(index + 1) * 100 / (Math.max(numberOfConnectedTargetHandles,getParameterElementUsingNumber(data.parameterValueData ?? [])) + 2)}%`, left: '20px' }}
                        />
                    </React.Fragment>
                ))}
                {Array.from(data.outputHandle).map((output, index) => (
                    <CustomHandle
                        key={index}
                        type="source"
                        position={Position.Right}
                        id={output.name}
                        datatype={output.dataType}
                        style={{height: 8, width: 8, top: `${(index + 1) * 100 / (data.outputHandle.length + 1)}%`}}
                    >
                        <div className="outputhandletext">{output.name}</div>
                    </CustomHandle>
                ))}
            </>
        );
    }else if (!data.stretchable) {
        content = (
            <>
                {Array.from(data.inputHandle).map((input, index) => (
                    <CustomHandle
                        key={index}
                        type="target"
                        position={Position.Left}
                        id={input.name}
                        datatype={input.dataType}
                        isConnectable={1}
                        style={{height: 8, width: 8, top: `${(index + 1) * 100 / (data.inputHandle.length + 1)}%`}}
                    >
                        <div className="inputhandletext">{input.name}</div>
                    </CustomHandle>
                ))}
                {Array.from(data.outputHandle).map((output, index) => (
                    <CustomHandle
                        key={index}
                        type="source"
                        position={Position.Right}
                        id={output.name}
                        datatype={output.dataType}
                        style={{height: 8, width: 8, top: `${(index + 1) * 100 / (data.outputHandle.length + 1)}%`}}
                    >
                        <div className="outputhandletext">{output.name}</div>
                    </CustomHandle>
                ))}
            </>

        );
    } else {
        content = (
            <>
                {Array.from({length: numberOfConnectedTargetHandles + 1}).map((_, index) => (
                    <CustomHandle
                        key={index}
                        type="target"
                        position={Position.Left}
                        id={data.inputHandle[0].name + index}
                        datatype={data.inputHandle[0].dataType}
                        isConnectable={1}
                        style={{
                            height: 8,
                            width: 8,
                            top: `${(index + 1) * 100 / (numberOfConnectedTargetHandles + 2)}%`
                        }}
                    >
                        <div className="inputhandletext">{data.inputHandle[0].name + index}</div>
                    </CustomHandle>
                ))}
                {Array.from(data.outputHandle).map((output, index) => (
                    <CustomHandle
                        key={index}
                        type="source"
                        position={Position.Right}
                        id={output.name}
                        datatype={output.dataType}
                        style={{height: 8, width: 8, top: `${(index + 1) * 100 / (data.outputHandle.length + 1)}%`}}
                    >
                        <div className="outputhandletext">{output.name}</div>
                    </CustomHandle>
                ))}
            </>
        );
    }

    return (
        <div
            className="react-flow__node-default logicalNode"
            style={{
                height: `${(Math.max(numberOfConnectedTargetHandles,getParameterElementUsingNumber(data.parameterValueData ?? []))+ 3) * 40}px`,
                width: data.label === "bool to string" ? "225px" : "80px",
                position: 'relative',
            }}
        >
            {data.label && <div>{data.label}</div>}
            {content}
        </div>
    );
};


export default LogicalNode;