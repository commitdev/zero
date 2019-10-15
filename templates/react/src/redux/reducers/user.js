import objectReducer from 'redux/reducers/object-reducer';

import {
  USER_LOADING,
  USER_LOADED,
  USER_ERROR,
} from 'redux/actions';

export default objectReducer(USER_LOADING, USER_LOADED, USER_ERROR);
