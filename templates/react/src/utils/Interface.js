const PropTypes = require('prop-types')

export default class Interface {
  constructor(props, propTypes) {
    // checks the proptype when in development
    if (process.env.NODE_ENV !== 'production') {
      PropTypes.checkPropTypes(propTypes, props, 'prop', this.constructor.name)
    }
  }
}
