import React from 'react';
/**
 * @ignore - internal component.
 */

var FormControlContext = React.createContext();
export function useFormControl() {
  return React.useContext(FormControlContext);
}
export default FormControlContext;