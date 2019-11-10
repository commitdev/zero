import React from 'react'
import Typography from '@material-ui/core/Typography'
import Lang from '../utils/lang'

export default Lang

export const Text = props => (
  <React.Fragment>{Lang.t(props.path || props.children)}</React.Fragment>
)
export const H1 = build('h1')
export const H2 = build('h2')
export const H3 = build('h3')
export const H4 = build('h4')
export const H5 = build('h5')
export const H6 = build('h6')
export const P = build('p', 'body1')
export const Span = build('span')
export const Small = build('small', 'body2')

function build(component, variant) {
  // return props => React.createElement(component, {}, Lang.t(props.children))
  return props => (
    <Typography {...props} component={component} variant={variant || component}>
      {Lang.t(props.children)}
    </Typography>
  )
}
