import { useNavigate } from 'react-router-dom';
import React, {useCallback, useEffect, useRef, useState} from 'react';
import {
    addEdge,
    Background,
    BackgroundVariant,
    Controls,
    OnConnect,
    Panel,
    ReactFlow, ReactFlowInstance,
    ReactFlowProvider,
    useEdgesState,
    useNodesState,
} from "reactflow";


import {ToastContainer, toast} from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

import 'reactflow/dist/style.css';

import Sidebar from './Sidebar';
import useKeyboardShortcuts from './hooks/useKeyboardShortcuts';
import {initialNodes, nodeTypes} from './nodes';
import {edgeTypes, initialEdges} from './edges';

export interface NodesData {
    nodes: Array<{
        accordion: string;
        primaryType: string;
        type: string;
        display: string;
        label: string;
        stretchable: boolean;
        services: { friendlyName: string; nameServices: string[] }[];
        subServices: { friendlyName: string; primary: string; secondary: { dataType: string; name: string }[] }[];
        inputHandle: { dataType: string; name: string }[];
        outputHandle: { dataType: string; name: string }[];
    }>;
}

type Node = {
    id: string;
};
type Connection = {
    source: string | null;
    target: string | null;
    sourceHandle: string | null;
    targetHandle: string | null;
};

let id = 0;
const getId = (): string => `${id++}`;


