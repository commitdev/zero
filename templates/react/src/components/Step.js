import React from 'react'
import classNames from 'classnames'
import { withStyles } from '@material-ui/core'
import CheckCircleOutlineIcon from '@material-ui/icons/CheckCircleOutline'

import { P } from '../components/Text'

const lineWidth = '1px'
const circleRadius = 5

const styles = theme => ({
  pipelineItem: {
    position: 'relative',
    flex: '1 1 auto',
    paddingBottom: theme.spacing(1),
    textAlign: 'center',
    '& svg': {
      position: 'absolute',
    },
    '& #start': {
      bottom: -circleRadius - 1,
      left: -circleRadius,
    },
    '& #end': {
      bottom: -circleRadius - 1,
      right: -circleRadius,
    },
    borderBottom: `${lineWidth} solid ${theme.palette.grey[300]}`,
    '& #progress': {
      position: 'absolute',
      bottom: -1,
    },
  },
  pipelineItemCompleted: {
    color: theme.palette.grey[600],
    borderBottom: `${lineWidth} solid ${theme.palette.primary.main}`,
  },
  pipelineItemInactive: {
    color: theme.palette.grey[300],
    '& svg': {
      fill: theme.palette.grey[300],
    },
  },
  pipelineItemActive: {
    color: theme.palette.primary.main,
    '& #progress': {
      width: '50%',
      borderBottom: `${lineWidth} solid ${theme.palette.primary.main}`,
    },
    '& #start': {
      fill: theme.palette.primary.main,
    },
    '& #end': {
      fill: theme.palette.grey[300],
    },
    '& p': {
      fontWeight: 'bold',
    },
  },
  checkmark: {
    color: theme.palette.grey[600],
    marginLeft: theme.spacing(1),
    // marginTop: theme.spacing(0.25),
    fontSize: theme.typography.h6.fontSize,
  },
})

class Step extends React.Component {
  render() {
    const { name, classes, progress, stepState } = this.props
    const statusStyles = {
      completed: classes.pipelineItemCompleted,
      inactive: classes.pipelineItemInactive,
      active: classes.pipelineItemActive,
    }

    return (
      <div
        className={classNames(classes.pipelineItem, statusStyles[stepState])}
      >
        <P inline color="inherit">
          {name}
        </P>
        {stepState === 'active' ? (
          <div id="progress" style={{ width: `${progress}%` }} />
        ) : null}
        {stepState === 'completed' ? (
          <CheckCircleOutlineIcon className={classes.checkmark} />
        ) : null}
        <svg id="start" height="10" width="10">
          <circle cx={circleRadius} cy={circleRadius} r={circleRadius} />
        </svg>
        <svg id="end" height="10" width="10">
          <circle cx={circleRadius} cy={circleRadius} r={circleRadius} />
        </svg>
      </div>
    )
  }
}

export default withStyles(styles)(Step)
