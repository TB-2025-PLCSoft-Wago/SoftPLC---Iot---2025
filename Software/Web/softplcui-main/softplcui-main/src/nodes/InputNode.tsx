import React, {useEffect, useState} from "react";
import {Handle, NodeProps, Position} from "reactflow";


interface InputNodeData {
    label: string;
    services: { friendlyName: string; nameServices: string[] }[];
    outputHandle: { dataType: string; name: string }[];
    subServices: { friendlyName: string; primary: string; secondary: { dataType: string; name: string }[] }[];
    type: string;
    id: string;
    selectedFriendlyNameData?: string;
    selectedServiceData?: string;
    selectedSubServiceData?: string;
    valueData?: string;
    dataType?: string;
    parameterValueData?: string[];
    parameterNameData?: string[];
}

const InputNode: React.FC<NodeProps<InputNodeData>> = (props) => {
    const {
        data = {
            dataType: "BUG",
            id: "BUG",
            label: "BUG",
            type: "BUG",
            outputHandle: [],
            subServices: [],
            services: [],
            handleType: "BUG",
            selectedServiceData: "BUG",
            selectedSubServiceData: "BUG",
            valueData: "BUG",
            selectedFriendlyNameData: "BUG",
            parameterValueData: undefined,
            parameterNameData: undefined,
        }
    } = props;
    const [selectedFriendlyName, setSelectedFriendlyName] = useState(data.selectedFriendlyNameData);
    const [selectedService, setSelectedService] = useState(data.selectedServiceData);
    const [selectedSubService, setSelectedSubService] = useState(data.selectedSubServiceData);
    const [inputValue, setInputValue] = useState(data.valueData);
    const [handleType, setHandleType] = useState("");
    const [inputValues, setInputValues] = useState<string[]>(data.parameterValueData ?? []);


    const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setInputValue(event.target.value);
    }

    const handleInputChange2 = (index: number) => (event: React.ChangeEvent<HTMLInputElement>) => {
        const newValues = [...inputValues];
        newValues[index] = event.target.value;
        setInputValues(newValues);
        data.parameterValueData = newValues;
        console.log("newValues : ",newValues)
    };

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
        data.valueData = inputValue;
        const newHandleType = data.subServices.find(sub => sub.friendlyName === selectedFriendlyName && sub.primary === selectedService)?.secondary.find(sec => sec.name === selectedSubService)?.dataType || "";
        setHandleType(newHandleType);
        data.dataType = newHandleType;
    }, [selectedService, selectedSubService, selectedFriendlyName, inputValue]);


    let content;
    if (data.type === "constantInput") {
        content = (
            <>
                <Handle
                    type={"source"}
                    position={Position.Right}
                    id={data.outputHandle[0].name}
                    isConnectable={true}
                    datatype={data.outputHandle[0].dataType}
                    style={{height: 8, width: 8}}
                />
                <input
                    type="number"
                    className="inputNodeSelect"
                    value={inputValue}
                    onChange={handleInputChange}
                    id={data.id}
                />
            </>
        );
        //console.log("constantInput : ", content);

    } else if (data.type === "viewWebInputBool" || data.type === "viewWebInputValue") {
        //console.log("dataType : ", data.outputHandle[0].dataType);
        content = (
            <>

                <Handle
                    type={"source"}
                    position={Position.Right}
                    id={data.outputHandle[0].name}
                    isConnectable={true}
                    datatype="bool"
                    style={{ height: 8, width: 8 }}
                />

                {[0, 1].map((index) => (
                    <div key={index} style={{ display: "flex", alignItems: "center", marginBottom: "4px" }}>
                            <input
                                id={`${data.id}-${index}`}
                                type="text"
                                className="inputNodeSelect"
                                value={inputValues[index]}
                                onChange={handleInputChange2(index)}
                                placeholder={data.parameterNameData?.[index] ?? ""}
                            />
                    </div>
                ))}

            </>
        );
        //console.log("constantInput : ", content);

    } else if (data.subServices.length > 0) {
        content = (
            <>
                <Handle
                    type={"source"}
                    position={Position.Right}
                    id={data.outputHandle[0].name}
                    isConnectable={true}
                    datatype={handleType}
                    style={{height: 8, width: 8}}
                />
                <select className="inputNodeSelect" value={selectedFriendlyName} onChange={handleFriendlyNameChange}
                        id={data.id}>
                    <option value="default" hidden> Select an appliance</option>
                    {data.services.map((service, index) => (
                        <option key={index} value={service.friendlyName}>
                            {service.friendlyName}
                        </option>
                    ))}
                </select>
                <select className="inputNodeSelect" value={selectedService} onChange={handleServiceChange}
                        id={"serv" + data.id}>
                    <option value="default" hidden> Select a service</option>
                    {data.services.find(service => service.friendlyName === selectedFriendlyName)?.nameServices.map((serviceName, index) => (
                        <option key={index} value={serviceName}>
                            {serviceName}
                        </option>
                    ))}
                </select>
                <select className="inputNodeSelect" value={selectedSubService} onChange={handleSubServiceChange}
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
                <Handle
                    type={"source"}
                    position={Position.Right}
                    id={data.outputHandle[0].name}
                    isConnectable={true}
                    datatype={data.outputHandle[0].dataType}
                    style={{height: 8, width: 8}}
                />
                <select className="inputNodeSelect" value={selectedService} onChange={handleServiceChange} id={data.id}>
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
        <div className="react-flow__node-default inputNode">
            {data.label && <div>{data.label}</div>}
            {content}
        </div>
    );
}

export default InputNode;