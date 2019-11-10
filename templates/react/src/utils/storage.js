/**
 * Localstorage wrapper
 * - automatically serializes JSON
 * - falls back to memory storage
 */

let storage
try {
  let fail
  let uid = new Date().toString()
  ;(storage = window.localStorage).setItem('uid', uid)
  fail = storage.getItem('uid') !== uid
  storage.removeItem('uid')
  fail && (storage = false)
} catch (e) {
  console.error('localstorage not supported')
}

/**
 * Shim localstorage to fallback to in memory storage when storage is disabled
 */
const createMemoryStorage = () => {
  const memoryStorage = {}
  return {
    getItem: function(key) {
      if (key === undefined) return memoryStorage
      else return memoryStorage[key]
    },
    setItem: function(key, value) {
      memoryStorage[key] = value
    },
    removeItem: function(key) {
      delete memoryStorage[key]
    },
  }
}

/**
 * fallback for private browsing mode (safari) where using localstorage will thrown an exception
 */
function shimLocalStorage() {
  return {
    getItem: function(key) {
      let value = localStorage.getItem(key)
      if (value !== undefined) {
        try {
          value = JSON.parse(value)
        } catch (e) {}
      }
      return value
    },
    setItem: function(key, value) {
      if (value !== undefined && value !== null) {
        localStorage.setItem(key, JSON.stringify(value))
      } else {
        localStorage.removeItem(key)
      }
    },
    removeItem: function(key) {
      localStorage.removeItem(key)
    },
  }
}

let localstore = storage ? shimLocalStorage() : createMemoryStorage()

module.exports = {
  getItem: localstore.getItem,
  setItem: localstore.setItem,
  removeItem: localstore.removeItem,
}