export default function App() {
    const reactFlowWrapper = useRef<HTMLDivElement | null>(null);
    const [nodes, setNodes, onNodesChange] = useNodesState(initialNodes);
    const [edges, setEdges, onEdgesChange] = useEdgesState(initialEdges);
    const [reactFlowInstance, setReactFlowInstance] = useState<ReactFlowInstance | null>(null);
    const [prevEdgeLength, setPrevEdgeLength] = useState(edges.length);
    const [nodesData, setNodesData] = useState<NodesData>({
        nodes: [{
            accordion: "q",
            primaryType: "q",
            type: "q",
            display: "q",
            label: "q",
            stretchable: false,
            services: [],
            subServices: [],
            inputHandle: [],
            outputHandle: [],
        }]
    });
    const nodesDataRef = useRef(nodesData);
    let initialized = false;

    useEffect(() => {
        if (!initialized) {
            fetch('http://localhost:8889/get-description')
                .then(response => response.json())
                .then(data => {
                    setNodesData(data);
                });
            initialized = true;
        }

    }, []);

    useEffect(() => {
        nodesDataRef.current = nodesData;
    }, [nodesData]);


    useEffect(() => {
        if (prevEdgeLength !== edges.length) {
            setPrevEdgeLength(edges.length);
        }
    }, [edges]);

    const onConnect: OnConnect = useCallback(
        (connection) => {
            setEdges((eds) => addEdge({...connection, type: "step"}, eds));
        },
        [setEdges]
    );

    const onDragOver = useCallback((event: React.DragEvent<HTMLDivElement>) => {
        event.preventDefault();
        event.dataTransfer.dropEffect = 'move';
    }, []);

    const onDrop = useCallback(
        (event: React.DragEvent<HTMLDivElement>) => {
            event.preventDefault();
            console.log("Drop detect");
            const subtype = event.dataTransfer.getData('subtype');
            const nodeDisplay = event.dataTransfer.getData('nodeLabel');

            if (reactFlowInstance) {
                const position = reactFlowInstance.screenToFlowPosition({
                    x: event.clientX,
                    y: event.clientY,
                });
                const nodeId = getId();
                const nodeDataWithId = {
                    ...nodesDataRef.current.nodes.find(node => node.display === nodeDisplay),
                    id: nodeId,
                    selectedFriendlyNameData: "default",
                    selectedServiceData: "default",
                    selectedSubServiceData: "default",
                    valueData: "",
                };
                const newNode = {
                    id: nodeId,
                    type: subtype,
                    position,
                    data: nodeDataWithId,
                };
                setNodes((nds) => nds.concat(newNode));
            }

        },
        [reactFlowInstance]

    );

    const onBuild = () => {
        const data = {nodes, edges};
        const selectedElements = {
            nodes: data.nodes.map(node => ({
                id: node.id,
                type: node.data.type,
                data: {
                    friendlyName: node.data.selectedFriendlyNameData,
                    service: node.data.selectedServiceData,
                    subService: node.data.selectedSubServiceData,
                    value: node.data.valueData,
                    parameterValueData : node.data.parameterValueData ?? [],
                    parameterNameData : node.data.parameterNameData ?? [],

                },
            })),
            edges: data.edges.map(edge => ({
                source: edge.source,
                sourceHandle: edge.sourceHandle,
                target: edge.target,
                targetHandle: edge.targetHandle,
            })),
        };
/*
        // download json
        const blob = new Blob([JSON.stringify(selectedElements, null, 2)], {type: 'application/json'});
        const url = URL.createObjectURL(blob);
        const link = document.createElement('a');
        link.href = url;
        link.download = 'data.json';
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
*/
        const myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");

        const raw = JSON.stringify(selectedElements);
        console.log("App raw selectedElements : ",raw)
        const requestOptions: RequestInit = {
            method: "POST",
            headers: myHeaders,
            body: raw,
            redirect: "follow",
        };

        let resp;
        const buildPromise = fetch("http://localhost:8889/json-graph", requestOptions)
            .then((response) => response.text())
            .then((result) => {
                resp = result;
                if (result != "Graph received") {
                    throw new Error(resp);
                }
            })
            .catch((error) => {
                if (error.message === "Failed to fetch") {
                    throw new Error("Server isn't running.")
                }
                throw error; // Rethrow the error to handle it in toast.promise
            });

        toast.promise(
            buildPromise,
            {
                pending: "Building graph...",
                success: "Graph is running on CC100",
                error: {
                    render: ({data}) => {
                        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                        // @ts-expect-error
                        return data.message;
                    },
                    autoClose: 7000
                }
            }
        )

    };

    const onSave = () => {
        if (!reactFlowInstance) {
            return;
        }
        const elements = reactFlowInstance.toObject();
        const json = JSON.stringify(elements);
        const myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");
        const requestOptions: RequestInit = {
            method: "POST",
            headers: myHeaders,
            body: json,
            redirect: "follow",
        };

        const savePromise = fetch("http://localhost:8889/json-save", requestOptions)
            .then((response) => response.text())
            .then((result) => {
                if (result != "Graph saved") {
                    throw new Error("Failed to save graph");
                }
            })
            .catch((error) => {
                if (error.message === "Failed to fetch") {
                    throw new Error("Server isn't running.")
                }
                throw error; // Rethrow the error to handle it in toast.promise
            });

        toast.promise(
            savePromise,
            {
                pending: "Saving graph...",
                success: "Graph saved successfully!",
                error: {
                    render: ({data}) => {
                        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                        // @ts-expect-error
                        return data.message;
                    },
                    autoClose: 7000
                }
            }
        )
    };

    const onRestore = () => {
        const restorePromise = fetch("http://localhost:8889/get-saved-json")
            .then((response) => response.json())
            .then((data) => {
                if (data === null) {
                    throw new Error("No data to restore");
                }
                setNodes(data.nodes);
                setEdges(data.edges);

                // Find the highest id among the nodes
                const highestId = Math.max(...data.nodes.map((node: Node) => Number(node.id)));
                // If highestId is a number (not NaN), set id to highestId + 1
                if (!isNaN(highestId)) {
                    id = highestId + 1;
                }

            })
            .catch((error) => {
                if (error.message === "Failed to fetch") {
                    throw new Error("Server isn't running.")
                }
                throw error; // Rethrow the error to handle it in toast.promise
            });

        toast.promise(
            restorePromise,
            {
                pending: "Restoring graph...",
                success: "Graph restored successfully!",
                error: {
                    render: ({data}) => {
                        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                        // @ts-expect-error
                        return data.message;
                    },
                    autoClose: 7000
                }
            }
        )
    };
    const navigate = useNavigate();

    const openView = () => {
        navigate('/websocket');
    };
    const isValidConnection = (connection: Connection) => {
        const sourceNode = nodes.find((node) => node.id === connection.source);
        const targetNode = nodes.find((node) => node.id === connection.target);
        if (!sourceNode || !targetNode) {
            return false;
        }
        let sourceHandleType;
        let targetHandleType;
        if (sourceNode.data.stretchable) {
            sourceHandleType = sourceNode.data.outputHandle[0].dataType;
        } else if (sourceNode.data.subServices.length > 0) {
            sourceHandleType = sourceNode.data.dataType;
        } else {
            sourceHandleType = sourceNode.data.outputHandle.find((output: {
                name: string;
            }) => output.name === connection.sourceHandle).dataType;
        }
        if (targetNode.data.stretchable) {
            targetHandleType = targetNode.data.inputHandle[0].dataType;
        } else if (targetNode.data.subServices.length > 0) {
            targetHandleType = targetNode.data.dataType;
        } else {
            targetHandleType = targetNode.data.inputHandle.find((input: {
                name: string;
            }) => input.name === connection.targetHandle).dataType;
        }
        if (sourceHandleType !== targetHandleType) {
            return false;
        }
        return true;
    }

    useKeyboardShortcuts({ nodes, edges, setNodes, setEdges, getId });
    return (
        <ReactFlowProvider>
            <div className="dndflow">
                <div className="reactflow-wrapper" ref={reactFlowWrapper} style={{height: '100vh', width: '100%'}}>
                    <ReactFlow

                        className="softPLCFlow"
                        nodes={nodes}
                        nodeTypes={nodeTypes}
                        onNodesChange={onNodesChange}
                        edges={edges}
                        edgeTypes={edgeTypes}
                        onEdgesChange={onEdgesChange}
                        onConnect={onConnect}
                        isValidConnection={isValidConnection}
                        onInit={setReactFlowInstance}
                        onDrop={onDrop}
                        onDragOver={onDragOver}
                    >
                        <Panel position="top-right">
                            <button onClick={openView}>Open view</button>
                            <button onClick={onSave}>Save</button>
                            <button onClick={onRestore}>Restore</button>
                            <button onClick={onBuild}>Build</button>
                        </Panel>
                        <Background
                            color="red"
                            variant={BackgroundVariant.Dots}
                        />
                        <Controls/>
                    </ReactFlow>
                </div>
                <Sidebar nodesData={nodesData}/>
            </div>
            <ToastContainer position="bottom-right"
                            theme="colored"
                            closeOnClick
                            closeButton={false}
                            hideProgressBar={true}
                            autoClose={3000}
                            pauseOnFocusLoss={false}
            />
        </ReactFlowProvider>
    );
}

