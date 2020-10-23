import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import ProvinceTable from './components/Table';
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/table', ProvinceTable);
  },
});
