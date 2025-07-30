import { useNavigate } from 'react-router-dom';
import React, {useCallback, useEffect, useRef, useState} from 'react';
import {
    addEdge,
    Background,
    BackgroundVariant,
    Controls, MarkerType,
    OnConnect,
    Panel,
    ReactFlow, ReactFlowInstance,
    ReactFlowProvider,
    useEdgesState,
    useNodesState,
    Edge, MiniMap,
} from "reactflow";


import {ToastContainer, toast} from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

import 'reactflow/dist/style.css';

import Sidebar from './Sidebar';
import useKeyboardShortcuts from './hooks/useKeyboardShortcuts';
import {initialNodes, nodeTypes} from './nodes';
import {edgeTypes, initialEdges} from './edges';
import CustomEdgeStartEndDebug  from "./edges/CustomEdgeStartEndDebug.tsx";
import CustomEdgeStepControl from "./CustomEdgeStepControl.tsx";
/*color*/
import ConnectionLine from './nodes/utils/ConnectionLine.tsx';
import ColorSelectorNode from './nodes/utils/ColorSelectorNode';
import Debug from "./webSocketInterface/debug.tsx";
import {sendEdgeClicked} from "./webSocketInterface/WebSocketInstanceEdgeClicked.tsx";
import ToolsMenu from "./nodes/Tool/ToolsMenu.tsx";
import {ToolProvider, useTool} from "./nodes/Tool/ToolContext.tsx";
import useCustomCursor from "./nodes/Tool/CustomCursor.tsx";
import type { Node } from 'reactflow';

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

    /*Color connections */

    const [selectedEdges, setSelectedEdges] = useState<Edge[]>([]);
    const [color, setColor] = useState('#30362F');
    const connectionLineStyle = { stroke: color, strokeWidth: 1 };
    const onConnect: OnConnect = useCallback(
        (params) =>
            setEdges((eds) =>
                addEdge(
                    {
                        ...params,
                        type: 'step',
                        //className: 'colored-edge',
                        style: { stroke: color, strokeWidth: 1},
                    },
                    eds
                )
            ),
        [color]
    );

    const onSelectionChange = useCallback(({ edges }) => {
        setSelectedEdges(edges ?? []);
    }, []);

    const handleColorChange = useCallback(
        (e: React.ChangeEvent<HTMLInputElement>) => {
            const newColor = e.target.value;
            setColor(newColor);

            // Met à jour les couleurs des edges sélectionnés
            setEdges((eds) =>
                eds.map((edge) =>
                    selectedEdges.some((sel) => sel.id === edge.id)
                        ? {
                            ...edge,
                            //className: 'colored-edge',
                            style: { ...edge.style, stroke: newColor },
                        }
                        : edge
                )
            );
        },
        [selectedEdges, setEdges]
    );


    /* interact */
    const onDragOver = useCallback((event: React.DragEvent<HTMLDivElement>) => {
        event.preventDefault();
        event.dataTransfer.dropEffect = 'move';
        //console.log("drag over"); //From Accordion, not already place
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
        onBuildSaveToDebug()
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
                // Adds onChange to "commentNode" nodes
                const restoredNodes = data.nodes.map((node: Node) => {
                    return node;
                });
                console.log("Restore graph : ",restoredNodes)
                setNodes(restoredNodes);
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
    const onDebug = () => {
        const restorePromise = fetch("http://localhost:8889/debug")
            .then((response) => response.json())
            .then((data) => {
                if (!data) {
                    throw new Error("No data to debug");
                }

                // Ajouter un fallback si la propriété label n'existe pas
                const edgesWithLabels = data.edges.map((edge: any) => ({
                    ...edge,
                    label: edge.label ?? '???', // ou undefined
                    labelBgPadding: [8, 4],
                    labelBgBorderRadius: 4,
                    labelBgStyle: { fill: 'white', color: '#333', fillOpacity: 0.8 },
                    style: {
                        ...edge.style,
                        strokeWidth: 1
                    }
                }));

                setNodes(data.nodes);
                setEdges(edgesWithLabels);

                // Find the highest id among the nodes
                const highestId = Math.max(...data.nodes.map((node: Node) => Number(node.id)));
                // If highestId is a number (not NaN), set id to highestId + 1
                if (!isNaN(highestId)) {
                    id = highestId + 1;
                }
            })
            .catch((error) => {
                if (error.message === "Failed to fetch") {
                    throw new Error("Server isn't running.");
                }
                throw error;
            });

        toast.promise(
            restorePromise,
            {
                pending: "Waiting debug graph...",
                success: "Debug started successfully!",
                error: {
                    render: ({data}) => data.message,
                    autoClose: 7000
                }
            }
        );
    };


    const onDebugStop = () => {
        const restorePromise = fetch("http://localhost:8889/debugStop")
            .then((response) => response.json())
            .then((data) => {
                if (data === null) {
                    throw new Error("No data to restore after debug");
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
                pending: "Waiting stop debug...",
                success: "debug stop successfully!",
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
    const onBuildSaveToDebug = () => {
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

        const savePromise = fetch("http://localhost:8889/json-save-toDebug", requestOptions)
            .then((response) => response.text())
            .then((result) => {
                if (result != "Graph saved to debug") {
                    throw new Error("Failed to save graph to debug");
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
                pending: "preparing debug...",
                success: "Graph saved to debug successfully!",
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

    /* view websocket */
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
    const [isDragging, setIsDragging] = useState(false);
    const handleNodeDragStart = () => {
        //console.log("start move");
        setIsDragging(true);
    };

    const handleNodeDrag = () => {
        //console.log("moving");
    };

    const handleNodeDragStop = () => {
        //console.log("stop move");
        setIsDragging(false);
    };

    useKeyboardShortcuts({ nodes, edges, setNodes, setEdges, getId, isDragging });

    /* Debug */
    const [checkedDebug, setChecked] = React.useState(false);

    const handleToggleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setChecked(event.target.checked);
        console.log("checked :",event.target.checked)
        if (event.target.checked == true){
            onDebug()
        }else{
            onDebugStop()
        }
    };
    useEffect(() => {
        // add or remove the class "debug-active" on the body
        if (checkedDebug) {
            document.body.classList.add('debug-active');
        } else {
            document.body.classList.remove('debug-active');
        }
    }, [checkedDebug]);

    /* tools */
    const { tool } = useTool();
    const { emoji } = useTool(); //cursor appearance

    const handleEdgeClick = (_event: React.MouseEvent, edge: Edge) => {
        console.log(`🖱️ Edge clicked: ${edge.source}, sourceHandle: ${edge.sourceHandle},  Tool: ${tool}`);
        console.log(edge);
        sendEdgeClicked({
            source: edge.source,
            sourceHandle: edge.sourceHandle,
            tool: tool,
        });
    };
    useCustomCursor(emoji, tool);
    const emojiCursor =
        tool === 'default'
            ? 'auto'
            : `url("data:image/svg+xml;utf8,<svg xmlns='http://www.w3.org/2000/svg' height='32' width='32'><text y='24' font-size='24'>${emoji}</text></svg>") 16 16, auto`;




    const onPaneClick = useCallback(
        (event: React.MouseEvent) => {
            const bounds = reactFlowWrapper.current?.getBoundingClientRect();
            /* add comment */
            if (tool === 'comment') {
                if (!bounds || !reactFlowInstance) return;
                const position = reactFlowInstance.screenToFlowPosition({
                    x: event.clientX - bounds.left,
                    y: event.clientY - bounds.top,
                });

                const nodeId = getId();
                const newNode = {
                    id: nodeId,
                    type: 'commentNode',
                    position,
                    data: {
                        text: '',
                    },
                };


                setNodes((nds) => [...nds, newNode]);
            }
        },
        [tool, reactFlowInstance, setNodes]
    );

    /* file manager */
    //Save As
    const handleSaveAs = async () => {
        try {
            if (!reactFlowInstance) {
                return;
            }
            const elements = reactFlowInstance.toObject();
            const json = JSON.stringify(elements, null, 2); // ← jolie mise en forme

            const options = {
                types: [
                    {
                        description: "JSON Files",
                        accept: { "application/json": [".json"] },
                    },
                ],
                suggestedName: "reactflow-diagram.json",
            };

            // Ouvre la boîte de dialogue pour enregistrer
            const handle = await (window as any).showSaveFilePicker(options);

            const writable = await handle.createWritable();
            await writable.write(json);
            await writable.close();

            toast.success("Graph saved successfully in file!");
        } catch (err: any) {
            if (err.name !== "AbortError") {
                toast.error("Error saving file: " + err.message);
            }
            // sinon, utilisateur a annulé, ne rien faire
        }
    };


    //Open File
    const handleOpen = (event: React.ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0];
        if (!file) return;

        const reader = new FileReader();

        reader.onload = (e) => {
            try {
                const text = e.target?.result as string;
                const data = JSON.parse(text);

                if (!data || !data.nodes || !data.edges) {
                    throw new Error("Invalid file format");
                }

                const restoredNodes = data.nodes.map((node: Node) => {
                    return node;
                });

                setNodes(restoredNodes);
                setEdges(data.edges);

                // Find the highest id among the nodes
                const highestId = Math.max(...data.nodes.map((node: Node) => Number(node.id)));
                // If highestId is a number (not NaN), set id to highestId + 1
                if (!isNaN(highestId)) {
                    id = highestId + 1;
                }

                toast.success("Graph loaded from file!");
            } catch (error: any) {
                toast.error(`Error loading file: ${error.message}`);
            }
        };

        reader.readAsText(file);
    };

    return (
        <ReactFlowProvider>
            <div className="dndflow">
                <div className="reactflow-wrapper" ref={reactFlowWrapper} style={{height: '100vh', width: '100%',cursor: emojiCursor,}}>

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

                        onNodeDragStart={handleNodeDragStart}
                        onNodeDrag={handleNodeDrag}
                        onNodeDragStop={handleNodeDragStop}

                        connectionLineComponent={ConnectionLine}
                        connectionLineStyle={connectionLineStyle}
                        onSelectionChange={onSelectionChange}

                        onEdgeClick={handleEdgeClick}
                        onPaneClick={onPaneClick}
                        fitView
                    >
                        <Panel position="top-right" className="menu-panel">
                            {/* Color Picker */}
                            <div className="color-picker">
                                <label>🎨 Color of the connections</label>
                                <input type="color" value={color} onChange={handleColorChange}/>
                            </div>
                            {/* tools */}
                            <ToolsMenu/>
                            {/*toggle debug*/}
                            <div className="switch-container">
                                <span className="switch-label">debug</span>
                                <label className="switch">
                                    <input type="checkbox" checked={checkedDebug} onChange={handleToggleChange}/>
                                    <span className="slider round"></span>
                                </label>
                            </div>
                            {/*button*/}
                            <button className={"button button1"} onClick={openView}>Open view</button>
                            <button className={"button button1 hide-when-debug"} onClick={onSave}>Save</button>
                            <button className={"button button1 hide-when-debug"} onClick={onRestore}>Restore</button>
                            <button className={"button button1 hide-when-debug"} onClick={onBuild}>Build</button>
                            <button className={"button button1 hide-when-debug"} onClick={handleSaveAs}>Save As</button>
                            {/*open File*/}
                            {!checkedDebug && (
                                <>
                                    <input
                                        type="file"
                                        accept=".json"
                                        onChange={handleOpen}
                                        style={{ display: "none" }}
                                        id="fileUpload"
                                    />

                                    <label
                                        htmlFor="fileUpload"
                                        className="button button1"
                                        style={{
                                            display: "inline-flex",
                                            alignItems: "center",
                                            justifyContent: "center",
                                            boxSizing: "border-box",
                                        }}
                                    >
                                        Open File
                                    </label>
                                </>
                            )}

                        </Panel>
                        {/*<MiniMap />*/}
                        <Controls/>
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
            {/*Debug webScoket*/}
            {checkedDebug && <Debug setEdges={setEdges} />}
        </ReactFlowProvider>
    );
}

