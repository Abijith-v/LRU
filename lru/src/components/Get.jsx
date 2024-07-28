import React, { useState } from 'react';
import axios from 'axios';

function Get() {
    const [key, setKey] = useState('');
    const [value, setValue] = useState('');
    const [message, setMessage] = useState('');

    const fetchKey = async () => {
        try {
            const response = await axios.get('http://localhost:8080/get'); // url to be updated
            setMessage('')
            setValue(response.data['value'])
            setKey(response.data['key'])
        } catch (error) {
            console.error("Error fetching value:", error);
            setValue('')
            setKey('')
            setMessage(error.response['data']);
        }
    };

    return (
        <div className= 'getContainer'>
            <button className="button" onClick={fetchKey}>Get</button>
    
            {key && (
                <div className='dataKey'>
                    <h3>key: {key}</h3>
                </div>
            )}
            {value && (
                <div className='dataValue'>
                    <h3>Value: {value}</h3>
                </div>
            )}
            {message && (
                <div className='dataValue'>
                    <h3>Message: {message}</h3>
                </div>
            )}
        </div>
    );
}

export default Get;
