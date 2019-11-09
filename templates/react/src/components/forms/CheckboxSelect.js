import React from 'react'
import PropTypes from 'prop-types'

import { makeStyles } from '@material-ui/core/styles'
import FormLabel from '@material-ui/core/FormLabel'
import FormControl from '@material-ui/core/FormControl'
import FormGroup from '@material-ui/core/FormGroup'
import FormControlLabel from '@material-ui/core/FormControlLabel'
import FormHelperText from '@material-ui/core/FormHelperText'
import Checkbox from '@material-ui/core/Checkbox'

const useStyles = makeStyles(theme => ({
  root: {
    display: 'flex',
  },
  formControl: {
    // margin: theme.spacing(3),
  },
}))

export default function CheckboxSelect({
  name,
  defaultValue,
  label,
  required,
  helpText,
  options,
}) {
  const classes = useStyles()
  let values = new Set()

  try {
    if (defaultValue) values = new Set(JSON.parse(defaultValue))
  } catch (e) {
    console.warn('invalid value format', defaultValue)
  }
  const [state, setState] = React.useState({
    values,
    value: defaultValue,
  })

  const handleChange = name => event => {
    event.target.checked ? state.values.add(name) : state.values.delete(name)
    setState({
      values: state.values,
      value: JSON.stringify(Array.from(state.values.values())),
    })
  }

  return (
    <div className={classes.root}>
      <FormControl component="fieldset" className={classes.formControl}>
        <FormLabel component="legend">{label || name}</FormLabel>
        <input type="hidden" name={name} value={state.value || ''} />
        <FormGroup>
          {options.map(({ label, value }, i) => (
            <FormControlLabel
              key={i}
              onChange={handleChange(value)}
              control={<Checkbox checked={state.values.has(value)} />}
              label={label || value}
            />
          ))}
        </FormGroup>
        <FormHelperText>{helpText}</FormHelperText>
      </FormControl>
    </div>
  )
}

CheckboxSelect.propTypes = {
  options: PropTypes.arrayOf(PropTypes.object),
  defaultValue: PropTypes.string,
  label: PropTypes.string.isRequired,
}
