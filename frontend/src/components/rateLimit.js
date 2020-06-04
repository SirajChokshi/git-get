import React, {useState, useEffect} from "react";
import './rateLimit.css'

const RateLimit = (props) => {

    const [rateLimit, setRateLimit] = useState(-1)

    const fetchRateLimit = () => {
        fetch(`http://localhost:8080/ratelimit/`,
            {method: "GET", headers: {'Content-Type': 'application/json'}}
        ).then(
            (limit => limit.json())
        ).then (
            json => {
                setRateLimit(json.requests)
            }
        ).catch((e) => {
            console.error(e);
        })
    }

    useEffect(() => fetchRateLimit(), [rateLimit])

    

    return (
        <div id="rate-limit">
            <span>{rateLimit}</span>
            <span>Remaining</span>
            Requests
        </div>
    );
}

export default RateLimit;