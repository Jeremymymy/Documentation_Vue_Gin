import { defineStore } from 'pinia';

export const useUserStore = defineStore('user', {
  state: () => ({
    user: null
  }),
  actions: {
    login (user) {
      // 执行登录操作，并更新用户状态
      this.user = user;
      this.name = user.name;
    },
    logout () {
      // 执行登出操作，并清空用户状态
      this.user = null;
    }
  }
});
