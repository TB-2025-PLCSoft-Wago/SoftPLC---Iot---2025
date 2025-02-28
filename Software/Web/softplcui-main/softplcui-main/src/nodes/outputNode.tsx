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
}

const OutputNode: React.FC<NodeProps<OutputNodeData>> = (props) => {
    const {
        data = {
            dataType: "BUG",
            id: "BUG",
            label: "BUG",
            subServices: [],
            services: [],
            inputHandle: [],
            handleType: "BUG",
            selectedFriendlyNameData: "BUG",
            selectedServiceData: "BUG",
            selectedSubServiceData: "BUG",
            valueData: "BUG"
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
    useEffect(() => {

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
        const newHandleType = data.subServices.find(sub => sub.friendlyName === selectedFriendlyName && sub.primary === selectedService)?.secondary.find(sec => sec.name === selectedSubService)?.dataType || "";
        setHandleType(newHandleType);
        data.dataType = newHandleType;
    }, [selectedService, selectedSubService, selectedFriendlyName]);

    let content;
    if (data.subServices.length > 0) {
        content = (
            <>
                <CustomHandle
                    type={"target"}
                    position={Position.Left}
                    id={data.inputHandle[0].name}
                    isConnectable={1}
                    datatype={handleType}
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
    return (
        <div className="react-flow__node-default outputNode">
            {data.label && <div>{data.label}</div>}
            {content}
        </div>
    );
}

export default OutputNode;