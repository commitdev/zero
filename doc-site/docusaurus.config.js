/** @type {import('@docusaurus/types').DocusaurusConfig} */

module.exports = {
  title: 'Zero',
  tagline: 'Opinionated infrastructure to take you from idea to production on day one',
  url: process.env.BUILD_DOMAIN ? `https://${process.env.BUILD_DOMAIN}` : 'https://staging.getzero.dev',
  baseUrl: '/',
  onBrokenLinks: 'warn',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',
  organizationName: 'commitdev',
  projectName: 'zero',
  themeConfig: {

    colorMode: {
      defaultMode: 'dark',
    },
    navbar: {
      logo: {
        alt: 'Zero Logo',
        src: 'img/zero.svg',
      },
      items: [
        {
          to: '/docs/zero/about/overview',
          label: 'Docs',
          className: 'header-docs-link header-logo-24',
          position: 'right'
        },
        {
          href: 'https://slack.getzero.dev',
          label: 'Slack',
          className: 'header-slack-link header-logo-24',
          position: 'right',
        },
        {
          href: 'https://github.com/commitdev/zero',
          label: 'Github',
          className: 'header-github-link header-logo-24',
          position: 'right',
        },
      ],
    },
    footer: {
      links: [
        {
          items: [
            {
              to: '/docs/zero/about/overview',
              label: 'Docs',
              className: 'header-docs-link header-logo-24',
              position: 'right'
            },
            {
              href: 'https://slack.getzero.dev',
              label: 'Slack',
              className: 'header-slack-link header-logo-24',
              position: 'right',
            },
            {
              href: 'https://github.com/commitdev/zero',
              label: 'Github',
              className: 'header-github-link header-logo-24',
              position: 'right',
            },
          ],
        },
      ],
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          path: 'docs',
          routeBasePath: 'docs/zero/',
          include: ['**/*.md', '**/*.mdx'],
          // editUrl: 'https://github.com/commitdev/zero/blob/main/doc-site/',
          editUrl: 'https://github.com/commitdev/zero/blob/doc-site/doc-site/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
        debug: true,
      },
    ],
  ],
  plugins: [
    'docusaurus-plugin-sass'
  ],
  stylesheets: [
    "https://fonts.googleapis.com/css2?family=Lato:wght@400;700;900&family=Montserrat:wght@400;600;700;800&display=swap",
  ]
};
