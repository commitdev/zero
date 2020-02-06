import React from 'react';

export default class Complete extends React.Component {
    render() {
        if (this.props.success) {
            return (
                <div><h3>Success</h3>
                    <p>You're project has been generated.</p>
                </div>
            )
        } else {
            return (
            <div><h3>Failed</h3>
                <p>You're project has not been generated.</p>
            </div>)
        }
    }
}
