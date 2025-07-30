import {NodeProps, Position} from "reactflow";
import CustomHandle from "./CustomHandle.tsx";
import React, {useEffect, useState} from "react";


interface OutputNodeData {
    label: string;
    services: { friendlyName: string; nameServices: string[] }[];
    inputHandle: { dataType: string; name: string }[];
    subServices: { friendlyName: string; primary: string; secondary: { dataType: string; name: string }[] }[];
    type: string;
    value: string;
    id: string;
    handleType?: string;
    selectedFriendlyNameData?: string;
    selectedServiceData?: string;
    selectedSubServiceData?: string;
    valueData?: string;
    dataType?: string;
    parameterValueData?: string[];
}

const OutputNode: React.FC<NodeProps<OutputNodeData>> = (props) => {
    const {
        data = {
            dataType: "BUG",
            id: "BUG",
            label: "BUG",
            type: "BUG",
            subServices: [],
            services: [],
            inputHandle: [],
            handleType: "BUG",
            selectedFriendlyNameData: "BUG",
            selectedServiceData: "BUG",
            selectedSubServiceData: "BUG",
            valueData: "BUG",
            parameterValueData: undefined,
            parameterNameData: undefined,
        }
    } = props;
    const [selectedFriendlyName, setSelectedFriendlyName] = useState(data.selectedFriendlyNameData);
    const [selectedService, setSelectedService] = useState(data.selectedServiceData);
    const [selectedSubService, setSelectedSubService] = useState(data.selectedSubServiceData);

    const [handleType, setHandleType] = useState("");
    data.valueData = "";
    const handleFriendlyNameChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        setSelectedFriendlyName(event.target.value);
        setSelectedService("default");
        setSelectedSubService("default");
    }

    const handleServiceChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        setSelectedService(event.target.value);
        setSelectedSubService("default");
    }

    const handleSubServiceChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        setSelectedSubService(event.target.value);
    }
    const handleServiceChangeInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        setSelectedService(event.target.value);
        //setSelectedSubService("default");
        console.log("service input : ",event.target.value)
        data.selectedServiceData = event.target.value
    }

    const handleSubServiceChangeInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        setSelectedSubService(event.target.value);
        console.log("sub service input : ",event.target.value)
        data.selectedSubServiceData = event.target.value
    }
    useEffect(() => {
        console.log("useEffect");
        if (!data.type.includes("viewWeb") && !data.type.includes("variable")) {
            if (selectedFriendlyName === "default") {
                data.selectedFriendlyNameData = "";
            } else {
                data.selectedFriendlyNameData = selectedFriendlyName;
            }
            if (data.services.length > 0) {
                data.selectedServiceData = selectedService;
            } else {
                data.selectedServiceData = "";
            }
            if (data.subServices.length > 0) {
                data.selectedSubServiceData = selectedSubService;
            } else {
                data.selectedSubServiceData = "";
            }
        }else {
            if (data.selectedServiceData === "default") {
                data.selectedServiceData = "";
            }
            if (data.selectedSubServiceData === "default") {
                data.selectedSubServiceData = "";
            }
        }
        const newHandleType = data.subServices.find(sub => sub.friendlyName === selectedFriendlyName && sub.primary === selectedService)?.secondary.find(sec => sec.name === selectedSubService)?.dataType || "";
        setHandleType(newHandleType);
        data.dataType = newHandleType;
    }, [selectedService, selectedSubService, selectedFriendlyName]);

    useEffect(() => {
        setSelectedFriendlyName(data.selectedFriendlyNameData ?? "default");
    }, [data.selectedFriendlyNameData]);

    useEffect(() => {
        setSelectedService(data.selectedServiceData ?? "default");
    }, [data.selectedServiceData]);

    useEffect(() => {
        setSelectedSubService(data.selectedSubServiceData ?? "default");
    }, [data.selectedSubServiceData]);

    let content;
    const {id, selectedServiceData, selectedSubServiceData } = data;
    if (data.type.includes("viewWeb") || data.type.includes("variable")) {
        content = (
            <>
                <CustomHandle
                    type={"target"}
                    position={Position.Left}
                    id={data.inputHandle[0].name}
                    isConnectable={1}
                    datatype={data.inputHandle[0].dataType}
                    className="inputhandleClass"
                    style={{height: 8, width: 8}}
                ></CustomHandle>

                <div style={{ display: "flex", flexDirection: "column", gap: "4px", marginBottom: "8px" }}>
                    {/* Input 1 */}
                    <div style={{ display: "flex", alignItems: "center" }}>
                        <input
                            id={`${id}-0`}
                            type="text"
                            className="inputNodeSelect"
                            value={selectedServiceData}
                            onChange={handleServiceChangeInput}
                            placeholder={data.type.includes("viewWeb") ? "appliance name" : "name"}
                        />
                    </div>

                    {/* Input 2 */}
                    <div style={{ display: "flex", alignItems: "center" }}>
                        <input
                            id={`${id}-1`}
                            type="text"
                            className="inputNodeSelect"
                            value={selectedSubServiceData}
                            onChange={handleSubServiceChangeInput}
                            placeholder={data.type.includes("viewWeb") ? "signal name":"default value" }
                        />
                    </div>
                </div>
            </>
        );
    } else if (data.subServices.length > 0) {
        content = (
            <>
                <CustomHandle
                    type={"target"}
                    position={Position.Left}
                    id={data.inputHandle[0].name}
                    isConnectable={1}
                    datatype={handleType}
                    className="inputhandleClass"
                    style={{height: 8, width: 8}}
                />
                <select className="outputNodeSelect" value={selectedFriendlyName} onChange={handleFriendlyNameChange}
                        id={data.id}>
                    <option value="default" hidden> Select an appliance</option>
                    {data.services.map((service, index) => (
                        <option key={index} value={service.friendlyName}>
                            {service.friendlyName}
                        </option>
                    ))}
                </select>
                <select className="outputNodeSelect" value={selectedService} onChange={handleServiceChange}
                        id={"serv" + data.id}>
                    <option value="default" hidden> Select a service</option>
                    {data.services.find(service => service.friendlyName === selectedFriendlyName)?.nameServices.map((serviceName, index) => (
                        <option key={index} value={serviceName}>
                            {serviceName}
                        </option>
                    ))}
                </select>
                <select className="outputNodeSelect" value={selectedSubService} onChange={handleSubServiceChange}
                        id={"subServ" + data.id}>
                    <option value="default" hidden>Select a sub service</option>
                    {data.subServices.find(sub => sub.friendlyName === selectedFriendlyName && sub.primary === selectedService)?.secondary.map((subService, index) => (
                        <option key={index} value={subService.name}>
                            {subService.name}
                        </option>
                    ))}
                </select>

            </>
        );
    } else {
        content = (
            <>
                <CustomHandle
                    type={"target"}
                    position={Position.Left}
                    id={data.inputHandle[0].name}
                    isConnectable={1}
                    datatype={data.inputHandle[0].dataType}
                    className="inputhandleClass"
                    style={{height: 8, width: 8}}
                ></CustomHandle>
                <select className="outputNodeSelect" value={selectedService} onChange={handleServiceChange}
                        id={data.id}>
                    <option value="default" hidden> Select a service</option>
                    {data.services.map((service) => (
                        service.nameServices.map((serviceName, index) => (
                            <option key={index} value={serviceName}>
                                {serviceName}
                            </option>
                        ))
                    ))}
                </select>
            </>
        );
    }
    let labelColor = "black";
    if (data.type.includes("view")) {
        labelColor = "#2F57A7";
    } else if (data.type.includes("variable")) {
        labelColor = "#8A429E";
    } else if (data.type.includes("constant")) {
        labelColor = "brown";
    }
    return (
        <div className="react-flow__node-default outputNode">
            {data.label && <div style={{ color: labelColor }}>{data.label}</div>}
            {content}
        </div>
    );
}

export default OutputNode;