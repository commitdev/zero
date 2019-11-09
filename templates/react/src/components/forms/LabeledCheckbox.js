import React, { useState } from 'react'
import FormControlLabel from '@material-ui/core/FormControlLabel'
import Checkbox from '@material-ui/core/Checkbox'
import PropTypes from 'prop-types'

export default function LabeledCheckbox({ name, value, label, required }) {
  const [checked, setChecked] = useState(false)

  return (
    <FormControlLabel
      control={
        <Checkbox
          id={name}
          checked={checked}
          color="default"
          inputProps={{ required }}
          onChange={() => setChecked(!checked)}
          value={value}
        />
      }
      label={label}
    />
  )
}

LabeledCheckbox.propTypes = {
  checked: PropTypes.bool.isRequired,
  value: PropTypes.string.isRequired,
  label: PropTypes.string.isRequired,
}
