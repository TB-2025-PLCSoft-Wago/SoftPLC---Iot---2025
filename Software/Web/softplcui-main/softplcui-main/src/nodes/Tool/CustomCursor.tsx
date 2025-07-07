import { useEffect } from 'react';

const useCustomCursor = (emoji: string, tool: string) => {
    useEffect(() => {
        const pane = document.querySelector('.react-flow__pane') as HTMLElement | null;
        if (!pane) return;

        if (tool === 'default') {
            pane.style.cursor = 'auto';
        } else {
            pane.style.cursor = `url("data:image/svg+xml;utf8,<svg xmlns='http://www.w3.org/2000/svg' height='32' width='32'><text y='24' font-size='24'>${emoji}</text></svg>") 16 16, auto`;
        }
    }, [emoji, tool]);
};
export default useCustomCursor;