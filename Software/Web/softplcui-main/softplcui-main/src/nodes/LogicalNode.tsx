import {getConnectedEdges, NodeProps, Position, useStore} from "reactflow";
import CustomHandle from "./CustomHandle.tsx";
import {useEffect, useMemo} from "react";
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


    let content;
    if (!data.stretchable) {
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
        <div className="react-flow__node-default logicalNode">
            {data.label && <div>{data.label}</div>}
            {content}
        </div>
    );
};


export default LogicalNode;