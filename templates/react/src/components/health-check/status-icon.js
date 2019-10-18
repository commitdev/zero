import React from 'react'
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

export default StatusIcon