import { defineStore } from 'pinia';
// 重新整理頁面資料會消失， 改用LocalStorage
export const useUserStore = defineStore('user', {
  state: () => ({
    user: null,
    email: 'abc@gmail.com',
    department: 'HR',
    name: 'test',
    password: null,
    session: null,
    userid: null
  }),
  getters: {
    getName: (state) => state.name,
    getEmail: (state) => state.email,
    getDepartment: (state) => state.department,
    getSession: (state) => state.session,
    getPassword: (state) => state.password,
    getUserid: (state) => state.userid
  },
  actions: {
    login (user) {
      // 执行登录操作，并更新用户状态
      this.user = user;
      this.name = user.name;
      this.email = user.email;
      this.deparement = user.deparement;
      this.session = user.session;
      this.password = user.password;
      this.userid = user.userid;
    },
    logout () {
      // 执行登出操作，并清空用户状态
      this.user = null;
    },
    modify (user) {
      this.user = user;
      this.name = user.name;
      this.email = user.email;
      this.deparement = user.deparement;
      this.password = user.password;
    }
  }
});
