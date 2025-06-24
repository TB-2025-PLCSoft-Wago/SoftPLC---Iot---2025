import React, { useEffect, useMemo, useState } from "react";
import { getConnectedEdges, NodeProps, useStore } from "reactflow";
import { useUpdateNodeInternals } from "reactflow";

import BoolToStringHandles from "./handles/BoolToStringHandles.tsx";
import FixedHandles from "./handles/FixedHandles.tsx";
import StretchableHandles from "./handles/StretchableHandles.tsx";

import { LogicalNodeData } from "./types.ts";
import { getParameterElementUsingNumber } from "./utils/getParameterCount.ts";
import CommunicationHandles from "./handles/CommunicationHandles.tsx";
import StringToBoolHandles from "./handles/StringToBoolHandles.tsx";
import TimerNode from "./TimerNode.tsx";
import EdgeDetectionNode from "./EdgeDetectionNode.tsx";


const LogicalNode: React.FC<NodeProps<LogicalNodeData>> = (props) => {
    const { data = { inputHandle: [], outputHandle: [], stretchable: false, label: "BUG", id: "BUG",
        parameterValueData: undefined,
        parameterNameData: undefined,
        type: undefined
    } } = props;

    const edges = useStore((state) => state.edges);
    const updateNodeInternals = useUpdateNodeInternals();

    const [inputValues, setInputValues] = useState<string[]>(data.parameterValueData ?? []);

    const numberOfConnectedTargetHandles = useMemo(() => {
        if (data.stretchable) {
            const connectedEdges = getConnectedEdges([props], edges);
            return connectedEdges.filter((e) => e.target === props.id).length;
        }
        return -1;
    }, [data.stretchable, data.inputHandle, props.id, edges]);

    useEffect(() => {
        updateNodeInternals(props.id);
    }, [numberOfConnectedTargetHandles]);

    useEffect(() => {
        const baseLength = data.parameterNameData?.length || 0;

        const dynamicParamCount = Math.max(
            numberOfConnectedTargetHandles + 1,
            getParameterElementUsingNumber(data.parameterValueData ?? [])
        ) + 1;

        // We want at least as many fields as parameterized or dynamically required names
        const totalParams = Math.max(baseLength, dynamicParamCount);

        const initialValues = Array.from(
            { length: totalParams },
            (_, i) => inputValues[i] || ""
        );

        setInputValues(initialValues);
        data.parameterValueData = initialValues;

        console.log("Champs init : ", initialValues);
    }, [
        numberOfConnectedTargetHandles,
        data.parameterNameData?.length,
    ]);


    const handleInputChange = (index: number) => (event: React.ChangeEvent<HTMLInputElement>) => {
        const newValues = [...inputValues];
        newValues[index] = event.target.value;
        setInputValues(newValues);
        data.parameterValueData = newValues;
        console.log("newValues : ",newValues)
    };

    let content;
    if (data.label === "bool to string") {
        content = (
            <BoolToStringHandles
                data={data}
                numberOfConnectedTargetHandles={numberOfConnectedTargetHandles}
                inputValues={inputValues}
                handleInputChange={handleInputChange}
            />
        );
    }else if (data.label === "string to bool") {
        //console.log("stoB type : ", data.label);
        //console.log("stoB type : ", data.type);
        content = (
            <StringToBoolHandles
                data={data}
                numberOfConnectedTargetHandles={numberOfConnectedTargetHandles}
                inputValues={inputValues}
                handleInputChange={handleInputChange}
            />
        );
    }else if (data.type.includes("ConfigurableNode")) {
        console.log("ConfigurableNode parameterNameData : ", data.parameterNameData);
        //console.log("Mqtt type : ", data.type);
        const [nodeSize, setNodeSize] = useState({ width: 225, height: (numberOfConnectedTargetHandles + 7) * 40 });

        const handleResize = (width: number, height: number) => {
            setNodeSize({ width, height });
        };
        content = (
            <CommunicationHandles
                data={data}
                inputValues={inputValues}
                handleInputChange={handleInputChange}
                setInputValues={setInputValues}
                onResize={handleResize}
            />
        );
        return (
            <div
                className="react-flow__node-default logicalNode"
                style={{
                    ...nodeSize,
                    position: "relative",
                }}
            >
                {data.label && <div className="data-label">{data.label}</div>}
                {content}
            </div>
        );
    } else if (!data.stretchable) {
        content = <FixedHandles data={data} />;
        if (data.type.includes("TON") || (data.type.includes("TOF"))) {
            //console.log("timer : ",data.type)
            return (
                <TimerNode
                    data={data}
                    numberOfConnectedTargetHandles={numberOfConnectedTargetHandles}
                />
            );

        }
        if (data.type.includes("trig")) {
            //console.log("timer : ",data.type)
            return (
                <EdgeDetectionNode
                    data={data}
                    numberOfConnectedTargetHandles={numberOfConnectedTargetHandles}
                />
            );

        }
    } else {
        content = <StretchableHandles data={data} numberOfConnectedTargetHandles={numberOfConnectedTargetHandles}/>;


    }

    return (
        <div
            className="react-flow__node-default logicalNode"
            style={{
                height: `${(Math.max(numberOfConnectedTargetHandles, getParameterElementUsingNumber(data.parameterValueData ?? [])) + 3) * 40}px`,
                width: (data.label === "bool to string"|| data.label ==="string to bool") ? "225px" : "80px",
                position: "relative",
            }}
        >
            {data.label && <div>{data.label}</div>}
            {content}
        </div>
    );
};

export default LogicalNode;
