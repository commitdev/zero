import React from 'react'
import Grid from '@material-ui/core/Grid'
import PropTypes from 'prop-types'

const Centered = props => {
  const { children, container, horizontal, ...passThrough } = props
  const style = 'container' in props ? { flex: '1 0 auto' } : {}
  const justify = 'horizontal' in props ? null : 'center'
  return (
    <Grid
      container
      className="animated fadeIn"
      justify={justify}
      alignItems="center"
      direction="column"
      style={style}
      {...passThrough}
    >
      {children}
    </Grid>
  )
}

Centered.propTypes = {
  container: PropTypes.bool,
  horizontal: PropTypes.bool,
}

export default Centered
