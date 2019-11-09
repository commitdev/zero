import React from 'react'
import Button from '@material-ui/core/Button'
import { withStyles } from '@material-ui/core/styles'
import Divider from '@material-ui/icons/KeyboardArrowRight'
import { router } from '../utils'
import PropTypes from 'prop-types'

class Breadcrumbs extends React.Component {
  render() {
    const { location = '', classes } = this.props
    const segments = location
      .replace(/^\//, '')
      .split('/')
      .filter(a => a)

    const paths = [segments[0]]
    segments.reduce((prev, curr, i) => {
      paths[i] = `${prev}/${curr}`
      return paths[i]
    })

    return (
      <div className={classes.breadcrumbs}>
        {paths.map((p, i) => (
          <React.Fragment key={i}>
            {i !== 0 && <Divider className={classes.divider} />}
            <Button onClick={router.to(`/${p}`)}>{segments[i]}</Button>
          </React.Fragment>
        ))}
      </div>
    )
  }
}

const styles = theme => ({
  breadcrumbs: {
    width: '100%',
    marginBottom: theme.spacing(2),
  },
  divider: {
    verticalAlign: 'middle',
  },
})

Breadcrumbs.propTypes = {
  location: PropTypes.string,
  classes: PropTypes.object,
}

export default withStyles(styles)(Breadcrumbs)
