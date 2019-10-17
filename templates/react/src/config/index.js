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
    items: [
      {
        path: '/',
        label: 'Home',
        icon: 'home',
      },
      {
        path: '/account',
        label: 'Account',
        icon: 'account_circle',
      },
    ],
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
