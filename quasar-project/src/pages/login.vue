<template>
  <div class="row justify-center full-height full-width text-center">
    <div class="column q-pa-lg">
      <div class="row">
        <q-card square dark class="q-pa-md q-ma-none no-shadow bg-grey-10" style="width:360px;">
          <q-card-section class="q-mt-xl q-mb-md">
            <q-btn round>
              <q-avatar rounded>
                <img src="~assets/tsmc.png">
              </q-avatar>
            </q-btn>
            <p class="text-h5 text-white">TSMC Document System</p>
            <p class="text-h5 text-white">Login </p>
          </q-card-section>
          <q-card-section>
            <q-form @submit.prevent="login" class="q-gutter-md">
              <q-input dark dense square filled clearable v-model="user.email" type="email" label="Email" class="form-control"
                id="inputEmail" aria-describedby="emailHelp">
                <template v-slot:prepend>
                  <q-icon name="email" />
                </template>
              </q-input>
              <q-input dark dense square filled clearable v-model="user.password" type="password" label="Password"
                class="form-control" id="inputPassword">
                <template v-slot:prepend>
                  <q-icon name="lock" />
                </template>
              </q-input>
              <q-btn outline rounded size="mid" color="red-4" class="full-width text-white" type="submit" label="Sign In" />
              <q-btn outline rounded size="mid" color="red-4" class="full-width text-white" label="register" router-link to="/register" />
            </q-form>
          </q-card-section>
          <q-card-actions>
            <div class="row full-width items-center">
              <div class="col-6">
              </div>
            </div>
          </q-card-actions>
        </q-card>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import { useUserStore } from 'src/stores/user';
// import { useQuasar } from 'quasar';
import { LocalStorage, SessionStorage } from 'quasar';
// const $q = useQuasar();
export default {
  name: 'login',
  data () {
    return {
      // 與使用者輸入的帳密欄位 v-model 雙向綁定
      user: {
        email: '',
        password: ''
      }
    };
  },
  methods: {
    login () {
      const userStore = useUserStore();

      // 如果一開始沒有 import axios 的話
      // 用 require("axios").default 也可以讓我們使用 axios
      // const axios = require("axios").default;
      axios
        .post('http://localhost:8000/TSMC/users/login', this.user)
        .then((res) => {
          console.log(res);
          if (res.statusText === 'OK') {
            // this.user.password
            userStore.login({ userid: res.data.User.EmployeeId, name: res.data.User.Name, email: res.data.User.Email, password: '$2a$10$/duU/Q6c2x.iPH5aQcngEOMo1iE5lgaFRolfj9I5o0cOLn/Px2m/i', session: res.data.Sessions, department: res.data.User.Department });

            try {
              LocalStorage.set('userInfo', res.data.User)
              SessionStorage.set('userSession', res.data.Sessions)
              const value = SessionStorage.getItem('userSession')
              const value2 = LocalStorage.getItem('userInfo')

              console.log(value)
              console.log(value2)
            } catch (e) {
              // 由于Web Storage API错误，
              // 数据未成功保存
              console.log(e)
            }
            console.log(res.data.User.Name)
            // 登录成功，进行页面跳转
            this.$router.push({
              path: '/index'
            }); // 跳转到指定页面
          }
        })
        .catch(error => {
          console.log(error)
          // 处理登录失败的逻辑
        });
    }
  }
}
</script>

<style>
</style>
