import { useMemo } from 'react';

export const useHandleStyles = (handleVisibility: boolean) => {
    const handleStyle = useMemo(
        () => ({
            height: handleVisibility ? 8 : 0,
            width: handleVisibility ? 8 : 0
        }),
        [handleVisibility]
    );

    const handleStyleSideBottom = useMemo(
        () => ({
            height: handleVisibility ? 8 : 0,
            width: handleVisibility ? 8 : 0,
            top: "70%"
        }),
        [handleVisibility]
    );

    const handleStyleSideTop = useMemo(
        () => ({
            height: handleVisibility ? 8 : 0,
            width: handleVisibility ? 8 : 0,
            top: "30%"
        }),
        [handleVisibility]
    );

    return { handleStyle, handleStyleSideBottom, handleStyleSideTop };
};