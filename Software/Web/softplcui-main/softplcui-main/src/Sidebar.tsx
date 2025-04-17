import React, { useEffect, useState } from 'react';
import { Accordion, AccordionItem, AccordionItemHeading, AccordionItemButton, AccordionItemPanel } from 'react-accessible-accordion';
import 'react-accessible-accordion/dist/fancy-example.css';
import { Node } from './nodes/Node.ts';
import  {NodesData}  from './App.tsx';



const Sidebar: React.FC<{ nodesData: NodesData }> = ({ nodesData }) => {
    const [nodes, setNodes] = useState<Record<string, Node[]>>({});

    useEffect(() => {
        const nodesByAccordion = nodesData.nodes.reduce((acc: Record<string, Node[]>, node: Node) => {
            if (!acc[node.accordion]) {
                acc[node.accordion] = [];
            }
            acc[node.accordion].push(node);
            return acc;
        }, {});
        setNodes(nodesByAccordion);
    }, [nodesData]);

    const onDragStart = (event: React.DragEvent<HTMLDivElement>, nodeType: string, nodeLabel: string, subtype: string) => {
        event.dataTransfer.setData('application/reactflow', nodeType);
        event.dataTransfer.setData('nodeLabel', nodeLabel);
        event.dataTransfer.setData('subtype', subtype);
        event.dataTransfer.effectAllowed = 'move';
    };

    return (
        <aside>
            <div className="description">Drag the bloc you want to add.</div>
            <div className="Sidebar">
                <Accordion allowMultipleExpanded allowZeroExpanded>
                    {Object.entries(nodes).map(([accordion, nodesOfAccordion], index) => (
                        <AccordionItem uuid={accordion.replace(/\s/g, '-')} key={index}>
                            <AccordionItemHeading>
                                <AccordionItemButton>{accordion}</AccordionItemButton>
                            </AccordionItemHeading>
                            <AccordionItemPanel>
                                {nodesOfAccordion.map((node, index) => (
                                    <div className={`dndnode ${node.type}`}
                                         onDragStart={(event) =>
                                             onDragStart(event, node.type, node.display, node.primaryType)}
                                         draggable={true}
                                         key={index}>
                                        {node.display}
                                    </div>
                                ))}
                            </AccordionItemPanel>
                        </AccordionItem>
                    ))}
                </Accordion>
            </div>
        </aside>
    );

};

export default Sidebar;