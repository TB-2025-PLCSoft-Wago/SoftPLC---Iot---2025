export interface LogicalNodeData {
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
    parameterNameData?: string[];
    type: string;
}
