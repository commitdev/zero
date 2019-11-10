import React from 'react'
import cn from './SuccessIndicator.module.css'

const SuccessIndicator = () => (
  <div className={cn.sn}>
    <div className={cn['sa-success']}>
      <div className={cn['sa-success-tip']} />
      <div className={cn['sa-success-long']} />
      <div className={cn['sa-success-placeholder']} />
      <div className={cn['sa-success-fix']} />
    </div>
  </div>
)

export default SuccessIndicator
