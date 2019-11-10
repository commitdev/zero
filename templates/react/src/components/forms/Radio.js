import React from 'react'
import FormControlLabel from '@material-ui/core/FormControlLabel'
import PropTypes from 'prop-types'

import { makeStyles } from '@material-ui/core/styles'
import Radio from '@material-ui/core/Radio'
import RadioGroup from '@material-ui/core/RadioGroup'
import FormControl from '@material-ui/core/FormControl'
import FormLabel from '@material-ui/core/FormLabel'

const useStyles = makeStyles(theme => ({
  root: {
    display: 'flex',
  },
  formControl: {},
  group: {
    margin: theme.spacing(1, 0),
  },
}))

export default function RadioButtons({
  name,
  defaultValue,
  label,
  required,
  options,
}) {
  const classes = useStyles()

  return (
    <div className={classes.root}>
      <FormControl component="fieldset" className={classes.formControl}>
        <FormLabel component="legend">{label || name}</FormLabel>
        <RadioGroup
          aria-label={label}
          name={name}
          className={classes.group}
          defaultValue={defaultValue}
        >
          {options.map(({ label, value }) => (
            <FormControlLabel
              value={value}
              control={<Radio />}
              label={label || value}
            />
          ))}
        </RadioGroup>
      </FormControl>
    </div>
  )
}

RadioButtons.propTypes = {
  options: PropTypes.arrayOf(PropTypes.object),
  defaultValue: PropTypes.string,
  label: PropTypes.string.isRequired,
}
