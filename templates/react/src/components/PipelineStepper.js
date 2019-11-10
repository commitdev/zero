import React from 'react'
import PropTypes from 'prop-types'
import { withStyles } from '@material-ui/core'

import Step from './Step'

const styles = theme => ({
  pipeline: {
    paddingLeft: theme.spacing(6),
    paddingRight: theme.spacing(6),
    display: 'flex',
    flexDirection: 'row',
    flexWrap: 'nowrap',
  },
})

export function PipelineStepper({
  classes,
  steps = [],
  activeStepIndex,
  stepProgress,
}) {
  return (
    <div className={classes.pipeline}>
      {steps.map((step, index) => (
        <Step
          key={index}
          name={step.label}
          progress={stepProgress}
          stepState={getStepState(index, activeStepIndex)}
        />
      ))}
    </div>
  )
}

PipelineStepper.propTypes = {
  steps: PropTypes.array.isRequired,
  activeStepIndex: PropTypes.number.isRequired,
  stepProgress: PropTypes.number,
}

function getStepState(index, activeStepIndex) {
  if (index === activeStepIndex) {
    return 'active'
  } else if (index < activeStepIndex) {
    return 'completed'
  } else {
    return 'inactive'
  }
}

export default withStyles(styles)(PipelineStepper)
