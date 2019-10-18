import React, { useState, useEffect } from 'react'
import axios from 'axios'
import StatusIcon from './status-icon'

function HealthCheck() {
  const [data, setData] = useState({ status: {} })

  useEffect(() => {
    const fetchData = async () => {
      const result = await axios(
        'http://127.0.0.1:8080/v1/health',
      )

      setData(result)
    }

    fetchData()
  }, [])

  return (
    <div>
      API Health Status: <StatusIcon value={data.status}/>
    </div>
  )
}

export default HealthCheck