import React, { useState } from 'react';
import axios from 'axios';

function Delete() {

    const deleteKeyValue = async () => {
        try {

            await axios.delete('http://localhost:8080/delete'); 
            window.alert('Key deleted sucessfully')
            console.log('Key deleted successfully')  
        } catch (error) {
            window.alert('Error deleting value')
            console.error("Error deleting value:", error);
        }
    };

    return (
        <div>
        <button className="button" onClick={deleteKeyValue}>Delete</button>
        </div>
    );
}

export default Delete;
