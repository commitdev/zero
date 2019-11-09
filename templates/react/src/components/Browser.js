import React from 'react'
import PropTypes from 'prop-types'
import cn from './Browser.module.css'
import cx from 'classnames'

const Browser = ({ src = '', width, height, children }) => (
  <React.Fragment>
    <div className={cx(cn.container, 'animated', 'fadeIn')}>
      <div className={cn.browser} style={{ width, minHeight: height }}>
        <img src={src} alt="" width={width} className={cn.source} />
        {children}
      </div>
      <div className={cn.overlay} />
    </div>
  </React.Fragment>
)

Browser.proptTypes = {
  src: PropTypes.string.required,
  width: PropTypes.number,
  height: PropTypes.number,
}

export default Browser
