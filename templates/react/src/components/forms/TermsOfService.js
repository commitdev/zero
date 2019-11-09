import React, { useState, useEffect } from 'react'
import PropTypes from 'prop-types'
import { withStyles } from '@material-ui/core/styles'
import {
  Checkbox,
  Button,
  Box,
  FormControlLabel,
  Paper,
  Typography,
} from '@material-ui/core'

import { P, Text } from '../Text'

function TermsOfService({ classes, agreed, label, name, defaultValue }) {
  const [termsText, setText] = useState('')
  const [checked, setChecked] = useState(Boolean(defaultValue))

  useEffect(() => {
    import('../../assets/TermsOfService.json').then(terms => {
      setText(terms['termsText'])
    })
  }, [])

  return (
    <React.Fragment>
      <P gutterBottom>phrase.tos_read</P>
      <Paper style={{ width: '100%' }}>
        <Box p={2} className={classes.textContainer}>
          <Typography color="textSecondary">{termsText}</Typography>
        </Box>
        <Box>
          <Button onClick={handleDownload(termsText)}>
            <Text>download</Text>
          </Button>
        </Box>
      </Paper>
      <FormControlLabel
        label={label || name}
        control={
          <Checkbox
            checked={checked}
            name={name}
            value="agreed"
            color="primary"
            onChange={e => setChecked(e.target.checked)}
          />
        }
      />
    </React.Fragment>
  )
}

function handleDownload(termsText) {
  return () => {
    if (termsText !== '') {
      var element = document.createElement('a')
      var file = new Blob([new TextEncoder().encode(termsText)], {
        type: 'text/plain',
      })
      element.href = URL.createObjectURL(file)
      element.download = 'TermsOfService.txt'
      element.click()
    }
  }
}

const styles = theme => ({
  textContainer: {
    borderBottom: `1px solid ${theme.palette.grey[300]}`,
    color: theme.palette.grey[300],
    maxHeight: 250,
    overflowY: 'auto',
    whiteSpace: 'pre-line',
  },
})

TermsOfService.propTypes = {
  classes: PropTypes.object.isRequired,
  name: PropTypes.string.isRequired,
  label: PropTypes.string,
  defaultValue: PropTypes.oneOfType([PropTypes.string, PropTypes.number]),
}

export default withStyles(styles)(TermsOfService)
