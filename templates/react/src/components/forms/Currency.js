import React from 'react'
import TextField from '@material-ui/core/TextField'
import InputAdornment from '@material-ui/core/InputAdornment'

export default function Currency({
  name,
  label,
  errorText,
  helpText,
  ...fieldProps
}) {
  return (
    <TextField
      id={name}
      name={name}
      label={label || name}
      helperText={errorText || helpText}
      InputProps={{
        inputProps: {
          type: 'number',
          min: 0,
        },
        startAdornment: <InputAdornment position="start">$</InputAdornment>,
      }}
      {...fieldProps}
    />
  )
}
