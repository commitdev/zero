import ownerDocument from './ownerDocument';

function ownerWindow(node) {
  var doc = ownerDocument(node);
  return doc.defaultView || window;
}

export default ownerWindow;