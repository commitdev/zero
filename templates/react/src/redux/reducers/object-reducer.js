const initialState = {
  loading: true,
  error: null,
  data: null,
};

export default (loading, loaded, error) => {
  return (state = initialState, action) => {
    switch (action.type) {
      case loading:
        return {
          ...state,
          loading: true,
        }
      case error:
        return {
          ...state,
          loading: false,
          error: action.error,
        }
      case loaded:
        return {
          loading: false,
          error: null,
          data: action.data
        }
      default:
        return state;
    }
  }
}
