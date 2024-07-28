import React, { useState } from "react";
import axios from "axios";

function Set() {
  const [key, setKey] = useState("");
  const [value, setValue] = useState("");

  const updateCache = async () => {
    if (!key || !value) {
      window.alert("Key and value cannot be empty");
      return;
    }
    try {
      const response = await axios.post(
        "http://localhost:8080/put",
        {
          key: key,
          value: value,
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );

      window.alert(response.data.message);
      setKey("");
      setValue("");
    } catch (error) {
      window.alert("Error setting key");
      console.error("Error setting key:", error);
    }
  };

  return (
    <div className="set">
      <button className="button" onClick={updateCache}>
        Set
      </button>
      <input
        className="dataKey"
        type="text"
        placeholder="Enter key"
        value={key}
        onChange={(e) => setKey(e.target.value)}
      />
      <input
        className="dataValue"
        type="text"
        placeholder="Enter value"
        value={value}
        onChange={(e) => setValue(e.target.value)}
      />
    </div>
  );
}

export default Set;
