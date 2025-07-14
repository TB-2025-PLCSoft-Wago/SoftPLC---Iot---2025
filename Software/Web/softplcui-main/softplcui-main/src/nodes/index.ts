import type { Node, NodeTypes } from "reactflow";
import InputNode from "./InputNode.tsx";
import LogicalNode from "./LogicalNode.tsx";
import OutputNode from "./outputNode.tsx";
import CommentNode from "./CommentNode.tsx";

/*const nodeDataWithId = {
  ...NodeJson.nodes[1],
  id: "12", // Ajoutez l'id ici
};*/

export const initialNodes = [
  /*{id:"f", type:"output", position:{x:400, y:100}, data: { label: "Output",
      externalConnection:"" ,value:""
    }},*/
  /*{
    id: "a",
    type: "inputNode",
    position: { x: 200, y: 200 },
    data:  nodeDataWithId,
  },*/
  ] satisfies Node[];

export const nodeTypes = {
  "outputNode": OutputNode,
  "LogicalNode": LogicalNode,
  "inputNode": InputNode,
  "commentNode": CommentNode,
  // Add any of your custom nodes here!
} satisfies NodeTypes;
