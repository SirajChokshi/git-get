import React from 'react';
import {FaSpinner} from 'react-icons/fa'

interface LoadingIndicatorProps {
    size: string;
}

const LoadingIndicator = ({ size }: LoadingIndicatorProps) => (
    <div className={"flex-centered"}>
        <FaSpinner size={size} className={"spin-icon"} />
    </div>
)

export default LoadingIndicator;