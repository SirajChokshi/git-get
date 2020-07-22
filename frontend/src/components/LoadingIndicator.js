import React from 'react';
import {FaSpinner} from 'react-icons/fa'

const LoadingIndicator = ({ size }) => (
    <div className={"flex-centered"}>
        <FaSpinner size={size ? size : "2.5em"} className={"spin-icon"} />
    </div>
)

export default LoadingIndicator;