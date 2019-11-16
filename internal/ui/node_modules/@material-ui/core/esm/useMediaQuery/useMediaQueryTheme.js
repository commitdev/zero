import useMediaQuery from './useMediaQuery'; // TODO to deprecate in v4.x and remove in v5

function useMediaQueryTheme() {
  return useMediaQuery.apply(void 0, arguments);
}

export default useMediaQueryTheme;