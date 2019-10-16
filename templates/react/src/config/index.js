export default {
  app: {
    name: 'Commit0',
  },
  account: {
    enabled: true,
    required: true,
  },
  header: {
    enabled: true,
  },
  sidenav: {
    enabled: true,
  },
  views: [
    {
      path: '/account',
      component: 'account',
    },
    {
      path: '/',
      component: 'home',
    },
  ],
}
