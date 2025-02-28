export interface Node {
    accordion: string;
    primaryType: string;
    type: string;
    display: string;
    label: string;
    stretchable: boolean;
    services: {friendlyName:string; nameServices:string[]}[];
    subServices: {friendlyName:string; primary:string; secondary:{dataType:string; name:string}[]}[];
    inputHandle: { dataType: string; name: string }[];
    outputHandle: { dataType: string; name: string }[];
    selectedServiceData?: string;
    selectedSubServiceData?: string;
    valueData?: string;
}