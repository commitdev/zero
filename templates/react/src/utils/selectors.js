import React from 'react'
import hoistNonReactStatic from 'hoist-non-react-statics'

// ref: https://reactjs.org/docs/higher-order-components.html
function withSelector(selector) {
  return WrappedComponent => {
    class Enhance extends React.PureComponent {
      render() {
        const { data, ...passThrough } = this.props
        const processedData = selector(this.props)
        return (
          <WrappedComponent
            {...passThrough}
            mutate={data.mutate}
            refetch={data.refetch}
            data={processedData}
          />
        )
      }
    }
    Enhance.displayName = `withSelector(${getDisplayName(WrappedComponent)})`
    hoistNonReactStatic(Enhance, WrappedComponent)

    return Enhance
  }
}

function getDisplayName(WrappedComponent) {
  return WrappedComponent.displayName || WrappedComponent.name || 'Component'
}

export default {
  withSelector,
  getDisplayName,
}
