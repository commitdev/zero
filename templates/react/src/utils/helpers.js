/**
 * getProp utility - simple implementation of lodash.get
 * @param {Object} object
 * @param {String|Array} path
 * @param {*} defaultVal
 */
function getProp(object, path, defaultVal) {
  const _path = Array.isArray(path)
    ? path
    : path.split('.').filter(i => i.length)

  if (!_path.length) {
    return object === undefined ? defaultVal : object
  } else if (object === undefined || object === null) {
    return undefined
  }

  return getProp(object[_path.shift()], _path, defaultVal)
}

/**
 * setProp utility - simple implementation of lodash.set
 * @param {Object} object
 * @param {String|Array} path
 * @param {*} val
 */
function setProp(object, keys, val) {
  keys = Array.isArray(keys) ? keys : keys.split('.').filter(i => i.length)
  if (keys.length > 1) {
    object[keys[0]] = object[keys[0]] || {}
    return setProp(object[keys[0]], keys.slice(1), val)
  }
  object[keys[0]] = val
}

/**
 * setProp utility - simple implementation of lodash.set
 * @param {Array} collection
 * @param {String} key
 */
function keyBy(collection = [], key) {
  return collection.reduce((obj, val) => {
    const keyValue = val[key]
    obj[keyValue] = val
    return obj
  }, {})
}

function clamp(number, min, max = Infinity) {
  return Math.min(Math.max(number, min), max)
}

module.exports = {
  getProp,
  setProp,
  keyBy,
  clamp,
}
