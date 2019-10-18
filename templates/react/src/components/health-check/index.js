import React, { useState, useEffect } from 'react'
import axios from 'axios'
import { withStyles } from '@material-ui/core/styles'
import { red, green } from '@material-ui/core/colors'
import Radio from '@material-ui/core/Radio'

const radioStyle = (customColor) => ({
  root: {
    color: customColor,
    '&$checked': {
      color: customColor,
    },
  },
  checked: {},
})

const CustomRadio = (customColor) => (
  withStyles(radioStyle(customColor))(props => <Radio checked={true} color="default" {...props} />)
)

const GreenRadio = CustomRadio(green[600])
const RedRadio = CustomRadio(red[600])

function StatusIcon(props) {
  return props.value === 200 ? <GreenRadio /> : <RedRadio />
}

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