import React from 'react'
import TextField from '@material-ui/core/TextField'
import InputLabel from '@material-ui/core/InputLabel'

export default function TextArea({
  name,
  label,
  errorText,
  helpText,
  ...fieldProps
}) {
  return (
    <div>
      <InputLabel>{label || name}</InputLabel>
      <TextField
        id={name}
        name={name}
        helperText={errorText || helpText}
        multiline
        rows={4}
        rowsMax={12}
        {...fieldProps}
      />
    </div>
  )
}
