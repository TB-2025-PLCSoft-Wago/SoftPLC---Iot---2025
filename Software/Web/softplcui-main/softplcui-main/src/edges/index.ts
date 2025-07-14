import type { Edge, EdgeTypes } from "reactflow";

import CustomEdgeStartEndDebug from "./CustomEdgeStartEndDebug.tsx";
import CustomEdgeStepControl from "./CustomEdgeStepControl.tsx";

export const initialEdges = [
] satisfies Edge[];

export const edgeTypes = {
  // Add your custom edge types here!
    customDebugEdge : CustomEdgeStartEndDebug,
    customStep  : CustomEdgeStepControl,
} satisfies EdgeTypes;
