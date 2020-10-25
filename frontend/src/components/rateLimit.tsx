import React, {useState, useEffect} from "react";
import LoadingIndicator from "./LoadingIndicator"
import './rateLimit.scss'

const RateLimit = () => {

    const BASE_URL = "https://arcane-ocean-76968.herokuapp.com/"

    const [rateLimit, setRateLimit] = useState({loading: true})

    const fetchRateLimit = () => {
        fetch(`${BASE_URL}ratelimit/`,
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
        {rateLimit.loading
                    ?
                    <>
                        <span><LoadingIndicator size={"1.3em"} /></span>
                        <span>Loading...</span>
                    </>
                    :
                    (
                        <>
                        <span>{rateLimit}</span>
                        <span>Remaining</span>
                        Requests
                        </>
                    )
                }
        </div>
    );
}

export default RateLimit;